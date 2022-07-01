-- +goose Up
-- +goose StatementBegin

CREATE DOMAIN email AS VARCHAR(255) CHECK(VALUE ~ '^\w+@[a-zA-Z_]+?\.[a-zA-Z]{2,3}$');
CREATE DOMAIN gender AS CHAR(1) CHECK(VALUE IN ('F', 'M', 'O'));

CREATE TABLE users (
    user_id BIGSERIAL PRIMARY KEY,
    user_firstname VARCHAR(100) NOT NULL,
    user_secondname VARCHAR(100) NOT NULL,
    user_password VARCHAR(100) NOT NULL,
    user_phone phone_number NOT NULL,
    user_email email NOT NULL UNIQUE,
    user_gender gender NOT NULL,
    user_status BOOLEAN DEFAULT TRUE,
    user_referral UUID DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    user_image TEXT DEFAULT NULL,
    address_id BIGINT REFERENCES address(address_id)
);

ALTER TABLE users ADD CONSTRAINT unique_email UNIQUE(user_email);
ALTER TABLE users ADD CONSTRAINT unique_phone UNIQUE(user_phone);

INSERT INTO users (user_firstname, user_secondname, user_password, user_phone, user_email, user_gender, user_image)
VALUES ('ijas', 'moopan', '31323334d41d8cd98f00b204e9800998ecf8427e', '7034464400', 'ijasmoopan46@gmail.com', 'M', 'images/defaultuser.png');

INSERT INTO users (user_firstname, user_secondname, user_password, user_phone, user_email, user_gender, user_image)
VALUES ('salman', 'butter', '31323334d41d8cd98f00b204e9800998ecf8427e', '8590385573', 'fasil2401@gmail.com', 'M', 'images/defaultuser.png');

INSERT INTO users (user_firstname, user_secondname, user_password, user_phone, user_email, user_gender, user_image)
VALUES ('rishal', 'p', '31323334d41d8cd98f00b204e9800998ecf8427e', '8136860910', 'muhammedrishalpaangil@gmail.com', 'M', 'images/defaultuser.png');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE users DROP CONSTRAINT unique_email;
ALTER TABLE users DROP CONSTRAINT unique_phone;
DROP TABLE users CASCADE;
DROP DOMAIN IF EXISTS gender CASCADE; 
DROP DOMAIN IF EXISTS email CASCADE;

-- +goose StatementEnd
