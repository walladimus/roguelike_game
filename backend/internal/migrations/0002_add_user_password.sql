-- 0002_add_user_password.sql
-- Adds password storage to users table for auth persistence.

ALTER TABLE users
ADD COLUMN IF NOT EXISTS password_hash TEXT NOT NULL DEFAULT '';

CREATE UNIQUE INDEX IF NOT EXISTS idx_users_username ON users (username);
