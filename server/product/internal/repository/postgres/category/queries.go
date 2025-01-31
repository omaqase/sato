package category

const createCategoryQuery = `
    INSERT INTO product_management.categories (title, slug, created_at, updated_at)
    VALUES ($1, $2, $3, $4)
    RETURNING id, title, slug, created_at, updated_at`

const getCategoryByIDQuery = `
    SELECT id, title, slug, created_at, updated_at
    FROM product_management.categories
    WHERE id = $1 AND deleted_at IS NULL`

const updateCategoryQuery = `
    UPDATE product_management.categories
    SET title = $1, slug = $2, updated_at = NOW()
    WHERE id = $3 AND deleted_at IS NULL
    RETURNING id, title, slug, created_at, updated_at`

const deleteCategoryQuery = `
    UPDATE product_management.categories
    SET deleted_at = NOW(), updated_at = NOW()
    WHERE id = $1 AND deleted_at IS NULL`

const listCategoriesQuery = `
    SELECT id, title, slug, created_at, updated_at
    FROM product_management.categories
    WHERE deleted_at IS NULL
    ORDER BY created_at DESC
    LIMIT $1 OFFSET $2`
