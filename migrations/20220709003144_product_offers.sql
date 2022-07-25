-- +goose Up
CREATE TABLE product_offers (id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    offer BIGINT NOT NULL,
    product_id BIGINT REFERENCES products(product_id) NOT NULL,
    offer_from DATE NOT NULL,
    offer_to DATE NOT NULL);

-- +goose Down
DROP TABLE product_offers CASCADE;
