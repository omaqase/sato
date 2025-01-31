-- Создание схемы
CREATE SCHEMA IF NOT EXISTS product_management;
SET search_path TO product_management, public;

-- Таблица для категорий
CREATE TABLE product_management.categories
(
    id         UUID PRIMARY KEY      DEFAULT gen_random_uuid(),
    title      VARCHAR(255) NOT NULL,
    slug       VARCHAR(255) NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- Индексы для категорий
CREATE UNIQUE INDEX idx_categories_slug
    ON product_management.categories (slug);
CREATE INDEX idx_categories_id_deleted
    ON product_management.categories (id)
    WHERE deleted_at IS NULL;
CREATE INDEX idx_categories_slug_active
    ON product_management.categories (slug)
    WHERE deleted_at IS NULL;

-- Таблица для продуктов
CREATE TABLE product_management.products
(
    id          UUID PRIMARY KEY        DEFAULT gen_random_uuid(),
    title       VARCHAR(255)   NOT NULL,
    description TEXT,
    price       NUMERIC(10, 2) NOT NULL CHECK (price >= 0),
    category_id UUID           NOT NULL REFERENCES product_management.categories (id),
    created_at  TIMESTAMP      NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP      NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMP
);

-- Индексы для продуктов
CREATE INDEX idx_products_category_id
    ON product_management.products (category_id)
    WHERE deleted_at IS NULL;
CREATE INDEX idx_products_price
    ON product_management.products (price)
    WHERE deleted_at IS NULL;
CREATE INDEX idx_products_id_deleted
    ON product_management.products (id)
    WHERE deleted_at IS NULL;