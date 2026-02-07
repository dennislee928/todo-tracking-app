package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// verifyAppleReceipt calls Apple's verifyReceipt API.
// Returns true if the receipt is valid and contains a non-consumable or subscription.
func VerifyAppleReceipt(receiptData, sharedSecret string) (bool, error) {
	if sharedSecret == "" {
		return false, fmt.Errorf("apple shared secret not configured")
	}

	reqBody := map[string]interface{}{
		"receipt-data": receiptData,
		"password":     sharedSecret,
	}
	body, _ := json.Marshal(reqBody)

	// Use production first, fallback to sandbox
	urls := []string{
		"https://buy.itunes.apple.com/verifyReceipt",
		"https://sandbox.itunes.apple.com/verifyReceipt",
	}

	for _, url := range urls {
		resp, err := http.Post(url, "application/json", bytes.NewReader(body))
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		data, _ := io.ReadAll(resp.Body)
		var result struct {
			Status  int `json:"status"`
			Receipt struct {
				InApp []struct {
					ProductID string `json:"product_id"`
				} `json:"in_app"`
			} `json:"receipt"`
		}
		if err := json.Unmarshal(data, &result); err != nil {
			continue
		}

		// 21007 = sandbox receipt sent to production
		if result.Status == 21007 {
			continue
		}
		if result.Status == 0 && len(result.Receipt.InApp) > 0 {
			return true, nil
		}
		return false, fmt.Errorf("apple status %d", result.Status)
	}
	return false, fmt.Errorf("apple verify failed")
}

// VerifyGooglePurchase verifies a Google Play purchase token.
// Requires Google Play Developer API and a service account with permissions.
func VerifyGooglePurchase(purchaseToken, productID, packageName, serviceAccountJSON string) (bool, error) {
	if packageName == "" || serviceAccountJSON == "" {
		return false, fmt.Errorf("google play not configured")
	}

	// Google Play Developer API: purchases.products.get
	// Requires OAuth2 token from service account.
	// For a minimal implementation, we return an error suggesting full setup.
	// Full implementation would use google.golang.org/api/androidpublisher and
	// google.golang.org/api/option with credentials.
	_ = purchaseToken
	_ = productID
	_ = packageName
	_ = serviceAccountJSON
	return false, fmt.Errorf("google play verification requires androidpublisher API setup - see implementation guide")
}
