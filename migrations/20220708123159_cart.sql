-- +goose Up
CREATE TABLE cart (
    cart_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(user_id) NOT NULL,
    product_id BIGINT REFERENCES products(product_id) NOT NULL,
    inventory_id BIGINT REFERENCES inventories(inventory_id) NOT NULL,
    quantity BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- +goose Down
DROP TABLE cart CASCADE;

