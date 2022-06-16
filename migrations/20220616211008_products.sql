-- +goose Up
-- +goose StatementBegin

CREATE TABLE products (
    product_id BIGSERIAL PRIMARY KEY,
    product_name VARCHAR(100) NOT NULL,
    product_price NUMERIC(16, 2) NOT NULL,
    product_color VARCHAR(30) NOT NULL,
    product_desc TEXT NOT NULL,
    product_quantity INT NOT NULL,
    product_image TEXT DEFAULT NULL
);

INSERT INTO products (product_name, product_price, product_color, product_desc, product_quantity, product_image)
VALUES ('Seven Friday', '1199', 'Red', 'Beautiful watch', 10, 'images/defaultproduct.png');

INSERT INTO products (product_name, product_price, product_color, product_desc, product_quantity, product_image)
VALUES ('Hublot', '899', 'Blue', 'Wonderful watch', 20, 'images/defaultproduct.png');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE products;

-- +goose StatementEnd
