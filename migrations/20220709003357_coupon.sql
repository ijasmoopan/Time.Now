-- +goose Up
CREATE TABLE coupon (id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    code TEXT NOT NULL,
    offer BIGINT NOT NULL,
    offer_from DATE NOT NULL,
    offer_to DATE NOT NULL);

-- +goose Down
DROP TABLE coupon CASCADE;
