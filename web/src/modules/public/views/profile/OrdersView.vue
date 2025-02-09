<template>
  <div class="flex-1">
    <!-- Filters -->
    <div class="bg-white rounded-2xl shadow-sm p-6 mb-6">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-4">
          <div v-for="filter in filters" :key="filter.id" class="relative">
            <button
              @click="filter.active = !filter.active"
              :class="[
                'px-4 py-2 text-sm font-medium rounded-xl transition-colors',
                filter.active
                  ? 'bg-blue-600 text-white'
                  : 'bg-gray-100 text-gray-700 hover:bg-gray-200',
              ]"
            >
              {{ filter.name }}
            </button>
          </div>
        </div>
        <div class="flex items-center gap-2 text-sm text-gray-500">
          <span>Сортировать:</span>
          <select class="px-3 py-1.5 rounded-lg bg-gray-100 border-0 text-gray-700">
            <option value="date">По дате</option>
            <option value="price">По цене</option>
            <option value="status">По статусу</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Orders List -->
    <div class="bg-white rounded-2xl shadow-sm overflow-hidden">
      <div class="divide-y divide-gray-100">
        <div v-for="order in orders" :key="order.id" class="p-6 hover:bg-gray-50 transition-colors">
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
              <p class="mt-1 text-sm text-gray-500">{{ formatPrice(order.price) }} ₽</p>
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

interface Filter {
  id: number
  name: string
  active: boolean
}

const filters = ref<Filter[]>([
  { id: 1, name: 'Все заказы', active: true },
  { id: 2, name: 'В пути', active: false },
  { id: 3, name: 'Доставлены', active: false },
  { id: 4, name: 'Отменены', active: false },
])

interface Order {
  id: number
  number: string
  date: string
  status: string
  statusColor: string
  statusDot: string
  name: string
  price: number
  image: string
}

const orders = ref<Order[]>([
  {
    id: 1,
    number: '12345',
    date: '12 марта 2024',
    status: 'В пути',
    statusColor: 'bg-blue-100 text-blue-700',
    statusDot: 'bg-blue-600',
    name: 'iPhone 13 Pro 256GB',
    price: 89990,
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
    price: 22990,
    image: 'https://via.placeholder.com/400',
  },
])

function formatPrice(price: number): string {
  return price.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ' ')
}
</script>
