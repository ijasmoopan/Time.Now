-- +goose Up
-- +goose StatementBegin

CREATE TABLE images(
    image_id BIGSERIAL PRIMARY KEY,
    product_id INT,
    product_image TEXT
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE images CASCADE;

-- +goose StatementEnd
