-- +goose Up
CREATE TABLE inventories(
    inventory_id BIGSERIAL PRIMARY KEY,
    product_id BIGINT REFERENCES products(product_id) NOT NULL,
    color_id BIGINT REFERENCES colors(color_id) NOT NULL,
    product_quantity BIGINT NOT NULL,
    inventory_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    inventory_updated_at TIMESTAMP DEFAULT NULL,
    inventory_deleted_at TIMESTAMP DEFAULT NULL
);

-- +goose Down
DROP TABLE inventories CASCADE;

