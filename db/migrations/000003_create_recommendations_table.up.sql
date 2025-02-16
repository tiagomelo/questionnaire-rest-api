CREATE TABLE IF NOT EXISTS recommendations (
    id SERIAL PRIMARY KEY,
    ulid TEXT UNIQUE NOT NULL CHECK (ulid ~ '^[0-9A-HJKMNP-TV-Z]{26}$'),
    product_ulid TEXT NOT NULL,
    CONSTRAINT fk_product FOREIGN KEY (product_ulid) REFERENCES products(ulid)
);