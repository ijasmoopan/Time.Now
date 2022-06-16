-- +goose Up
-- +goose StatementBegin

CREATE DOMAIN phone_number AS VARCHAR(10) CHECK(VALUE ~ '^[0-9]{10}$');
CREATE DOMAIN email AS VARCHAR(255) CHECK(VALUE ~ '^\w+@[a-zA-Z_]+?\.[a-zA-Z]{2,3}$');

CREATE TABLE users (
    user_id BIGSERIAL PRIMARY KEY,
    user_name VARCHAR(100) NOT NULL,
    user_password VARCHAR(100) NOT NULL,
    user_phone phone_number NOT NULL,
    user_email email NOT NULL,
    user_status BOOLEAN DEFAULT TRUE,
    user_referral UUID DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    user_image TEXT DEFAULT NULL
);
ALTER TABLE users ADD CONSTRAINT unique_user_name UNIQUE(user_name);
ALTER TABLE users ADD CONSTRAINT unique_email UNIQUE(user_email);

INSERT INTO users (user_name, user_password, user_phone, user_email, user_image)
VALUES ('ijas', '1234', '7034464400', 'ijasmoopan46@gmail.com', 'images/defaultuser.png');

INSERT INTO users (user_name, user_password, user_phone, user_email, user_image)
VALUES ('salman', '1234', '8590385573', 'fasil2401@gmail.com', 'images/defaultuser.png');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE users DROP CONSTRAINT unique_user_name;
ALTER TABLE users DROP CONSTRAINT unique_email;
DROP TABLE users;
DROP DOMAIN IF EXISTS phone_number CASCADE;
DROP DOMAIN IF EXISTS email CASCADE;

-- +goose StatementEnd
