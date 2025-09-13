CREATE DATABASE userservicedb;

\c userservicedb;

CREATE TYPE user_role_enum AS ENUM ('user', 'buyer', 'seller');

CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role user_role_enum DEFAULT 'user',
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE INDEX idx_products_user_id ON products(user_id);
