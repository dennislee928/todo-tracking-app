-- Add subscription columns to users table
ALTER TABLE users ADD COLUMN IF NOT EXISTS is_premium BOOLEAN NOT NULL DEFAULT FALSE;
ALTER TABLE users ADD COLUMN IF NOT EXISTS premium_expires_at TIMESTAMPTZ;

CREATE INDEX IF NOT EXISTS idx_users_is_premium ON users(is_premium);
