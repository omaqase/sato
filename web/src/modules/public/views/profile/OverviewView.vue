<template>
  <div class="flex-1">
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
    <div class="mt-8 bg-white rounded-2xl shadow-sm overflow-hidden">
      <div class="p-6 border-b border-gray-100">
        <div class="flex items-center justify-between">
          <h2 class="text-lg font-semibold text-gray-900">Последние заказы</h2>
          <router-link
            to="/account/orders"
            class="text-sm font-medium text-blue-600 hover:text-blue-700"
          >
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
              <div :class="`w-2 h-2 rounded-full ${order.statusDot}`" />
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
              <p class="mt-1 text-sm text-gray-500">{{ formatPrice(Number(order.price)) }} ₽</p>
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
</template>

<script setup lang="ts">
import { ref } from 'vue'

interface Stat {
  id: number
  name: string
  value: string
  icon: string
  bgColor: string
  iconColor: string
}

const stats = ref<Stat[]>([
  {
    id: 1,
    name: 'Активные заказы',
    value: '3',
    icon: 'BoxIcon',
    bgColor: 'bg-blue-50',
    iconColor: 'text-blue-600',
  },
  {
    id: 2,
    name: 'Бонусные баллы',
    value: '1 250',
    icon: 'StarIcon',
    bgColor: 'bg-purple-50',
    iconColor: 'text-purple-600',
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
])

interface Order {
  id: number
  number: string
  date: string
  status: string
  statusColor: string
  statusDot: string
  name: string
  price: string
  image: string
}

const recentOrders = ref<Order[]>([
  {
    id: 1,
    number: '12345',
    date: '12 марта 2024',
    status: 'В пути',
    statusColor: 'bg-blue-100 text-blue-700',
    statusDot: 'bg-blue-600',
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
    statusDot: 'bg-green-600',
    name: 'AirPods Pro 2nd generation',
    price: '22 990',
    image: 'https://via.placeholder.com/400',
  },
])

function formatPrice(price: number): string {
  return price.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ' ')
}
</script>
