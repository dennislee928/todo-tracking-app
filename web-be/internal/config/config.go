package config

import "os"

// Config holds application configuration.
type Config struct {
	DatabaseURL       string
	JWTSecret         string
	SupabaseURL       string
	SupabaseAnonKey   string
	SupabaseJWTSecret string
}

// Load reads configuration from environment variables.
func Load() (*Config, error) {
	return &Config{
		DatabaseURL:       getEnv("DATABASE_URL", "postgres://localhost:5432/todo?sslmode=disable"),
		JWTSecret:         getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
		SupabaseURL:       getEnv("SUPABASE_URL", ""),
		SupabaseAnonKey:   getEnv("SUPABASE_ANON_KEY", ""),
		SupabaseJWTSecret: getEnv("SUPABASE_JWT_SECRET", ""),
	}, nil
}

func getEnv(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
