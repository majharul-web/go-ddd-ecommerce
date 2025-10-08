UPDATE users
SET
    first_name = 'Alicia',
    last_name = 'Brown',
    email = 'alicia@example.com',
    password = 'newpassword',
    avatar = 'avatar2.png',
    is_shop_owner = true,
    
WHERE id = 1
RETURNING *;
