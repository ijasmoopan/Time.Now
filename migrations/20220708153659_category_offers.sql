-- +goose Up
CREATE TABLE category_offers (id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    offer BIGINT NOT NULL,
    category_id BIGINT REFERENCES categories(category_id) NOT NULL,
    offer_from DATE NOT NULL,
    offer_to DATE NOT NULL);

-- +goose Down
DROP TABLE category_offers CASCADE;
