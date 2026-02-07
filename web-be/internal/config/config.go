package config

import "os"

// Config holds application configuration.
type Config struct {
	DatabaseURL       string
	JWTSecret         string
	SupabaseURL       string
	SupabaseAnonKey   string
	SupabaseJWTSecret string
	// Stripe
	StripeSecretKey      string
	StripeWebhookSecret  string
	StripePriceID        string
	// Apple IAP
	AppleSharedSecret string
	// Google Play
	GooglePackageName       string
	GoogleServiceAccountJSON string
}

// Load reads configuration from environment variables.
func Load() (*Config, error) {
	return &Config{
		DatabaseURL:            getEnv("DATABASE_URL", "postgres://localhost:5432/todo?sslmode=disable"),
		JWTSecret:              getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
		SupabaseURL:            getEnv("SUPABASE_URL", ""),
		SupabaseAnonKey:        getEnv("SUPABASE_ANON_KEY", ""),
		SupabaseJWTSecret:      getEnv("SUPABASE_JWT_SECRET", ""),
		StripeSecretKey:        getEnv("STRIPE_SECRET_KEY", ""),
		StripeWebhookSecret:    getEnv("STRIPE_WEBHOOK_SECRET", ""),
		StripePriceID:          getEnv("STRIPE_PRICE_ID", ""),
		AppleSharedSecret:      getEnv("APPLE_SHARED_SECRET", ""),
		GooglePackageName:      getEnv("GOOGLE_PACKAGE_NAME", ""),
		GoogleServiceAccountJSON: getEnv("GOOGLE_SERVICE_ACCOUNT_JSON", ""),
	}, nil
}

func getEnv(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
