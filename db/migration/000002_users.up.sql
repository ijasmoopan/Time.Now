
-- Users
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


-- Address
CREATE DOMAIN phone_number_address AS VARCHAR(10) CHECK(VALUE ~ '^[0-9]{10}$');
CREATE DOMAIN pincode AS VARCHAR(6) CHECK(VALUE ~ '^[0-9]{6}$');

CREATE TABLE address(
    address_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(user_id) NOT NULL,
    address_name VARCHAR(100) NOT NULL,
    address_phone phone_number_address NOT NULL,
    address_pincode pincode NOT NULL,
    address_housename VARCHAR(100) NOT NULL,
    address_streetname VARCHAR(100) NOT NULL,
    address_city VARCHAR(100) NOT NULL,
    address_state VARCHAR(100) NOT NULL,
    address_desc VARCHAR(255),
    address_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    address_updated_at TIMESTAMP DEFAULT NULL,
    address_deleted_at TIMESTAMP DEFAULT NULL
);