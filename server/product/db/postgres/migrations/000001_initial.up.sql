-- Создание схемы и установка пути поиска
CREATE SCHEMA IF NOT EXISTS product_management;
SET search_path TO product_management, public;

-- Таблица категорий
CREATE TABLE categories
(
    id         UUID PRIMARY KEY   DEFAULT gen_random_uuid(),
    title      TEXT      NOT NULL,
    slug       TEXT      NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- Таблица иерархии категорий
CREATE TABLE categories_hierarchy
(
    ancestor   UUID NOT NULL REFERENCES categories (id),
    descendant UUID NOT NULL REFERENCES categories (id),
    depth      INT  NOT NULL,
    PRIMARY KEY (ancestor, descendant)
);

-- Определение ENUM типа для статуса одобрения продукта
CREATE TYPE product_approve_status AS ENUM ('pending', 'approved', 'rejected');

-- Таблица продуктов
CREATE TABLE products
(
    id          UUID PRIMARY KEY                DEFAULT gen_random_uuid(),
    title       TEXT                   NOT NULL,
    category_id UUID                   NOT NULL REFERENCES categories (id),
    seller_id   UUID                   NOT NULL REFERENCES user_management.sellers (user_id),
    stock       INT                    NOT NULL CHECK (stock >= 0),
    rating      NUMERIC(3, 2)          NOT NULL CHECK (rating >= 0 AND rating <= 5),
    approved    product_approve_status NOT NULL DEFAULT 'pending',
    created_at  TIMESTAMP              NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP              NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMP
);

-- Таблица отзывов о продуктах
CREATE TABLE products_reviews
(
    id         UUID PRIMARY KEY   DEFAULT gen_random_uuid(),
    product_id UUID      NOT NULL REFERENCES products (id),
    user_id    UUID      NOT NULL REFERENCES user_management.users (id),
    rating     INT       NOT NULL CHECK (rating >= 1 AND rating <= 5),
    comment    TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    UNIQUE (user_id, product_id)
);

-- Таблица ответов на отзывы
CREATE TABLE products_reviews_replies
(
    id         UUID PRIMARY KEY   DEFAULT gen_random_uuid(),
    review_id  UUID      NOT NULL REFERENCES products_reviews (id),
    seller_id  UUID      NOT NULL REFERENCES user_management.sellers (user_id),
    reply      TEXT      NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- Таблица вопросов о продуктах
CREATE TABLE products_questions
(
    id         UUID PRIMARY KEY   DEFAULT gen_random_uuid(),
    user_id    UUID      NOT NULL REFERENCES user_management.users (id),
    product_id UUID      NOT NULL REFERENCES products (id),
    question   TEXT      NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    UNIQUE (user_id, product_id)
);

-- Таблица ответов на вопросы
CREATE TABLE products_questions_replies
(
    id          UUID PRIMARY KEY   DEFAULT gen_random_uuid(),
    question_id UUID      NOT NULL REFERENCES products_questions (id),
    seller_id   UUID      NOT NULL REFERENCES user_management.sellers (user_id),
    reply       TEXT      NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMP
);