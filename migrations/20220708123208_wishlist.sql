-- +goose Up
CREATE TABLE wishlist (
    wishlist_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(user_id) NOT NULL,
    product_id BIGINT REFERENCES products(product_id) NOT NULL,
    inventory_id BIGINT REFERENCES inventories(inventory_id) NOT NULL
);

-- +goose Down
DROP TABLE wishlist CASCADE;

