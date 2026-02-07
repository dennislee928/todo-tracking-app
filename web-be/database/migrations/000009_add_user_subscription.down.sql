-- Remove subscription columns from users table
DROP INDEX IF EXISTS idx_users_is_premium;
ALTER TABLE users DROP COLUMN IF EXISTS premium_expires_at;
ALTER TABLE users DROP COLUMN IF EXISTS is_premium;
