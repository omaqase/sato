SET search_path TO user_management, public;

CREATE TABLE favorites
(
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id    UUID        NOT NULL REFERENCES users (id),
    product_id UUID        NOT NULL,
    created_at TIMESTAMP   NOT NULL DEFAULT NOW(),
    UNIQUE (user_id, product_id)
);

CREATE INDEX idx_favorites_user_id ON favorites (user_id);