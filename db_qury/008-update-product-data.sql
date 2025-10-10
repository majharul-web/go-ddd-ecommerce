UPDATE products
SET title = 'Updated Product 1', description = 'Updated description', price = 89.99, img_url = 'https://example.com/updated-image.jpg'
WHERE id = 1
RETURNING *;
