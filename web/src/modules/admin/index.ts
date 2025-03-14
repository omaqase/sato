import { RouteLocationNormalized, NavigationGuardNext } from 'vue-router'
import { useRoleStore } from '@/stores/useRoleStore'

export * from './store'
export * from './types'
export * from './services'
export * from './composables'

export { default as AdminLayout } from './layouts/AdminLayout.vue'
export { default as ProductsView } from './views/ProductsView.vue'
export { default as CategoriesView } from './views/CategoriesView.vue'

export const adminModule = {
  name: 'admin',
  routes: () => import('./routes'),
  store: () => import('./store'),
  guards: {
    requireAdmin: async (to: RouteLocationNormalized, from: RouteLocationNormalized, next: NavigationGuardNext) => {
      const roleStore = useRoleStore()
      if (await roleStore.isAdmin()) {
        next()
      } else {
        next('/login')
      }
    }
  }
}