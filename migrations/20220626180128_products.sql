-- +goose Up
-- +goose StatementBegin

CREATE TABLE products (
    product_id BIGSERIAL PRIMARY KEY,
    product_name VARCHAR(100) NOT NULL,
    product_price NUMERIC(16, 2) NOT NULL,
    product_desc TEXT NOT NULL,
    product_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    product_updated_at TIMESTAMP DEFAULT NULL,
    product_deleted_at TIMESTAMP DEFAULT NULL,
    -- product_color_id BIGINT REFERENCES colors(color_id) NOT NULL,
    product_image_id BIGINT REFERENCES images(image_id),
    product_brand_id BIGINT REFERENCES brands(brand_id) NOT NULL,
    product_category_id BIGINT REFERENCES categories(category_id) NOT NULL,
    product_subcategory_id BIGINT REFERENCES subcategories(subcategory_id) NOT NULL
    -- product_inventory_id BIGINT REFERENCES inventories(inventory_id) NOT NULL
);

-- INSERT INTO products (product_name, product_price, product_color, product_desc, product_quantity, product_image)
-- VALUES ('Seven Friday', '1199', 'Red', 'Beautiful watch', 10, 'images/defaultproduct.png');

-- INSERT INTO products (product_name, product_price, product_color, product_desc, product_quantity, product_image)
-- VALUES ('Hublot', '899', 'Blue', 'Wonderful watch', 20, 'images/defaultproduct.png');

-- INSERT INTO products (product_name, product_price, product_color, product_desc, product_quantity, product_image)
-- VALUES ('Seven Friday', '1199', 'Red', 'Beautiful watch', 10, 'images/defaultproduct.png');

-- INSERT INTO products (product_name, product_price, product_color, product_desc, product_quantity, product_image)
-- VALUES ('Hublot', '899', 'Blue', 'Wonderful watch', 20, 'images/defaultproduct.png');

-- INSERT INTO products (product_name, product_price, product_color, product_desc, product_quantity, product_image)
-- VALUES ('Seven Friday', '1199', 'Red', 'Beautiful watch', 10, 'images/defaultproduct.png');

-- INSERT INTO products (product_name, product_price, product_color, product_desc, product_quantity, product_image)
-- VALUES ('Hublot', '899', 'Blue', 'Wonderful watch', 20, 'images/defaultproduct.png');

-- INSERT INTO products (product_name, product_price, product_color, product_desc, product_quantity, product_image)
-- VALUES ('Seven Friday', '1199', 'Red', 'Beautiful watch', 10, 'images/defaultproduct.png');

-- INSERT INTO products (product_name, product_price, product_color, product_desc, product_quantity, product_image)
-- VALUES ('Hublot', '899', 'Blue', 'Wonderful watch', 20, 'images/defaultproduct.png');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE products CASCADE;

-- +goose StatementEnd
