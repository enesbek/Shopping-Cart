package constants

const CREATE_PRODUCT_TABLE_QUERY = `
DROP TABLE IF EXISTS products;
CREATE TABLE products
(
	Id SERIAL PRIMARY KEY,
	name TEXT,
	price INTEGER,
	image TEXT,
	quantity INTEGER,
	basket BOOLEAN
);
INSERT INTO products(name, price, image, quantity, basket)
VALUES 
('iphone', 1000, 'https://cdn.vatanbilgisayar.com/Upload/PRODUCT/apple/thumb/ip-3_large.jpg', 1, false),
('samsung', 900, 'https://cdn.vatanbilgisayar.com/Upload/PRODUCT/samsung/thumb/117735-ana_large.jpg', 1, false),
('huawei', 750, 'https://cdn.vatanbilgisayar.com/Upload/PRODUCT/huawei/thumb/huawei-p40-gorseli-106926_large.jpg', 1, false),
('iphone 8', 500, 'https://cdn.vatanbilgisayar.com/Upload/PRODUCT/apple/thumb/v2-88458-3_large.jpg', 1, false);`

const GET_ALL_PRODUCTS_QUERY = `
SELECT * FROM products
ORDER BY id ASC;`

const ADD_PRODUCT_TO_BASKET = `
UPDATE products
SET basket = true
WHERE id = $1;`

const GET_ALL_BASKET = `
SELECT * FROM products
WHERE basket = true
ORDER BY id ASC;`

const INCREMENT_QUANTITY_OF_PRODUCT = `
UPDATE products
SET quantity = quantity + 1
WHERE id = $1;`

const DECREMENT_QUANTITY_OF_PRODUCT = `
UPDATE products
SET quantity = quantity - 1
WHERE id = $1;`

const DELETE_PRODUCT_FROM_BASKET = `
UPDATE products
SET basket = false, quantity = 1
WHERE id = $1;`

const GET_QUANTITY_OF_PRODUCT = `
SELECT quantity FROM products
WHERE id = $1;`

