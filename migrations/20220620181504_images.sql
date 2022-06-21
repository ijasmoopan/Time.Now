-- +goose Up
-- +goose StatementBegin

CREATE TABLE images(
    product_id INT,
    product_image TEXT
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE images;

-- +goose StatementEnd
