<template>
  <div class="min-h-screen bg-gradient-to-b from-gray-50 to-white pb-12">
    <!-- Hero Section -->
    <div class="bg-gradient-to-r from-blue-600 to-blue-700 text-white">
      <div class="max-w-7xl mx-auto px-4">
        <div class="py-12">
          <div class="flex items-center gap-6">
            <div
              class="w-24 h-24 rounded-2xl bg-white/10 backdrop-blur-sm flex items-center justify-center"
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
              <h1 class="text-3xl font-bold mb-1">Александр Иванов</h1>
              <div class="flex items-center gap-4 text-white/80">
                <span>+7 (999) 123-45-67</span>
                <span class="text-white/40">•</span>
                <span>alexander@example.com</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="max-w-7xl mx-auto px-4 -mt-8">
      <div class="flex flex-col lg:flex-row gap-8">
        <!-- Sidebar -->
        <div class="lg:w-80">
          <div class="bg-white rounded-2xl shadow-sm overflow-hidden">
            <!-- Navigation -->
            <nav class="divide-y divide-gray-100">
              <a
                v-for="item in menuItems"
                :key="item.id"
                :class="[
                  'flex items-center gap-3 px-6 py-4 hover:bg-gray-50 transition-colors',
                  item.active ? 'bg-blue-50 text-blue-600 hover:bg-blue-50' : 'text-gray-700',
                ]"
                href="#"
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
              </a>
            </nav>
          </div>

          <!-- Support Card -->
          <div class="mt-6 bg-gradient-to-br from-gray-900 to-gray-800 rounded-2xl p-6 text-white">
            <div class="flex items-center gap-4 mb-4">
              <div class="p-3 bg-white/10 rounded-xl">
                <svg class="w-6 h-6" viewBox="0 0 24 24" fill="none">
                  <path
                    d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
                    stroke="currentColor"
                    stroke-width="2"
                  />
                  <path
                    d="M12 8v4l3 3"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                  />
                </svg>
              </div>
              <div>
                <h3 class="font-medium">Нужна помощь?</h3>
                <p class="text-sm text-white/70">Мы всегда на связи</p>
              </div>
            </div>
            <button
              class="w-full py-3 bg-white/10 hover:bg-white/20 rounded-xl text-sm font-medium transition-colors"
            >
              Написать в поддержку
            </button>
          </div>
        </div>

        <!-- Content -->
        <div class="flex-1 space-y-8">
          <!-- Stats Grid -->
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
            <div
              v-for="stat in stats"
              :key="stat.id"
              class="bg-white rounded-2xl p-6 hover:shadow-lg transition-all duration-300"
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

          <!-- Recent Orders -->
          <div class="bg-white rounded-2xl shadow-sm overflow-hidden">
            <div class="p-6 border-b border-gray-100">
              <div class="flex items-center justify-between">
                <h2 class="text-lg font-semibold text-gray-900">Последние заказы</h2>
                <a href="#" class="text-sm font-medium text-blue-600 hover:text-blue-700">
                  Все заказы
                </a>
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
                    <p
                      class="font-medium text-gray-900 truncate hover:text-blue-600 transition-colors"
                    >
                      {{ order.name }}
                    </p>
                    <p class="mt-1 text-sm text-gray-500">
                      {{ formatPrice(Number(order.price)) }} ₽
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
import { ref } from 'vue'

const menuItems = [
  {
    id: 1,
    name: 'Обзор',
    icon: 'HomeIcon',
    active: true,
    badge: true,
  },
  {
    id: 2,
    name: 'Мои заказы',
    icon: 'ShoppingBagIcon',
    active: false,
    badge: true,
  },
  {
    id: 3,
    name: 'Избранное',
    icon: 'HeartIcon',
    active: false,
    badge: true,
  },
  {
    id: 4,
    name: 'Уведомления',
    icon: 'BellIcon',
    active: false,
    badge: false,
  },
  {
    id: 5,
    name: 'Настройки',
    icon: 'CogIcon',
    active: false,
    badge: true,
  },
]

const stats = [
  {
    id: 1,
    name: 'Активные заказы',
    value: '3',
    icon: 'ShoppingBagIcon',
    bgColor: 'bg-blue-50',
    iconColor: 'text-blue-600',
  },
  {
    id: 2,
    name: 'Бонусные баллы',
    value: '2 450',
    icon: 'GiftIcon',
    bgColor: 'bg-green-50',
    iconColor: 'text-green-600',
  },
  {
    id: 3,
    name: 'Товары в избранном',
    value: '12',
    icon: 'HeartIcon',
    bgColor: 'bg-red-50',
    iconColor: 'text-red-600',
  },
  {
    id: 4,
    name: 'Отзывы',
    value: '8',
    icon: 'ChatIcon',
    bgColor: 'bg-yellow-50',
    iconColor: 'text-yellow-600',
  },
]

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
