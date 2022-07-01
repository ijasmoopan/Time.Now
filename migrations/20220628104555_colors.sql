-- +goose Up
-- +goose StatementBegin
CREATE TABLE colors(
    color_id BIGSERIAL PRIMARY KEY,
    color VARCHAR(100),
    product_id BIGINT REFERENCES products(product_id) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE colors CASCADE;
-- +goose StatementEnd
