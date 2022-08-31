CREATE TABLE admins(
    admin_id BIGSERIAL PRIMARY KEY,
    admin_name VARCHAR(100) NOT NULL,
    admin_password VARCHAR(100) NOT NULL 
);

INSERT INTO admins (admin_name, admin_password) 
VALUES ('admin', '31323334d41d8cd98f00b204e9800998ecf8427e');
