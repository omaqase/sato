-- Создание схемы и установка пути поиска
CREATE SCHEMA IF NOT EXISTS user_management;
SET search_path TO user_management, public;

-- Определение ENUM типа для ролей пользователей
CREATE TYPE user_role AS ENUM ('customer', 'seller', 'admin', 'super_admin');

-- Таблица методов оплаты (перенесена выше users)
CREATE TABLE payment_methods
(
    id      UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    card    TEXT NOT NULL,
    date    TEXT NOT NULL,
    cvv     TEXT NOT NULL,
    UNIQUE (user_id, card)
);

-- Таблица пользователей
CREATE TABLE users
(
    id                UUID PRIMARY KEY   DEFAULT gen_random_uuid(),
    email             TEXT      NOT NULL UNIQUE,
    role              user_role NOT NULL DEFAULT 'customer',
    first_name        TEXT      NOT NULL,
    last_name         TEXT      NOT NULL,
    phone             TEXT,
    payment_method_id UUID REFERENCES payment_methods (id),
    promotions        BOOLEAN   NOT NULL DEFAULT FALSE,
    email_verified    BOOLEAN   NOT NULL DEFAULT FALSE,
    two_factor        BOOLEAN   NOT NULL DEFAULT FALSE,
    created_at        TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at        TIMESTAMP
);

-- Добавление внешнего ключа для payment_methods.user_id после создания таблицы users
ALTER TABLE payment_methods
    ADD CONSTRAINT fk_payment_methods_user_id
        FOREIGN KEY (user_id) REFERENCES users (id);

-- Таблица продавцов
CREATE TABLE sellers
(
    user_id    UUID PRIMARY KEY REFERENCES users (id),
    name       TEXT          NOT NULL,
    rating     NUMERIC(3, 2) NOT NULL CHECK (rating >= 0 AND rating <= 5),
    approved   BOOLEAN       NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP     NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP     NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- Таблица корзин
CREATE TABLE baskets
(
    id      UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users (id),
    UNIQUE (user_id)
);

-- Таблица элементов корзины
CREATE TABLE basket_items
(
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    basket_id  UUID NOT NULL REFERENCES baskets (id),
    product_id UUID NOT NULL,
    quantity   INT  NOT NULL CHECK (quantity > 0),
    UNIQUE (basket_id, product_id)
);