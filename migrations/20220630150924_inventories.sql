-- +goose Up
-- +goose StatementBegin

CREATE TABLE inventories(
    inventory_id BIGSERIAL PRIMARY KEY,
    product_id BIGINT REFERENCES products(product_id) NOT NULL,
    product_color VARCHAR(50) NOT NULL,
    product_quantity BIGINT NOT NULL,
    inventory_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    inventory_updated_at TIMESTAMP DEFAULT NULL,
    inventory_deleted_at TIMESTAMP DEFAULT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE inventories CASCADE;

-- +goose StatementEnd
