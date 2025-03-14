export const API_CONFIG = {
  BASE_URL: 'http://localhost:8090',
  ENDPOINTS: {
    AUTH: {
      LOGIN: '/api/v1/auth/login',
      REGISTER: '/api/v1/auth/register',
      REFRESH: '/api/v1/auth/refresh',
    },
    CATALOGUE: {
      CATEGORIES: '/api/v1/category',
      PRODUCTS: '/api/v1/product',
    },
    USER: {
      FAVORITES: '/api/v1/user/favorites',
      CART: '/api/v1/user/cart',
      PROFILE: '/api/v1/user/profile',
      ORDERS: '/api/v1/user/orders',
    },
    SELLER: {
      PRODUCTS: '/seller/api/v1/product',
    },
    ADMIN: {
      CATEGORIES: '/admin/api/v1/category',
      PRODUCTS: '/admin/api/v1/product',
    },
  },
  AUTH_HEADER: 'Authorization',
  TOKEN_PREFIX: 'Bearer',
}
