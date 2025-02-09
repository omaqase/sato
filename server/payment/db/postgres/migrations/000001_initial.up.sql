create schema if not exists payment_management;

set search_path to payment_management, public;

create type payment_status as enum ('pending', 'completed', 'failed');

create table payment_methods
(
    id                    uuid primary key default gen_random_uuid(),
    user_id               uuid references user_management.users (id) on delete cascade,
    card_number_encrypted bytea        not null,
    cardholder_name       varchar(255) not null,
    expiration_month      smallint     not null,
    expiration_year       smallint     not null,
    cvv_encrypted         bytea        not null,
    created_at            timestamp        default current_timestamp,
    updated_at            timestamp        default current_timestamp
);

create table payments
(
    id                    uuid primary key        default gen_random_uuid(),
    user_id               uuid references user_management.users (id) on delete cascade,
    payment_method_id     uuid           references payment_methods (id) on delete set null,
    amount                decimal(10, 2) not null,
    currency              char(3)        not null,
    status                payment_status not null default 'pending',
    transaction_reference varchar(255),
    created_at            timestamp               default current_timestamp,
    updated_at            timestamp               default current_timestamp
);

create index idx_payment_methods_user_id on payment_methods (id);
create index idx_payments_user_id on payments (user_id);
create index idx_payments_status on payments (status);