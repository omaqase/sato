package category

const CreateCategorySQL = `
    INSERT INTO product_management.categories (title, slug)
    VALUES ($1, $2)
    RETURNING id, title, slug, created_at
`

const UpdateCategorySQL = `
    UPDATE product_management.categories
    SET title = $1, slug = $2, updated_at = $3
    WHERE id = $4 AND deleted_at IS NULL
    RETURNING id, title, slug, updated_at
`

const GetCategoryByIDSQL = `
    SELECT id, title, slug, created_at, updated_at
    FROM product_management.categories
    WHERE id = $1 AND deleted_at IS NULL
`

const GetCategoryBySlugSQL = `
    SELECT id, title, slug, created_at, updated_at
    FROM product_management.categories
    WHERE slug = $1 AND deleted_at IS NULL
`

const DeleteCategoryByIDSQL = `
    UPDATE product_management.categories
    SET deleted_at = NOW()
    WHERE id = $1 AND deleted_at IS NULL
`

const DeleteCategoryBySlugSQL = `
    UPDATE product_management.categories
    SET deleted_at = NOW()
    WHERE slug = $1 AND deleted_at IS NULL
`

const ListCategoriesSQL = `
   SELECT id, title, slug, created_at, updated_at
   FROM product_management.categories
   WHERE deleted_at IS NULL
   ORDER BY created_at DESC
   LIMIT $1 OFFSET $2
`

const AddCategoryHierarchySQL = `
    INSERT INTO product_management.categories_hierarchy (ancestor, descendant, depth)
    VALUES ($1, $2, $3)
`

const GetCategoryAncestorsSQL = `
    SELECT c.id, c.title, c.slug, c.created_at, c.updated_at
    FROM product_management.categories c
    INNER JOIN product_management.categories_hierarchy h ON h.ancestor = c.id
    WHERE h.descendant = $1 AND c.deleted_at IS NULL
    ORDER BY h.depth DESC
`

const GetCategoryDescendantsSQL = `
    SELECT c.id, c.title, c.slug, c.created_at, c.updated_at
    FROM product_management.categories c
    INNER JOIN product_management.categories_hierarchy h ON h.descendant = c.id
    WHERE h.ancestor = $1 AND c.deleted_at IS NULL
    ORDER BY h.depth ASC
`
