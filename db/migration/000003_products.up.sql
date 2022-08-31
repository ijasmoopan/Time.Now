-- Images
CREATE TABLE images(
    image_id BIGSERIAL PRIMARY KEY,
    product_id INT,
    product_image TEXT
);

-- Categories
CREATE TABLE categories (
    category_id BIGSERIAL PRIMARY KEY,
    category_name VARCHAR(100) NOT NULL,
    category_desc VARCHAR(255),
    category_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    category_updated_at TIMESTAMP DEFAULT NULL,
    category_deleted_at TIMESTAMP DEFAULT NULL
);

-- Subcategories
CREATE TABLE subcategories(
    subcategory_id BIGSERIAL PRIMARY KEY,
    subcategory_name VARCHAR(100) NOT NULL,
    subcategory_desc VARCHAR(255),
    subcategory_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    subcategory_updated_at TIMESTAMP DEFAULT NULL,
    subcategory_deleted_at TIMESTAMP DEFAULT NULL
);

-- Brands
CREATE TABLE brands(
    brand_id BIGSERIAL PRIMARY KEY,
    brand_name VARCHAR(100) NOT NULL,
    brand_desc VARCHAR(255),
    brand_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    brand_updated_at TIMESTAMP DEFAULT NULL,
    brand_deleted_at TIMESTAMP DEFAULT NULL
);

-- Products
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

-- Colors
CREATE TABLE colors(
    color_id BIGSERIAL PRIMARY KEY,
    color VARCHAR(100) NOT NULL,
    color_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    color_updated_at TIMESTAMP DEFAULT NULL,
    color_deleted_at TIMESTAMP DEFAULT NULL
);

-- Inventories
CREATE TABLE inventories(
    inventory_id BIGSERIAL PRIMARY KEY,
    product_id BIGINT REFERENCES products(product_id) NOT NULL,
    color_id BIGINT REFERENCES colors(color_id) NOT NULL,
    product_quantity BIGINT NOT NULL,
    inventory_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    inventory_updated_at TIMESTAMP DEFAULT NULL,
    inventory_deleted_at TIMESTAMP DEFAULT NULL
);