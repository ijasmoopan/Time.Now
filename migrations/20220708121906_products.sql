-- +goose Up
CREATE TABLE products (
    product_id BIGSERIAL PRIMARY KEY,
    product_name VARCHAR(100) NOT NULL,
    product_price NUMERIC(16, 2) NOT NULL,
    product_desc TEXT NOT NULL,
    product_status BOOLEAN DEFAULT TRUE NOT NULL,
    product_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    product_updated_at TIMESTAMP DEFAULT NULL,
    product_deleted_at TIMESTAMP DEFAULT NULL,
    product_brand_id BIGINT REFERENCES brands(brand_id) NOT NULL,
    product_category_id BIGINT REFERENCES categories(category_id) NOT NULL,
    product_subcategory_id BIGINT REFERENCES subcategories(subcategory_id) NOT NULL
);

-- +goose Down
DROP TABLE products CASCADE;
