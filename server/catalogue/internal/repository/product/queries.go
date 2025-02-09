package product

const CreateProductSQL = `
    INSERT INTO product_management.products (title, category_id, seller_id, stock)
    VALUES ($1, $2, $3, $4)
    RETURNING id, title, category_id, seller_id, stock, rating, approved, created_at
`

const UpdateProductSQL = `
    UPDATE product_management.products
    SET title = $1, category_id = $2, seller_id = $3, stock = $4, updated_at = NOW()
    WHERE id = $5 AND deleted_at IS NULL
    RETURNING id, title, category_id, seller_id, stock, rating, approved, updated_at
`

const GetProductByIDSQL = `
    SELECT id, title, category_id, seller_id, stock, rating, approved, created_at, updated_at
    FROM product_management.products
    WHERE id = $1 AND deleted_at IS NULL
`

const DeleteProductByIDSQL = `
    UPDATE product_management.products
    SET deleted_at = NOW()
    WHERE id = $1 AND deleted_at IS NULL
`

const ListProductsSQL = `
    SELECT id, title, category_id, seller_id, stock, rating, approved, created_at, updated_at
    FROM product_management.products
    WHERE deleted_at IS NULL
    ORDER BY created_at DESC
    LIMIT $1 OFFSET $2
`
