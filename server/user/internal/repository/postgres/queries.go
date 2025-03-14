package postgres

const getUserByEmailSQL = `
	SELECT id, email, role
	FROM user_management.users
	WHERE email = $1 AND deleted_at IS NULL
`

const createUserSQL = `
	INSERT INTO user_management.users(email)
	VALUES ($1)
	RETURNING id, email, role
`

const updateUserSQL = `
	UPDATE user_management.users
	SET first_name = $1, last_name = $2, phone = $3, promotions = $4, updated_at = $5
	WHERE id = $6 AND deleted_at IS NULL
	RETURNING id, first_name, last_name, email, role, phone, promotions, updated_at	
`

const getSubscribedToPromotionUsersSQL = `
    SELECT email, COALESCE(first_name, ''), COALESCE(last_name, '')
    FROM user_management.users
    WHERE promotions = true AND deleted_at IS NULL
    LIMIT $1 OFFSET $2
`

const getUserByIDSQL = `
	SELECT id FROM user_management.users WHERE id = $1 AND deleted_at IS NULL
`

const createSellerSQL = `
	INSERT INTO user_management.sellers(user_id, rating)
	VALUES ($1, $2)
`

const addToFavoritesSQL = `
    INSERT INTO user_management.favorites (user_id, product_id)
    VALUES ($1, $2)
    RETURNING id, user_id, product_id, created_at
`

const removeFromFavoritesSQL = `
    DELETE FROM user_management.favorites
    WHERE user_id = $1 AND product_id = $2
`

const getFavoritesSQL = `
    SELECT id, user_id, product_id, created_at
    FROM user_management.favorites
    WHERE user_id = $1
    ORDER BY created_at DESC
    LIMIT $2 OFFSET $3
`

const getFavoritesCountSQL = `
    SELECT COUNT(*)
    FROM user_management.favorites
    WHERE user_id = $1
`
