CREATE SCHEMA IF NOT EXISTS user_management;

SET search_path TO user_management, public;

CREATE TABLE user_management.users
(
    id                          UUID PRIMARY KEY      DEFAULT gen_random_uuid(),
    email                       VARCHAR(255) NOT NULL,
    password                    VARCHAR(255) NOT NULL,
    first_name                  VARCHAR(255) NOT NULL,
    last_name                   VARCHAR(255) NOT NULL,
    phone                       VARCHAR(255),
    created_at                  TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at                  TIMESTAMP    NOT NULL DEFAULT NOW(),
    deleted_at                  TIMESTAMP,
    is_subscribed_to_promotions BOOLEAN      NOT NULL DEFAULT FALSE
);

CREATE UNIQUE INDEX idx_users_email
    ON user_management.users (email);

CREATE INDEX idx_users_promotion_subscribers
    ON user_management.users (is_subscribed_to_promotions)
    WHERE is_subscribed_to_promotions = TRUE;

CREATE INDEX idx_users_id_deleted
    ON user_management.users (id)
    WHERE deleted_at IS NULL;

CREATE INDEX idx_users_email_active
    ON user_management.users (email)
    WHERE deleted_at IS NULL;

CREATE INDEX idx_users_promo_active_email
    ON user_management.users (is_subscribed_to_promotions, email)
    WHERE deleted_at IS NULL;