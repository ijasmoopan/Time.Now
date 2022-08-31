-- +goose Up
-- +goose StatementBegin
CREATE DOMAIN phone_number AS VARCHAR(10) CHECK(VALUE ~ '^[0-9]{10}$');
CREATE DOMAIN email AS VARCHAR(255) CHECK(VALUE ~ '^\w+@[a-zA-Z_]+?\.[a-zA-Z]{2,3}$');
CREATE DOMAIN gender AS CHAR(1) CHECK(VALUE IN ('F', 'M', 'O'));

CREATE TABLE users (
    user_id BIGSERIAL PRIMARY KEY,
    user_firstname VARCHAR(100) NOT NULL,
    user_secondname VARCHAR(100) NOT NULL,
    user_password VARCHAR(100) NOT NULL,
    user_phone phone_number NOT NULL,
    user_email email NOT NULL,
    user_gender gender NOT NULL,
    user_status BOOLEAN DEFAULT TRUE,
    user_referral UUID DEFAULT gen_random_uuid() NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    user_image TEXT DEFAULT NULL
    -- address_id BIGINT REFERENCES address(address_id) NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- ALTER TABLE users DROP CONSTRAINT unique_email;
-- ALTER TABLE users DROP CONSTRAINT unique_phone;
DROP TABLE users CASCADE;
DROP DOMAIN IF EXISTS gender CASCADE; 
DROP DOMAIN IF EXISTS phone_number CASCADE;
DROP DOMAIN IF EXISTS email CASCADE;

-- +goose StatementEnd
