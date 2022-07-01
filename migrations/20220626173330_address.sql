-- +goose Up
-- +goose StatementBegin

CREATE DOMAIN phone_number AS VARCHAR(10) CHECK(VALUE ~ '^[0-9]{10}$');
CREATE DOMAIN pincode AS VARCHAR(6) CHECK(VALUE ~ '^[0-9]{6}$');

CREATE TABLE address(
    address_id BIGSERIAL PRIMARY KEY,
    address_name VARCHAR(100) NOT NULL,
    address_phone phone_number NOT NULL,
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

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE address CASCADE;
DROP DOMAIN IF EXISTS pincode CASCADE;
DROP DOMAIN IF EXISTS phone_number CASCADE;

-- +goose StatementEnd
