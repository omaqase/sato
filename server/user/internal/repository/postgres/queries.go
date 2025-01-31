package postgres

const createUserQuery = `
    INSERT INTO user_management.users (email, password, first_name, last_name, created_at, updated_at)
    VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING id, email, first_name, last_name, created_at, updated_at, is_subscribed_to_promotions`

const updateUserQuery = `
    UPDATE user_management.users
    SET 
        first_name = $1,
        last_name = $2,
        phone = $3,
        is_subscribed_to_promotions = $4,
        updated_at = NOW()
    WHERE 
        id = $5 
        AND deleted_at IS NULL
    RETURNING 
        id, email, first_name, last_name, phone, 
        created_at, updated_at, is_subscribed_to_promotions`

const deleteUserQuery = `
    UPDATE user_management.users
    SET 
        deleted_at = NOW(),
        updated_at = NOW()
    WHERE 
        id = $1 
        AND deleted_at IS NULL
    RETURNING 
        id, deleted_at`

const getUserByCredentialsQuery = `
    SELECT id, email, password 
    FROM user_management.users 
    WHERE email = $1 AND deleted_at IS NULL`

const getUserSubscribedToPromotions = `
    SELECT email 
    FROM user_management.users 
    WHERE is_subscribed_to_promotions = TRUE 
    AND deleted_at IS NULL 
    AND email > $1
    ORDER BY email 
    LIMIT $2`
