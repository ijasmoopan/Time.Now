-- +goose Up
-- +goose StatementBegin

CREATE TABLE subcategories(
    subcategory_id BIGSERIAL PRIMARY KEY,
    subcategory_name VARCHAR(100) NOT NULL,
    subcategory_desc VARCHAR(255),
    subcategory_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    subcategory_updated_at TIMESTAMP DEFAULT NULL,
    subcategory_deleted_at TIMESTAMP DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE subcategories CASCADE;
-- +goose StatementEnd
