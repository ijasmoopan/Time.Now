-- +goose Up
CREATE TABLE payments (
    payment_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(user_id) NOT NULL,
    total_price FLOAT NOT NULL,
    payment_type VARCHAR(40) NOT NULL,
    payment_status BOOLEAN DEFAULT FALSE NOT NULL,
    created_at TIMESTAMP DEFAULT NULL,
    paid_at TIMESTAMP DEFAULT NULL,
    cancelled_at TIMESTAMP DEFAULT NULL
);

-- +goose Down
DROP TABLE payments CASCADE;

