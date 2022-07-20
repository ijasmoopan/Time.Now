-- +goose Up

CREATE TABLE images(
    image_id BIGSERIAL PRIMARY KEY,
    product_id INT,
    product_image TEXT
);


-- +goose Down

DROP TABLE images CASCADE;

