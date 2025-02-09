import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/account/favorites',
      name: 'account-favorites',
      component: () => import('../modules/public/views/FavoritesView.vue'),
    },
    {
      path: '/account/cart',
      name: 'account-cart',
      component: () => import('../modules/public/views/CartView.vue'),
    },
    {
      path: '/account',
      name: 'account-profile',
      component: () => import('../modules/public/views/profile/ProfileView.vue'),
    },
    {
      path: '/account/orders',
      name: 'account-orders',
      component: () => import('../modules/public/views/profile/OrdersView.vue'),
    },
    {
      path: '/account/notifications',
      name: 'account-notifications',
      component: () => import('../modules/public/views/profile/NotificationsView.vue'),
    },
    {
      path: '/account/settings',
      name: 'account-settings',
      component: () => import('../modules/public/views/profile/SettingsView.vue'),
    },
    {
      path: '/account/overview',
      name: 'account-overview',
      component: () => import('../modules/public/views/profile/OverviewView.vue'),
    },
    {
      path: '/security/login',
      name: 'security-login',
      component: () => import('../modules/public/components/forms/LoginForm.vue'),
    },
    {
      path: '/privacy-policy',
      name: 'privacy-policy',
      component: () => import('../views/PrivacyPolicyView.vue'),
    },
    {
      path: '/terms',
      name: 'terms',
      component: () => import('@/views/TermsView.vue'),
    },
    {
      path: '/requisites',
      name: 'requisites',
      component: () => import('@/views/RequisitesView.vue'),
    },
    {
      path: '/press',
      name: 'press',
      component: () => import('@/views/PressView.vue'),
    },
    {
      path: '/press/:id',
      name: 'press-article',
      component: () => import('@/views/PressArticleView.vue'),
    },
    {
      path: '/careers',
      name: 'careers',
      component: () => import('@/views/CareersView.vue'),
    },
    {
      path: '/how-to-order',
      name: 'how-to-order',
      component: () => import('@/views/HowToOrderView.vue'),
    },
    {
      path: '/payment-methods',
      name: 'payment-methods',
      component: () => import('@/views/PaymentMethodsView.vue'),
    },
    {
      path: '/delivery',
      name: 'delivery',
      component: () => import('@/views/DeliveryView.vue'),
    },
    {
      path: '/return-policy',
      name: 'return-policy',
      component: () => import('@/views/ReturnPolicyView.vue'),
    },
    {
      path: '/sell',
      name: 'sell',
      component: () => import('@/views/SellOnMarketplaceView.vue'),
    },
    {
      path: '/transport',
      name: 'transport',
      component: () => import('@/views/TransportView.vue'),
    },
    {
      path: '/pickup-point',
      name: 'pickup-point',
      component: () => import('@/views/PickupPointView.vue'),
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('@/views/AboutView.vue'),
    },
    {
      path: '/product/:id',
      name: 'product',
      component: () => import('@/modules/public/views/ProductView.vue'),
    },
  ],
})

export default router
