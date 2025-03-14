-- Create new migration file: server/catalogue/db/postgres/migrations/000004_create_favorites.up.sql
CREATE TABLE IF NOT EXISTS product_management.user_favorites (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    product_id UUID NOT NULL REFERENCES product_management.products(id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE(user_id, product_id)
);