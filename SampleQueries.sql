------------------------------- Product Quantity ------------------------------------

SELECT P.PRODUCT_ID,
	P.PRODUCT_NAME,
	P.PRODUCT_PRICE,
	P.PRODUCT_DESC,
	INV.PRODUCT_COLOR,
	INV.PRODUCT_QUANTITY,
	P.PRODUCT_IMAGE_ID,
	P.PRODUCT_BRAND_ID,
	P.PRODUCT_CATEGORY_ID,
	P.PRODUCT_SUBCATEGORY_ID,
	I.PRODUCT_IMAGE,
	C.CATEGORY_NAME,
	S.SUBCATEGORY_NAME,
	B.BRAND_NAME
FROM PRODUCTS P
INNER JOIN CATEGORIES C ON P.PRODUCT_CATEGORY_ID = C.CATEGORY_ID
INNER JOIN SUBCATEGORIES S ON P.PRODUCT_SUBCATEGORY_ID = S.SUBCATEGORY_ID
INNER JOIN BRANDS B ON P.PRODUCT_BRAND_ID = B.BRAND_ID
LEFT JOIN IMAGES I ON P.PRODUCT_IMAGE_ID = I.IMAGE_ID
INNER JOIN INVENTORIES INV ON P.PRODUCT_ID = INV.PRODUCT_ID
WHERE P.PRODUCT_DELETED_AT IS NULL;

-------------------------------- Product Inventories ----------------------
SELECT *
FROM INVENTORIES;
SELECT * FROM colors;

SELECT INVENTORY_ID,
	PRODUCT_ID,
	PRODUCT_COLOR,
	PRODUCT_QUANTITY
FROM INVENTORIES;

SELECT * FROM categories;
SELECT * FROM products;
SELECT * FROM inventories;
SELECT * FROM colors;
SELECT * FROM images;
DROP TABLE wishlist CASCADE;

-----------------------------Inserting product ---------------------------
INSERT INTO products (product_name,
					product_price,
					product_desc,
					product_created_at,
					product_brand_id,
					product_category_id,
					product_subcategory_id)
			VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING product_id;
----------------- Inserting product colors ----------------------
INSERT INTO colors (color)
			VALUES ($1) RETURNING color_id;
--------- Inserting product quantity according to color -----------
INSERT INTO inventories (product_id,
					color_id,
					product_quantity,
					inventory_created_at)
			VALUES ($1, $2, $3, $4);
------------------------ Inserting Product Image -------------------------
INSERT INTO images (product_id,
					product_image)
			VALUES ($1, $2) RETURNING image_id;	
--------------------------------------------------------------------------




-------------------------------------------------------------------------
SELECT * FROM products;
SELECT * FROM categories;
SELECT * FROM category_offers;
CREATE TABLE products (id BIGSERIAL PRIMARY KEY,
					   name TEXT NOT NULL,
					   price FLOAT NOT NULL,
					   category_id BIGINT REFERENCES categories(category_id) NOT NULL,
					  offer_price FLOAT);
-- CREATE TABLE categories (id BIGSERIAL PRIMARY KEY,
-- 						name TEXT NOT NULL);
						 
-- 						category_offer_id BIGINT REFERENCES category_offer(id) NOT NULL)
						
CREATE TABLE category_offers (id BIGSERIAL PRIMARY KEY,
							 name TEXT NOT NULL,
							 offer BIGINT NOT NULL,
							 category_id BIGINT REFERENCES categories(category_id) NOT NULL,
							 offer_from DATE NOT NULL,
							 offer_to DATE NOT NULL,
							 offer_status BOOLEAN DEFAULT TRUE NOT NULL);
INSERT INTO category_offers (name, offer, category_id, offer_from, offer_to)
			VALUES ('super', 50, 2, '2022-04-04', '2022-05-05');
					 
------------- Join categories and offers --------------------
SELECT * FROM categories c
	INNER JOIN category_offers co
		ON co.category_id = c.category_id;
		
ALTER TABLE categories ADD COLUMN offer_price FLOAT;
ALTER TABLE categories DROP COLUMN offer_price;

INSERT INTO products (name, price, category_id)
	VALUES ('titan', 35.00, 2);
	
SELECT * FROM category_offers;

UPDATE category_offers SET offer_from = '2022-08-01', offer_to = '2022-08-31' WHERE id = 2;

---------------------- Product with category offer ----------------------
SELECT p.id,
		p.name,
		p.price,
		p.category_id,
		c.category_name,
		o.name,
		o.offer,
		CASE 
			WHEN o.offer_from <= NOW() AND offer_to >= NOW() THEN true
			ELSE false
		END AS offer_status,
		CASE 
			WHEN o.offer_from <= NOW() AND offer_to >= NOW() THEN (p.price - ((p.price * o.offer)/100))
			ELSE NULL
		END AS discount_price 
	FROM products p
	INNER JOIN categories c
		ON p.id = c.category_id
	INNER JOIN category_offers o
		ON o.category_id = c.category_id;
		
		
		
SELECT *, CASE 
		  	WHEN offer_from <= NOW() AND offer_to >= NOW() THEN true 
		  	ELSE false 
		  END AS status
	FROM category_offers;
		
		
------------------------------ Product in Detail Original ---------------------
SELECT p.product_id,
						p.product_name,
						p.product_price,
						p.product_desc,
						p.product_brand_id,
						p.product_category_id,
						p.product_subcategory_id,
						i.image_id,
						i.product_image,
						c.category_name,
						c.category_desc,
						s.subcategory_name,
						s.subcategory_desc,
						b.brand_name,
						b.brand_desc
					FROM   products p
						INNER JOIN categories c
								ON p.product_category_id = c.category_id
						INNER JOIN subcategories s
								ON p.product_subcategory_id = s.subcategory_id
						INNER JOIN brands b
								ON p.product_brand_id = b.brand_id
						LEFT JOIN images i
								ON p.product_id = i.product_id
						INNER JOIN colors r
								ON 
					WHERE  p.product_deleted_at IS NULL 
						AND p.product_status = true 

Ordered
Shipped
In-Transit
Delivered

