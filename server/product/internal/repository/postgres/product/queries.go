package product

const createProductQuery = `
    INSERT INTO product_management.products (title, description, price, category_id, created_at, updated_at)
    VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING id, title, description, price, category_id, created_at, updated_at`

const getProductByIDQuery = `
    SELECT id, title, description, price, category_id, created_at, updated_at
    FROM product_management.products
    WHERE id = $1 AND deleted_at IS NULL`

const updateProductQuery = `
    UPDATE product_management.products
    SET title = $1, description = $2, price = $3, category_id = $4, updated_at = NOW()
    WHERE id = $5 AND deleted_at IS NULL
    RETURNING id, title, description, price, category_id, created_at, updated_at`

const deleteProductQuery = `
    UPDATE product_management.products
    SET deleted_at = NOW(), updated_at = NOW()
    WHERE id = $1 AND deleted_at IS NULL`

const listProductsQuery = `
    SELECT id, title, description, price, category_id, created_at, updated_at
    FROM product_management.products
    WHERE deleted_at IS NULL
    ORDER BY created_at DESC
    LIMIT $1 OFFSET $2`
