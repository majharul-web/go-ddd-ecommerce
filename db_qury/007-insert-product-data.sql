INSERT INTO products (title, description, price, img_url)
VALUES ('Product 1', 'Sample description', 99.99, 'https://example.com/image.jpg')
RETURNING id;
