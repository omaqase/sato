export const adminRoutes = [
  {
    path: '/admin',
    name: 'admin',
    component: () => import('@/modules/admin/views/AdminLayout.vue'),
    meta: { 
      requiresAuth: true,
      requiresAdmin: true 
    },
    children: [
      {
        path: 'categories',
        redirect: { name: 'admin-categories' },
      },
      {
        path: 'categories',
        name: 'admin-categories',
        component: () => import('@/modules/admin/views/CategoriesView.vue'),
      },
      {
        path: 'products',
        name: 'admin-products',
        component: () => import('@/modules/admin/views/ProductsView.vue'),
      },
    ],
  },
] 