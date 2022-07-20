-- +goose Up
CREATE TABLE colors(
    color_id BIGSERIAL PRIMARY KEY,
    color VARCHAR(100) NOT NULL,
    color_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    color_updated_at TIMESTAMP DEFAULT NULL,
    color_deleted_at TIMESTAMP DEFAULT NULL
);

-- +goose Down
DROP TABLE colors CASCADE;
