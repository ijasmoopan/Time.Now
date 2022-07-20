-- +goose Up
CREATE TABLE categories (
    category_id BIGSERIAL PRIMARY KEY,
    category_name VARCHAR(100) NOT NULL,
    category_desc VARCHAR(255),
    category_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    category_updated_at TIMESTAMP DEFAULT NULL,
    category_deleted_at TIMESTAMP DEFAULT NULL
);

-- +goose Down
DROP TABLE categories CASCADE;
