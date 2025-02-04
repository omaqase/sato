package postgres

//const createUserQuery = `
//    INSERT INTO user_management.users (
//        email,
//        password_hash,
//        personal_data,
//        role
//    ) VALUES (
//        $1,
//        $2,
//        jsonb_build_object(
//            'first_name', $3,
//            'last_name', $4,
//            'phone', $5
//        ),
//        COALESCE($6, 'customer')
//    )
//    RETURNING
//        id,
//        email,
//        personal_data->>'first_name' as first_name,
//        personal_data->>'last_name' as last_name,
//        role,
//        created_at,
//        updated_at,
//        is_email_verified,
//        is_active`

const updateUserQuery = `
    UPDATE user_management.users
    SET 
        personal_data = COALESCE(
            personal_data || jsonb_build_object(
                'first_name', COALESCE($1, personal_data->>'first_name'),
                'last_name', COALESCE($2, personal_data->>'last_name'),
                'phone', COALESCE($3, personal_data->>'phone')
            ), 
            personal_data
        ),
        role = COALESCE($4, role),
        is_active = COALESCE($5, is_active),
        updated_at = NOW()
    WHERE 
        id = $6 
        AND is_deleted = false
    RETURNING 
        id,
        email,
        personal_data->>'first_name' as first_name,
        personal_data->>'last_name' as last_name,
        personal_data->>'phone' as phone,
        role,
        is_active,
        created_at,
        updated_at`

const deleteUserQuery = `
    UPDATE user_management.users
    SET 
        is_deleted = true,
        deleted_at = NOW(),
        updated_at = NOW()
    WHERE 
        id = $1 
        AND is_deleted = false
    RETURNING 
        id, 
        deleted_at`

const getUserByCredentialsQuery = `
    SELECT 
        id, 
        email, 
        password_hash,
        role,
        is_active
    FROM user_management.users 
    WHERE 
        email = $1 
        AND is_deleted = false 
        AND is_email_verified = true`

const getUserSubscribedToPromotions = `
    SELECT 
        email,
        personal_data->>'phone' as phone
    FROM user_management.users 
    WHERE 
        personal_data->>'marketing_opt_in' = 'true'
        AND is_deleted = false
        AND email > $1
    ORDER BY email 
    LIMIT $2`

const getUserByEmailSQL = `
	SELECT id, email, role
	FROM user_management.users
	WHERE email = $1 AND deleted_at IS NULL
`

const createUserSQL = `
	INSERT INTO user_management.users(email, first_name, last_name)
	VALUES ($1, $2, $3)
	RETURNING id, email, role
`

const updateUserSQL = `
	UPDATE user_management.users
	SET first_name = $1, last_name = $2, phone = $3, promotions = $4
	WHERE id = $5 AND deleted_at IS NULL
	RETURNING id, first_name, last_name, email, role, phone, promotions
`
