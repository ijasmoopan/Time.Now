-- Category offer
CREATE TABLE category_offers (id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    offer BIGINT NOT NULL,
    category_id BIGINT REFERENCES categories(category_id) NOT NULL,
    offer_from DATE NOT NULL,
    offer_to DATE NOT NULL);

-- Product offer
CREATE TABLE product_offers (id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    offer BIGINT NOT NULL,
    product_id BIGINT REFERENCES products(product_id) NOT NULL,
    offer_from DATE NOT NULL,
    offer_to DATE NOT NULL);

-- Coupon
CREATE TABLE coupon (id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    code TEXT NOT NULL,
    offer BIGINT NOT NULL,
    offer_from DATE NOT NULL,
    offer_to DATE NOT NULL);
