<template>
  <div class="min-h-screen bg-gray-50">
    <aside class="fixed inset-y-0 left-0 w-64 bg-white border-r border-gray-200">
      <div class="flex items-center gap-3 px-6 h-16 border-b border-gray-200">
        <span class="font-bold text-gray-900">Seller Dashboard</span>
      </div>

      <nav class="p-4 space-y-1">
        <router-link
          v-for="item in menuItems"
          :key="item.path"
          :to="item.path"
          class="flex items-center gap-3 px-3 py-2 text-gray-600 rounded-lg hover:bg-gray-50"
          :class="{ 'bg-blue-50 text-blue-600': isActive(item.path) }"
        >
          <component :is="item.icon" class="w-5 h-5" />
          <span>{{ item.name }}</span>
        </router-link>
      </nav>
    </aside>

    <main class="pl-64">
      <header class="h-16 bg-white border-b border-gray-200 px-8 flex items-center justify-between">
        <h1 class="text-xl font-bold text-gray-900">{{ currentPage }}</h1>

        <div class="flex items-center gap-4">
          <div class="flex items-center gap-2 text-sm">
            <span class="text-gray-500">Баланс:</span>
            <span class="font-medium text-gray-900">{{ formatPrice(balance) }} ₽</span>
          </div>

          <button class="relative group">
            <div
              class="hidden group-hover:block absolute right-0 top-full mt-2 w-48 bg-white rounded-xl shadow-lg border border-gray-100 py-2"
            >
              <router-link
                to="/dashboard/profile"
                class="block px-4 py-2 text-sm text-gray-600 hover:bg-gray-50"
              >
                Профиль
              </router-link>
              <router-link
                to="/dashboard/settings"
                class="block px-4 py-2 text-sm text-gray-600 hover:bg-gray-50"
              >
                Настройки
              </router-link>
              <button
                @click="logout"
                class="w-full text-left px-4 py-2 text-sm text-red-600 hover:bg-gray-50"
              >
                Выйти
              </button>
            </div>
          </button>
        </div>
      </header>

      <!-- Page Content -->
      <div class="p-8">
        <router-view />
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import BoxIcon from '@/components/icons/BoxIcon.vue'
import ChartHistogramIcon from '@/components/icons/ChartHistogramIcon.vue'
import CoinsIcon from '@/components/icons/CoinsIcon.vue'
import HomeIcon from '@/components/icons/HomeIcon.vue'
import SettingsIcon from '@/components/icons/SettingsIcon.vue'
import ShoppingCartIcon from '@/components/icons/ShoppingCartIcon.vue'
import UserExperienceIcon from '@/components/icons/UserExperienceIcon.vue'
import { computed, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()
const balance = ref(125000)

const menuItems = [
  { name: 'Обзор', path: '/dashboard', icon: HomeIcon },
  { name: 'Товары', path: '/dashboard/products', icon: BoxIcon },
  { name: 'Заказы', path: '/dashboard/orders', icon: ShoppingCartIcon },
  { name: 'Аналитика', path: '/dashboard/analytics', icon: ChartHistogramIcon },
  { name: 'Финансы', path: '/dashboard/finances', icon: CoinsIcon },
  { name: 'Клиенты', path: '/dashboard/customers', icon: UserExperienceIcon },
  { name: 'Настройки', path: '/dashboard/settings', icon: SettingsIcon },
]

const currentPage = computed(() => {
  return menuItems.find((item) => isActive(item.path))?.name || ''
})

function isActive(path: string): boolean {
  return route.path === path
}

function formatPrice(price: number): string {
  return new Intl.NumberFormat('ru-RU').format(price)
}

function logout() {
  // Логика выхода
  router.push('/login')
}
</script>
