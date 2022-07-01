-- +goose Up
-- +goose StatementBegin
CREATE TABLE brands(
    brand_id BIGSERIAL PRIMARY KEY,
    brand_name VARCHAR(100) NOT NULL,
    brand_desc VARCHAR(255),
    brand_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    brand_updated_at TIMESTAMP DEFAULT NULL,
    brand_deleted_at TIMESTAMP DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE brands CASCADE;
-- +goose StatementEnd
