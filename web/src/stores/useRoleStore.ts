import { authApi } from '@/api/auth'
import { defineStore } from 'pinia'

export const useRoleStore = defineStore('role', {
  state: () => ({
    isAdmin: false,
  }),
  
  actions: {
    setAdmin(value: boolean) {
      this.isAdmin = value
    },
    
    // Проверяем роль при инициализации приложения
    async checkAdminRole() {
        const token = localStorage.getItem('access_token')
        if (!token) {
          this.isAdmin = false
          return
        }
        
        try {
          const response = await authApi.checkAdminRole()
          this.isAdmin = response.isAdmin
          return this.isAdmin
        } catch (error) {
          console.error('Error checking admin role:', error)
          this.isAdmin = false
          return false
        }
      },
    },

  persist: true,
}) 