-- +goose Up
CREATE TABLE orders (
    order_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(user_id) NOT NULL,
    address_id BIGINT REFERENCES address(address_id) NOT NULL,
    product_id BIGINT REFERENCES products(product_id) NOT NULL,
    payment_id BIGINT REFERENCES payments(payment_id) NOT NULL,
    inventory_id BIGINT REFERENCES inventories(inventory_id) NOT NULL,
    order_status VARCHAR(100) NOT NULL DEFAULT 'Ordered',
    cart_id BIGINT REFERENCES cart(cart_id) NOT NULL,
    ordered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    replaced_at TIMESTAMP DEFAULT NULL,
    delivered_at TIMESTAMP DEFAULT NULL,
    cancelled_at TIMESTAMP DEFAULT NULL
);

-- +goose Down
DROP TABLE orders CASCADE;

