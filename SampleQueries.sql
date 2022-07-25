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


------------------------------------------------------------------------
-------------------------- 22 - 07 - 2022 -------------------------------
------------------------------------------------------------------------

SELECT * FROM cart;
SELECT * FROM orders;
SELECT * FROM address;
SELECT * FROM products;
SELECT * FROM users;
SELECT * FROM inventories;

ALTER TABLE orders ADD COLUMN cart_id BIGINT REFERENCES cart(cart_id) NOT NULL DEFAULT 3;
ALTER TABLE orders ADD COLUMN payment_status BOOLEAN DEFAULT TRUE NOT NULL;
ALTER TABLE orders ADD COLUMN payment_type VARCHAR(50) DEFAULT 'COD' NOT NULL;
ALTER TABLE orders ADD COLUMN sold_price FLOAT DEFAULT 0 NOT NULL;

ALTER TABLE orders ADD COLUMN ordered_at TIMESTAMP DEFAULT NOW() NOT NULL;
ALTER TABLE orders ADD COLUMN delivered_at TIMESTAMP NULL;
ALTER TABLE orders ADD COLUMN replaced_at TIMESTAMP NULL;
ALTER TABLE orders ADD COLUMN cancelled_at TIMESTAMP NULL;

ALTER TABLE orders DROP COLUMN ordered_at;
ALTER TABLE orders DROP COLUMN delivered_at;
ALTER TABLE orders DROP COLUMN replaced_at;
ALTER TABLE orders DROP COLUMN cancelled_at;
ALTER TABLE cart ADD COLUMN deleted_at TIMESTAMP NULL;

SELECT CASE WHEN updated_at IS NOT NULL THEN updated_at
			WHEN updated_at IS NULL THEN created_at
		END AS status
	FROM cart
	WHERE cart_id = 3
		AND deleted_at IS NULL;

SELECT cart_id FROM cart WHERE cart_id = 3;
SELECT * FROM inventories;
UPDATE products SET product_deleted_at = NULL;

--------------------------------- Get Products --------------------------------------
SELECT p.product_id,
		p.product_name,
		p.product_price,
		p.product_desc,
		p.product_status,
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
		b.brand_desc,
		inv.inventory_id, 
		inv.product_id, 
		col.color_id,
		col.color, 
		inv.product_quantity 
	FROM   products p
		INNER JOIN categories c
				ON p.product_category_id = c.category_id
		INNER JOIN subcategories s
				ON p.product_subcategory_id = s.subcategory_id
		INNER JOIN brands b
				ON p.product_brand_id = b.brand_id
		LEFT JOIN images i
				ON p.product_id = i.product_id
		INNER JOIN inventories inv
				ON inv.product_id = p.product_id
		INNER JOIN colors col
				ON inv.color_id = col.color_id
	WHERE  p.product_deleted_at IS NULL
			AND inv.inventory_deleted_at IS NULL
			AND p.product_name ILIKE '%shock%'
			AND c.category_name ILIKE '%men%'
			AND s.subcategory_name ILIKE '%collections%'
			AND b.brand_name ILIKE '%casio%'
			AND p.product_price >= 25
			AND p.product_price <= 50
			AND p.product_status = true
			AND col.color ILIKE '%black%'
			AND inv.product_quantity <= 10
			 
---------------------------- CTE Query -----------------------
WITH cte_getproducts (product_id, 
					product_name, 
					product_price,
					product_desc, 
					product_status, 
					product_brand_id,
					product_category_id,
					product_subcategory_id,
					image_id,
					product_image,
					category_name,
					category_desc,
					subcategory_name,
					subcategory_desc,
					brand_name,
					brand_desc,
					inventory_id, 
					color_id,
					color, 
					product_quantity,
					  
					offer_id,
					offer_name,
					offer,
					offer_from,
					offer_to,
					offer_status,
					offer_price
					 ) 
	AS (SELECT p.product_id,
				p.product_name,
				p.product_price,
				p.product_desc,
				p.product_status,
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
				b.brand_desc,
				inv.inventory_id, 
				col.color_id,
				col.color, 
				inv.product_quantity,
		
				o.id,
				o.name,
				o.offer,
				o.offer_from,
				o.offer_to,
				CASE 
					WHEN o.offer_from <= NOW() AND o.offer_to >= NOW() THEN true
					ELSE false
				END AS offer_status,
				CASE 
					WHEN o.offer_from <= NOW() AND o.offer_to >= NOW() THEN (p.product_price - ((p.product_price * o.offer)/100))
					ELSE NULL
				END AS offer_price
	FROM   products p
		INNER JOIN categories c
			ON p.product_category_id = c.category_id
		INNER JOIN subcategories s
			ON p.product_subcategory_id = s.subcategory_id
		INNER JOIN brands b
			ON p.product_brand_id = b.brand_id
		LEFT JOIN images i
			ON p.product_id = i.product_id
		INNER JOIN inventories inv
			ON inv.product_id = p.product_id
		INNER JOIN colors col
			ON inv.color_id = col.color_id
		
		LEFT JOIN category_offers o
			ON o.category_id = p.product_category_id
	WHERE  p.product_deleted_at IS NULL
		AND inv.inventory_deleted_at IS NULL)
		
SELECT * 
	FROM cte_getproducts
		ORDER BY product_id
		OFFSET 0
		LIMIT 5
		WHERE category_name ILIKE '%men%'
			OR category_name ILIKE '%women%'
		
		WHERE product_name ILIKE '%shock%'
			OR product_name ILIKE '%diz%'
			AND category_name ILIKE '%men%'
			AND subcategory_name ILIKE '%collections%'
			AND brand_name ILIKE '%casio%'
			AND product_price >= 25
			AND product_price <= 50
			AND product_status = true
			AND color ILIKE '%black%'
			AND product_quantity <= 10;




---------------------------------------------------------------------------

SELECT p.product_id,
		p.product_name,
		p.product_price,
		p.product_desc,
		p.product_status,
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
		b.brand_desc,
		inv.inventory_id, 
		inv.product_id, 
		col.color_id,
		col.color, 
		inv.product_quantity 
	FROM   products p
		INNER JOIN categories c
				ON p.product_category_id = c.category_id
		INNER JOIN subcategories s
				ON p.product_subcategory_id = s.subcategory_id
		INNER JOIN brands b
				ON p.product_brand_id = b.brand_id
		LEFT JOIN images i
				ON p.product_id = i.product_id
		INNER JOIN inventories inv
				ON inv.product_id = p.product_id
		INNER JOIN colors col
				ON inv.color_id = col.color_id
	WHERE  p.product_deleted_at IS NULL
			AND inv.inventory_deleted_at IS NULL
			AND p.product_name ILIKE '%shock%'
			AND c.category_name ILIKE '%men%'
			AND s.subcategory_name ILIKE '%collections%'
			AND b.brand_name ILIKE '%casio%'
			AND p.product_price >= 25
			AND p.product_price <= 50
			AND p.product_status = true
			AND col.color ILIKE '%black%'
			AND inv.product_quantity <= 10

SELECT * FROM category_offers;
SELECT * FROM products;
SELECT * FROM inventories;
SELECT * FROM products 
	INNER JOIN inventories
		ON products.product_id = inventories.inventory_id;

------------------------- GET OFFER PRODUCTS -----------------
SELECT p.product_id,
		p.product_name,
		p.product_price,
		p.product_category_id,
		c.category_name,
		o.id,
		o.name,
		o.offer,
		o.offer_from,
		o.offer_to,
		CASE 
			WHEN o.offer_from <= NOW() AND o.offer_to >= NOW() THEN true
			ELSE false
		END AS offer_status,
		CASE 
			WHEN o.offer_from <= NOW() AND o.offer_to >= NOW() THEN (p.product_price - ((p.product_price * o.offer)/100))::NUMERIC(10, 2)
			ELSE NULL
		END AS discount_price
	FROM products p
	INNER JOIN categories c
		ON p.product_category_id = c.category_id
	INNER JOIN category_offers o
		ON o.category_id = c.category_id;
----------------------- Get Offers ------------------------
	SELECT c.category_id,
		c.category_name,
		c.category_desc,
		o.id,
		o.name,
		o.offer,
		o.offer_from,
		o.offer_to,
		CASE 
			WHEN o.offer_from <= NOW() AND o.offer_to >= NOW() THEN true
			ELSE false
		END AS offer_status
	FROM category_offers o
		INNER JOIN categories c
			ON o.category_id = c.category_id
	WHERE category_deleted_at IS NULL

SELECT * FROM category_offers;


INSERT INTO category_offers (name, offer, category_id, offer_from, offer_to)
	VALUES ('Fresh offer', 50, 4, '2022-07-01', '2022-07-31');

SELECT * FROM categories;

------------------------------------









