package rest

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
	"github.com/stripe/stripe-go/v81/webhook"
	"gorm.io/gorm"

	"github.com/todo-tracking-app/web-be/internal/config"
	"github.com/todo-tracking-app/web-be/internal/model"
	"github.com/todo-tracking-app/web-be/internal/service"
)

// RegisterSubscriptionRoutes registers subscription routes.
// Stripe webhook must NOT use auth middleware (Stripe sends raw POST).
func RegisterSubscriptionRoutes(r *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	h := &subscriptionHandler{db: db, cfg: cfg}
	// Webhook: no auth
	r.POST("/subscription/stripe-webhook", h.StripeWebhook)
}

// RegisterSubscriptionProtectedRoutes registers subscription routes that require auth.
func RegisterSubscriptionProtectedRoutes(r *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	h := &subscriptionHandler{db: db, cfg: cfg}
	r.POST("/subscription/create-checkout-session", h.CreateCheckoutSession)
	r.POST("/subscription/apple-verify", h.AppleVerify)
	r.POST("/subscription/google-verify", h.GoogleVerify)
}

type subscriptionHandler struct {
	db  *gorm.DB
	cfg *config.Config
}

// CreateCheckoutSessionRequest is the request for creating a Stripe Checkout session.
type CreateCheckoutSessionRequest struct {
	SuccessURL string `json:"success_url" binding:"required"`
	CancelURL  string `json:"cancel_url" binding:"required"`
}

// CreateCheckoutSessionResponse is the response with checkout session URL.
type CreateCheckoutSessionResponse struct {
	URL string `json:"url"`
}

// CreateCheckoutSession creates a Stripe Checkout session for premium upgrade.
// @Summary Create Stripe Checkout session
// @Tags subscription
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body CreateCheckoutSessionRequest true "Checkout session request"
// @Success 200 {object} CreateCheckoutSessionResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /subscription/create-checkout-session [post]
func (h *subscriptionHandler) CreateCheckoutSession(c *gin.Context) {
	userID, _ := c.Get("user_id")
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if h.cfg.StripeSecretKey == "" || h.cfg.StripePriceID == "" {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Stripe not configured"})
		return
	}

	var req CreateCheckoutSessionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stripe.Key = h.cfg.StripeSecretKey

	params := &stripe.CheckoutSessionParams{
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(req.SuccessURL),
		CancelURL:  stripe.String(req.CancelURL),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(h.cfg.StripePriceID),
				Quantity: stripe.Int64(1),
			},
		},
		Metadata: map[string]string{
			"user_id": userID.(string),
		},
	}

	sess, err := session.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, CreateCheckoutSessionResponse{URL: sess.URL})
}

// StripeWebhook handles Stripe webhook events (checkout.session.completed).
func (h *subscriptionHandler) StripeWebhook(c *gin.Context) {
	if h.cfg.StripeWebhookSecret == "" {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "webhook secret not configured"})
		return
	}

	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "read body failed"})
		return
	}

	sigHeader := c.GetHeader("Stripe-Signature")
	event, err := webhook.ConstructEvent(payload, sigHeader, h.cfg.StripeWebhookSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid signature"})
		return
	}

	switch event.Type {
	case "checkout.session.completed":
		var sess stripe.CheckoutSession
		if err := json.Unmarshal(event.Data.Raw, &sess); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid session object"})
			return
		}
		userID, ok := sess.Metadata["user_id"]
		if !ok || userID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing user_id in metadata"})
			return
		}

		expiresAt := time.Now().Add(365 * 24 * time.Hour) // 1 year for one-time payment
		if err := h.db.Model(&model.User{}).Where("id = ?", userID).
			Updates(map[string]interface{}{
				"is_premium":         true,
				"premium_expires_at": expiresAt,
			}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
			return
		}
	default:
		// Ignore other events
	}

	c.JSON(http.StatusOK, gin.H{"received": true})
}

// AppleVerifyRequest is the request for Apple IAP verification.
type AppleVerifyRequest struct {
	ReceiptData string `json:"receipt_data" binding:"required"`
}

// AppleVerify verifies Apple IAP receipt and upgrades user to premium.
// @Summary Verify Apple IAP receipt
// @Tags subscription
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body AppleVerifyRequest true "Receipt data"
// @Success 200 {object} map[string]bool
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /subscription/apple-verify [post]
func (h *subscriptionHandler) AppleVerify(c *gin.Context) {
	userID, _ := c.Get("user_id")
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req AppleVerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call Apple verifyReceipt API
	valid, err := service.VerifyAppleReceipt(req.ReceiptData, h.cfg.AppleSharedSecret)
	if err != nil || !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid receipt"})
		return
	}

	expiresAt := time.Now().Add(365 * 24 * time.Hour)
	if err := h.db.Model(&model.User{}).Where("id = ?", userID.(string)).
		Updates(map[string]interface{}{
			"is_premium":         true,
			"premium_expires_at": expiresAt,
		}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"is_premium": true})
}

// GoogleVerifyRequest is the request for Google Play purchase verification.
type GoogleVerifyRequest struct {
	PurchaseToken string `json:"purchase_token" binding:"required"`
	ProductID     string `json:"product_id" binding:"required"`
}

// GoogleVerify verifies Google Play purchase and upgrades user to premium.
// @Summary Verify Google Play purchase
// @Tags subscription
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body GoogleVerifyRequest true "Purchase token and product ID"
// @Success 200 {object} map[string]bool
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /subscription/google-verify [post]
func (h *subscriptionHandler) GoogleVerify(c *gin.Context) {
	userID, _ := c.Get("user_id")
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req GoogleVerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	valid, err := service.VerifyGooglePurchase(req.PurchaseToken, req.ProductID, h.cfg.GooglePackageName, h.cfg.GoogleServiceAccountJSON)
	if err != nil || !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid purchase"})
		return
	}

	expiresAt := time.Now().Add(365 * 24 * time.Hour)
	if err := h.db.Model(&model.User{}).Where("id = ?", userID.(string)).
		Updates(map[string]interface{}{
			"is_premium":         true,
			"premium_expires_at": expiresAt,
		}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"is_premium": true})
}
