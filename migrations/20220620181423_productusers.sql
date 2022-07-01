-- +goose Up
-- +goose StatementBegin

CREATE TABLE productusers(
    product_id INT,
    user_id INT
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE productusers CASCADE;

-- +goose StatementEnd
