-- +goose Up
-- +goose StatementBegin
CREATE TABLE admins(
    admin_id BIGSERIAL PRIMARY KEY,
    admin_name VARCHAR(100) NOT NULL,
    admin_password VARCHAR(100) NOT NULL 
);
INSERT INTO admins (admin_name, admin_password) 
VALUES ('admin', '1234');
INSERT INTO admins (admin_name, admin_password)
VALUES ('ijas', '1234');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE admins;
-- +goose StatementEnd
