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
