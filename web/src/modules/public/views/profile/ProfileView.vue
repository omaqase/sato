<template>
  <div class="min-h-screen bg-gradient-to-b from-gray-50 via-blue-50/10 to-white pb-12">
    <!-- Hero Section -->
    <div class="bg-gradient-to-r from-blue-600 via-blue-500 to-blue-700 text-white shadow-lg relative overflow-hidden">
      <div class="absolute inset-0 bg-[url('data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjAwIiBoZWlnaHQ9IjIwMCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48ZGVmcz48cGF0dGVybiBpZD0iYSIgcGF0dGVyblVuaXRzPSJ1c2VyU3BhY2VPblVzZSIgd2lkdGg9IjQwIiBoZWlnaHQ9IjQwIiBwYXR0ZXJuVHJhbnNmb3JtPSJyb3RhdGUoNDUpIj48cGF0aCBkPSJNLTEwIDMwaDYwdi0ySDEwdjJaIiBmaWxsPSIjZmZmIiBmaWxsLW9wYWNpdHk9Ii4xIi8+PC9wYXR0ZXJuPjwvZGVmcz48cmVjdCB3aWR0aD0iMTAwJSIgaGVpZ2h0PSIxMDAlIiBmaWxsPSJ1cmwoI2EpIi8+PC9zdmc+')] opacity-10"/>
      <div class="max-w-7xl mx-auto px-4">
        <div class="py-12">
          <div class="flex items-center gap-6">
            <div
              class="w-24 h-24 rounded-2xl bg-white/10 backdrop-blur-sm flex items-center justify-center transform hover:scale-105 transition-all duration-300 shadow-xl ring-1 ring-white/20"
            >
              <svg class="w-12 h-12 text-white/90" viewBox="0 0 24 24" fill="none">
                <path
                  d="M20 21v-2a4 4 0 00-4-4H8a4 4 0 00-4 4v2M12 11a4 4 0 100-8 4 4 0 000 8z"
                  stroke="currentColor"
                  stroke-width="2"
                />
              </svg>
            </div>
            <div>
              <!-- Use real user data for name and contact info -->
              <h1 class="text-3xl font-bold mb-1">{{ user?.first_name }} {{ user?.last_name }}</h1>
              <div class="flex items-center gap-4 text-white/80">
                <span>{{ user?.phone || 'Телефон не указан' }}</span>
                <span class="text-white/40">•</span>
                <span>{{ user?.email }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="max-w-7xl mx-auto px-4 -mt-8">
      <div class="flex flex-col lg:flex-row gap-8">
        <!-- Sidebar (unchanged) -->
        <div class="lg:w-80">
          <div class="bg-white rounded-2xl shadow-sm overflow-hidden hover:shadow-lg transition-all duration-300 border border-gray-100">
            <nav class="divide-y divide-gray-100">
              <router-link
                v-for="item in menuItems"
                :key="item.id"
                :to="item.route"
                :class="[
                  'flex items-center gap-3 px-6 py-4 hover:bg-gray-50 transition-colors',
                  item.active ? 'bg-blue-50 text-blue-600 hover:bg-blue-50' : 'text-gray-700',
                ]"
                @click="item.inDevelopment && showFeatureInDevelopment()"
              >
                <component :is="item.icon" class="w-5 h-5" />
                <span class="font-medium">{{ item.name }}</span>
                <svg
                  v-if="item.badge"
                  class="w-5 h-5 ml-auto text-blue-600"
                  viewBox="0 0 24 24"
                  fill="none"
                >
                  <circle cx="12" cy="12" r="4" fill="currentColor" />
                </svg>
              </router-link>
            </nav>
          </div>
          <div class="mt-6 bg-gradient-to-br from-gray-900 via-gray-800 to-gray-900 rounded-2xl p-6 text-white shadow-lg transform hover:-translate-y-1 transition-all duration-300 relative overflow-hidden">
            <div class="absolute inset-0 bg-[url('data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjAwIiBoZWlnaHQ9IjIwMCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48ZGVmcz48cGF0dGVybiBpZD0iYSIgcGF0dGVyblVuaXRzPSJ1c2VyU3BhY2VPblVzZSIgd2lkdGg9IjQwIiBoZWlnaHQ9IjQwIiBwYXR0ZXJuVHJhbnNmb3JtPSJyb3RhdGUoNDUpIj48cGF0aCBkPSJNLTEwIDMwaDYwdi0ySDEwdjJaIiBmaWxsPSIjZmZmIiBmaWxsLW9wYWNpdHk9Ii4wNSIvPjwvcGF0dGVybj48L2RlZnM+PHJlY3Qgd2lkdGg9IjEwMCUiIGhlaWdodD0iMTAwJSIgZmlsbD0idXJsKCNhKSIvPjwvc3ZnPg==')] opacity-50"/>
            <div class="flex items-center gap-4 mb-4">
              <div class="p-3 bg-white/10 rounded-xl">
                <svg class="w-6 h-6" viewBox="0 0 24 24" fill="none">
                  <path d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z" stroke="currentColor" stroke-width="2" />
                  <path d="M12 8v4l3 3" stroke="currentColor" stroke-width="2" stroke-linecap="round" />
                </svg>
              </div>
              <div>
                <h3 class="font-medium">Нужна помощь?</h3>
                <p class="text-sm text-white/70">Мы всегда на связи</p>
              </div>
            </div>
            <button
              @click="showFeatureInDevelopment"
              class="w-full py-3 bg-white/10 hover:bg-white/20 rounded-xl text-sm font-medium transition-colors"
            >
              Написать в поддержку
            </button>
          </div>
        </div>

        <!-- Content -->
        <div class="flex-1 space-y-8">
          <!-- User Info (unchanged, already uses real data) -->
          <div class="bg-white rounded-2xl p-6 shadow-sm hover:shadow-lg transition-all duration-300 border border-gray-100">
            <h2 class="text-2xl font-bold text-gray-900 mb-4">Профиль</h2>
            <div class="space-y-4">
              <div>
                <p class="text-sm text-gray-500">Email</p>
                <p class="font-medium">{{ user?.email }}</p>
              </div>
              <div>
                <p class="text-sm text-gray-500">Имя</p>
                <p class="font-medium">{{ user?.first_name || 'Не указано' }}</p>
              </div>
              <div>
                <p class="text-sm text-gray-500">Фамилия</p>
                <p class="font-medium">{{ user?.last_name || 'Не указана' }}</p>
              </div>
            </div>
          </div>

          <!-- Stats Grid -->
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
            <div
              v-for="stat in stats"
              :key="stat.id"
              class="bg-white rounded-2xl p-6 shadow-sm hover:shadow-lg transition-all duration-300 border border-gray-100 transform hover:-translate-y-1"
            >
              <div class="flex items-center gap-4">
                <div :class="`p-4 rounded-xl ${stat.bgColor}`">
                  <component :is="stat.icon" class="w-6 h-6" :class="stat.iconColor" />
                </div>
                <div>
                  <p class="text-sm font-medium text-gray-500">{{ stat.name }}</p>
                  <p class="mt-1 text-2xl font-bold text-gray-900">{{ stat.value }}</p>
                </div>
              </div>
            </div>
          </div>

          <!-- Recent Orders (unchanged for now, see notes below) -->
          <div class="bg-white rounded-2xl shadow-sm overflow-hidden hover:shadow-lg transition-all duration-300 border border-gray-100">
            <div class="p-6 border-b border-gray-100">
              <div class="flex items-center justify-between">
                <h2 class="text-lg font-semibold text-gray-900">Последние заказы</h2>
                <router-link href="#" class="text-sm font-medium text-blue-600 hover:text-blue-700">
                  Все заказы
                </router-link>
              </div>
            </div>
            <div class="divide-y divide-gray-100">
              <div
                v-for="order in recentOrders"
                :key="order.id"
                class="p-6 hover:bg-gray-50 transition-colors"
              >
                <div class="flex items-center justify-between mb-4">
                  <div class="flex items-center gap-4">
                    <div :class="`w-2 h-2 rounded-full ${order.statusColor}`" />
                    <div>
                      <div class="font-medium text-gray-900">Заказ #{{ order.number }}</div>
                      <div class="text-sm text-gray-500">{{ order.date }}</div>
                    </div>
                  </div>
                  <div :class="`px-4 py-1.5 rounded-full text-xs font-medium ${order.statusColor}`">
                    {{ order.status }}
                  </div>
                </div>
                <div class="flex items-center gap-4">
                  <div class="relative group">
                    <img
                      :src="order.image"
                      :alt="order.name"
                      class="w-20 h-20 rounded-xl object-cover group-hover:scale-105 transition-transform duration-300"
                    />
                    <div class="absolute inset-0 rounded-xl ring-1 ring-inset ring-black/5" />
                  </div>
                  <div class="flex-1 min-w-0">
                    <p class="font-medium text-gray-900 truncate hover:text-blue-600 transition-colors">
                      {{ order.name }}
                    </p>
                    <p class="mt-1 text-sm text-gray-500">
                      {{ formatPrice(Number(order.price)) }} ₸
                    </p>
                  </div>
                  <button
                    class="flex items-center gap-2 px-4 py-2 rounded-xl border border-gray-200 text-sm font-medium text-gray-700 hover:bg-gray-50 transition-colors"
                  >
                    <span>Подробнее</span>
                    <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none">
                      <path
                        d="M9 5l7 7-7 7"
                        stroke="currentColor"
                        stroke-width="2"
                        stroke-linecap="round"
                      />
                    </svg>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useAuthStore } from '@/stores/useAuthStore'
import { useFavoritesStore } from '@/stores/useFavoritesStore'
import { storeToRefs } from 'pinia'
import { useRouter } from 'vue-router'
import { useToast } from '@/composables/useToast'

const authStore = useAuthStore()
const { user } = storeToRefs(authStore)
const router = useRouter()
const { showFeatureInDevelopment } = useToast()
const favoritesStore = useFavoritesStore()

// Fetch favorites data when the component mounts
onMounted(() => {
  favoritesStore.fetchFavorites()
})

const menuItems = [
  {
    id: 1,
    name: 'Обзор',
    icon: 'HomeIcon',
    active: true,
    badge: true,
    route: '/account/overview',
    inDevelopment: false,
  },
  {
    id: 2,
    name: 'Мои заказы',
    icon: 'ShoppingBagIcon',
    active: false,
    badge: true,
    route: '/account/orders',
    inDevelopment: false,
  },
  {
    id: 3,
    name: 'Избранное',
    icon: 'HeartIcon',
    active: false,
    badge: true,
    route: '/favorites',
    inDevelopment: true,
  },
  {
    id: 4,
    name: 'Уведомления',
    icon: 'BellIcon',
    active: false,
    badge: false,
    route: '/account/notifications',
    inDevelopment: true,
  },
  {
    id: 5,
    name: 'Настройки',
    icon: 'CogIcon',
    active: false,
    badge: true,
    route: '/account/settings',
    inDevelopment: false,
  },
]

// Define stats as a computed property for reactivity
const stats = computed(() => [
  {
    id: 1,
    name: 'Имя',
    value: user.value?.first_name || 'Не указано',
    icon: 'UserIcon',
    bgColor: 'bg-blue-50',
    iconColor: 'text-blue-600',
  },
  {
    id: 2,
    name: 'Бонусные баллы',
    value: user.value?.bonus_points?.toString() || '0',
    icon: 'GiftIcon',
    bgColor: 'bg-green-50',
    iconColor: 'text-green-600',
  },
  {
    id: 3,
    name: 'Товары в избранном',
    value: favoritesStore.favorites.length.toString(),
    icon: 'HeartIcon',
    bgColor: 'bg-red-50',
    iconColor: 'text-red-600',
  },
  {
    id: 4,
    name: 'Отзывы',
    value: user.value?.review_count?.toString() || '0',
    icon: 'ChatIcon',
    bgColor: 'bg-yellow-50',
    iconColor: 'text-yellow-600',
  },
])

// Placeholder for recent orders (still mock data, see notes)
const recentOrders = [
  {
    id: 1,
    number: '12345',
    date: '12 марта 2024',
    status: 'В пути',
    statusColor: 'bg-blue-100 text-blue-700',
    name: 'iPhone 13 Pro 256GB',
    price: '89 990',
    image: 'https://via.placeholder.com/400',
  },
  {
    id: 2,
    number: '12344',
    date: '10 марта 2024',
    status: 'Доставлен',
    statusColor: 'bg-green-100 text-green-700',
    name: 'AirPods Pro 2nd generation',
    price: '22 990',
    image: 'https://via.placeholder.com/400',
  },
]

function formatPrice(price: number): string {
  return price.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ' ')
}
</script>