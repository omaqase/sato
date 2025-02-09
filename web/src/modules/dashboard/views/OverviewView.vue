<template>
  <div class="space-y-8">
    <!-- Stats Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <div
        v-for="stat in stats"
        :key="stat.title"
        class="bg-white rounded-xl p-6 border border-gray-100"
      >
        <div class="flex items-center justify-between mb-2">
          <span class="text-sm text-gray-500">{{ stat.title }}</span>
          <component :is="stat.icon" class="w-6 h-6 text-gray-400" />
        </div>
        <div class="space-y-2">
          <div class="text-2xl font-bold text-gray-900">{{ stat.value }}</div>
          <div
            class="flex items-center gap-1 text-sm"
            :class="stat.trend > 0 ? 'text-green-500' : 'text-red-500'"
          >
            <span>{{ Math.abs(stat.trend) }}%</span>
            <svg
              class="w-4 h-4"
              :class="stat.trend > 0 ? 'rotate-0' : 'rotate-180'"
              viewBox="0 0 24 24"
              fill="none"
            >
              <path
                d="M12 5l7 7-7 7"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
              />
            </svg>
            <span class="text-gray-500">vs прошлый месяц</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Recent Orders -->
    <div class="bg-white rounded-xl border border-gray-100">
      <div class="p-6 border-b border-gray-100">
        <div class="flex items-center justify-between">
          <h2 class="text-lg font-semibold text-gray-900">Последние заказы</h2>
          <router-link to="/dashboard/orders" class="text-sm text-blue-600 hover:text-blue-700">
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
          <div class="flex items-center justify-between mb-2">
            <div class="flex items-center gap-4">
              <div :class="`w-2 h-2 rounded-full ${order.statusDot}`"></div>
              <div>
                <div class="font-medium text-gray-900">Заказ #{{ order.number }}</div>
                <div class="text-sm text-gray-500">{{ order.date }}</div>
              </div>
            </div>
            <div :class="`px-3 py-1 rounded-full text-sm font-medium ${order.statusClass}`">
              {{ order.status }}
            </div>
          </div>
          <div class="flex items-center justify-between mt-4">
            <div class="text-sm text-gray-600">{{ order.items }} товаров</div>
            <div class="font-medium text-gray-900">{{ formatPrice(order.total) }} ₽</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import ChartHistogramIcon from '@/components/icons/ChartHistogramIcon.vue'
import CoinsIcon from '@/components/icons/CoinsIcon.vue'
import ShoppingCartIcon from '@/components/icons/ShoppingCartIcon.vue'
import UserExperienceIcon from '@/components/icons/UserExperienceIcon.vue'

const stats = [
  {
    title: 'Продажи',
    value: '346 842 ₽',
    trend: 12.5,
    icon: CoinsIcon,
  },
  {
    title: 'Заказы',
    value: '156',
    trend: 8.2,
    icon: ShoppingCartIcon,
  },
  {
    title: 'Клиенты',
    value: '89',
    trend: -3.1,
    icon: UserExperienceIcon,
  },
  {
    title: 'Конверсия',
    value: '3.2%',
    trend: 4.3,
    icon: ChartHistogramIcon,
  },
]

const recentOrders = [
  {
    id: 1,
    number: '12345',
    date: '12.03.2024',
    status: 'Новый',
    statusDot: 'bg-blue-500',
    statusClass: 'bg-blue-50 text-blue-600',
    items: 3,
    total: 12990,
  },
  {
    id: 2,
    number: '12344',
    date: '12.03.2024',
    status: 'В доставке',
    statusDot: 'bg-yellow-500',
    statusClass: 'bg-yellow-50 text-yellow-600',
    items: 1,
    total: 2990,
  },
  {
    id: 3,
    number: '12343',
    date: '11.03.2024',
    status: 'Выполнен',
    statusDot: 'bg-green-500',
    statusClass: 'bg-green-50 text-green-600',
    items: 2,
    total: 5980,
  },
]

function formatPrice(price: number): string {
  return new Intl.NumberFormat('ru-RU').format(price)
}
</script>
