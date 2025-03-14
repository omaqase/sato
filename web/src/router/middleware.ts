import type { NavigationGuardNext, RouteLocationNormalized } from 'vue-router'
import { useAuthStore } from '@/stores/useAuthStore'
import { useRoleStore } from '@/stores/useRoleStore'

export async function authMiddleware(
  to: RouteLocationNormalized,
  from: RouteLocationNormalized,
  next: NavigationGuardNext,
) {
  const token = localStorage.getItem('access_token')
  
  // Редирект с /security/login если уже авторизован
  if (to.path === '/security/login' && token) {
    next({ path: '/account' })
    return
  }

  if (to.matched.some((record) => record.meta.requiresAuth)) {
    if (!token) {
      next({ path: '/security/login' })
      return
    }

    // Проверяем роль админа для защищенных роутов
    if (to.matched.some((record) => record.meta.requiresAdmin)) {
      const roleStore = useRoleStore()
      await roleStore.checkAdminRole()
      
      if (!roleStore.isAdmin) {
        next({ path: '/' })
        return
      }
    }
  }
  
  next()
}
