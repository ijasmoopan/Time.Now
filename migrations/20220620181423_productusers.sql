-- +goose Up

CREATE TABLE productusers(
    product_id INT,
    user_id INT
);


-- +goose Down

DROP TABLE productusers CASCADE;

