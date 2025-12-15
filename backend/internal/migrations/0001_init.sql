-- 0001_init.sql
-- Initial schema for the roguelike backend

-- Users: minimal identity for players (can be guest or later linked to real auth)
CREATE TABLE IF NOT EXISTS users (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username     TEXT NOT NULL,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Friends: simple directed friend relationship with a status
CREATE TABLE IF NOT EXISTS friends (
    id            BIGSERIAL PRIMARY KEY,
    user_id       UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    friend_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    status        TEXT NOT NULL DEFAULT 'pending', -- pending / accepted / blocked etc.
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (user_id, friend_id)
);

-- Lobbies: where players gather before/while playing
CREATE TABLE IF NOT EXISTS lobbies (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    host_user_id  UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    code          TEXT NOT NULL UNIQUE,           -- short join code
    status        TEXT NOT NULL DEFAULT 'open',   -- open / in_game / closed
    max_players   INT  NOT NULL DEFAULT 4,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Lobby members: which users are currently in which lobby
CREATE TABLE IF NOT EXISTS lobby_members (
    id          BIGSERIAL PRIMARY KEY,
    lobby_id    UUID NOT NULL REFERENCES lobbies(id) ON DELETE CASCADE,
    user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    joined_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (lobby_id, user_id)
);

-- Notices: in-game notifications to users (match invites, friend requests, etc.)
CREATE TABLE IF NOT EXISTS notices (
    id          BIGSERIAL PRIMARY KEY,
    user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type        TEXT NOT NULL,               -- e.g. 'friend_request', 'system', etc.
    title       TEXT NOT NULL,
    body        TEXT NOT NULL,
    is_read     BOOLEAN NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Achievements: list of possible achievements
CREATE TABLE IF NOT EXISTS achievements (
    id          BIGSERIAL PRIMARY KEY,
    code        TEXT NOT NULL UNIQUE,        -- stable identifier
    name        TEXT NOT NULL,
    description TEXT NOT NULL
);

-- User achievements: which achievements a user has unlocked
CREATE TABLE IF NOT EXISTS user_achievements (
    id              BIGSERIAL PRIMARY KEY,
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    achievement_id  BIGINT NOT NULL REFERENCES achievements(id) ON DELETE CASCADE,
    unlocked_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (user_id, achievement_id)
);

-- Useful indexes
CREATE INDEX IF NOT EXISTS idx_friends_user_id
    ON friends (user_id);

CREATE INDEX IF NOT EXISTS idx_friends_friend_id
    ON friends (friend_id);

CREATE INDEX IF NOT EXISTS idx_lobby_members_lobby_id
    ON lobby_members (lobby_id);

CREATE INDEX IF NOT EXISTS idx_notices_user_id_is_read
    ON notices (user_id, is_read);

CREATE INDEX IF NOT EXISTS idx_user_achievements_user_id
    ON user_achievements (user_id);