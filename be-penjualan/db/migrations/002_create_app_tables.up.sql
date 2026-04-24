-- Migration: 002_create_app_tables
-- Creates application configuration tables in the `app` schema.

-- Users of the InsightFlow application (admin, manager, sales, viewer)
CREATE TABLE app.users (
    id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nama             VARCHAR(100)  NOT NULL,
    email            VARCHAR(150)  UNIQUE NOT NULL,
    password         VARCHAR(255)  NOT NULL,      -- bcrypt hash
    role             VARCHAR(50)   NOT NULL CHECK (role IN ('admin','manager','sales','viewer')),
    telegram_user_id BIGINT        UNIQUE,         -- maps to Telegram user for Q&A per-role
    aktif            BOOLEAN       NOT NULL DEFAULT TRUE,
    created_at       TIMESTAMPTZ   NOT NULL DEFAULT NOW()
);

-- Telegram group configuration (daily summary, anomaly alerts)
CREATE TABLE app.telegram_config (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nama_grup   VARCHAR(100)  NOT NULL,
    chat_id     BIGINT        UNIQUE NOT NULL,     -- Telegram chat_id (negative for groups)
    aktif       BOOLEAN       NOT NULL DEFAULT TRUE,
    jam_summary TIME          NOT NULL DEFAULT '00:00', -- stored in UTC; 00:00 UTC = 07:00 WIB
    created_at  TIMESTAMPTZ   NOT NULL DEFAULT NOW()
);

-- Per-metric anomaly thresholds (answered: threshold is per-metric, not global)
CREATE TABLE app.anomaly_config (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    metric_key    VARCHAR(100)   UNIQUE NOT NULL,  -- e.g. daily_revenue, order_count
    threshold_pct DECIMAL(5,2)   NOT NULL DEFAULT 10.00,  -- e.g. 10.00 = 10% deviation triggers alert
    aktif         BOOLEAN        NOT NULL DEFAULT TRUE,
    updated_at    TIMESTAMPTZ    NOT NULL DEFAULT NOW()
);

-- Saved dashboard configurations per user
CREATE TABLE app.saved_dashboards (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id      UUID         REFERENCES app.users(id) ON DELETE CASCADE,
    nama         VARCHAR(150) NOT NULL,
    konfigurasi  JSONB,                            -- stores report type, filters, etc.
    created_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

-- Index for quick lookup by Telegram user during Q&A
CREATE INDEX idx_users_telegram_user_id ON app.users (telegram_user_id) WHERE telegram_user_id IS NOT NULL;
