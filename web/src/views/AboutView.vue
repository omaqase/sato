<template>
  <div class="min-h-screen bg-gradient-to-b from-gray-50 to-white pb-12">
    <div class="bg-gradient-to-r from-blue-600 to-blue-700 text-white">
      <div class="max-w-7xl mx-auto px-4 py-16">
        <div class="max-w-2xl">
          <h1 class="text-4xl font-bold mb-6">О компании Sato</h1>
          <p class="text-xl text-blue-100">
            Ведущий маркетплейс Казахстана с сотнями довольных покупателей
          </p>
        </div>
      </div>
    </div>

    <div class="max-w-7xl mx-auto px-4 py-12">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8 mb-16">
        <div
          v-for="stat in stats"
          :key="stat.id"
          class="bg-white p-8 rounded-2xl hover:shadow-xl transition-all duration-300"
        >
          <div class="w-16 h-16 rounded-xl bg-blue-100 flex items-center justify-center mb-6">
            <component :is="stat.icon" class="w-8 h-8 text-blue-600" />
          </div>
          <div class="text-3xl font-bold text-gray-900 mb-2">{{ stat.value }}</div>
          <p class="text-gray-600">{{ stat.description }}</p>
        </div>
      </div>

      <div class="bg-white rounded-2xl p-8 mb-16">
        <h2 class="text-2xl font-bold text-gray-900 mb-8">Наша миссия</h2>
        <div class="prose prose-blue max-w-none">
          <p class="text-gray-600">
            Мы создаем удобную и надежную платформу для онлайн-торговли, которая помогает продавцам
            развивать свой бизнес, а покупателям находить лучшие товары по выгодным ценам.
          </p>
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div v-for="section in sections" :key="section.id" class="bg-white rounded-2xl p-8">
          <h2 class="text-2xl font-bold text-gray-900 mb-4">{{ section.title }}</h2>
          <p class="text-gray-600 mb-6">{{ section.description }}</p>
          <ul class="space-y-3">
            <li v-for="item in section.items" :key="item" class="flex items-start gap-3">
              <svg class="w-5 h-5 text-blue-600 mt-1" viewBox="0 0 24 24" fill="none">
                <path
                  d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linecap="round"
                />
              </svg>
              <span class="text-gray-600">{{ item }}</span>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import ChartHistogramIcon from '@/components/icons/ChartHistogramIcon.vue'
import TruckIcon from '@/components/icons/TruckIcon.vue'
import UsersIcon from '@/components/icons/UsersIcon.vue'

// import UsersIcon from '@/components/icons/UsersIcon.vue'
// import BoxesIcon from '@/components/icons/BoxesIcon.vue'
// import TruckIcon from '@/components/icons/TruckIcon.vue'

const stats = [
  {
    id: 1,
    value: '300K+',
    description: 'Активных покупателей ежемесячно',
    icon: UsersIcon,
  },
  {
    id: 2,
    value: '10K+',
    description: 'Товаров в каталоге',
    icon: ChartHistogramIcon,
  },
  {
    id: 3,
    value: '15+',
    description: 'Городов доставки',
    icon: TruckIcon,
  },
]

const sections = [
  {
    id: 1,
    title: 'Для покупателей',
    description: 'Мы делаем покупки простыми и безопасными',
    items: [
      'Гарантия лучшей цены',
      'Проверенные продавцы',
      'Удобная доставка',
      'Безопасная оплата',
      'Простой возврат',
    ],
  },
  {
    id: 2,
    title: 'Для продавцов',
    description: 'Мы помогаем бизнесу расти и развиваться',
    items: [
      'Доступ к миллионной аудитории',
      'Простая интеграция',
      'Маркетинговая поддержка',
      'Логистическая инфраструктура',
      'Аналитика продаж',
    ],
  },
]
</script>
