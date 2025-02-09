<template>
  <div class="min-h-screen bg-gradient-to-b from-gray-50 to-white pb-12">
    <div class="bg-gradient-to-r from-blue-600 to-blue-700 text-white">
      <div class="max-w-7xl mx-auto px-4 py-16">
        <div class="max-w-2xl">
          <h1 class="text-4xl font-bold mb-6">Продавайте на sato</h1>
          <p class="text-xl text-blue-100">
            Откройте свой бизнес для миллионов покупателей по всему Казахстану
          </p>
          <button
            class="mt-8 px-8 py-4 bg-white text-blue-600 rounded-xl font-medium hover:bg-blue-50 transition-colors"
          >
            Стать продавцом
          </button>
        </div>
      </div>
    </div>

    <div class="max-w-7xl mx-auto px-4 py-12">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8 mb-16">
        <div
          v-for="benefit in benefits"
          :key="benefit.id"
          class="bg-white p-8 rounded-2xl hover:shadow-xl transition-all duration-300"
        >
          <div class="w-16 h-16 rounded-xl bg-blue-100 flex items-center justify-center mb-6">
            <component :is="benefit.icon" class="w-8 h-8 text-blue-600" />
          </div>
          <h3 class="text-xl font-semibold text-gray-900 mb-4">{{ benefit.title }}</h3>
          <p class="text-gray-600">{{ benefit.description }}</p>
        </div>
      </div>

      <div class="bg-white rounded-2xl p-8 mb-16">
        <h2 class="text-2xl font-bold text-gray-900 mb-8">Как начать продавать</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
          <div v-for="step in steps" :key="step.id" class="text-center">
            <div
              class="w-12 h-12 rounded-full bg-blue-100 flex items-center justify-center mx-auto mb-4"
            >
              <span class="text-xl font-bold text-blue-600">{{ step.id }}</span>
            </div>
            <h3 class="font-medium text-gray-900 mb-2">{{ step.title }}</h3>
            <p class="text-sm text-gray-500">{{ step.description }}</p>
          </div>
        </div>
      </div>

      <div class="bg-blue-50 rounded-2xl p-8">
        <div class="flex flex-col md:flex-row items-center gap-8">
          <div class="flex-1">
            <h2 class="text-2xl font-bold text-gray-900 mb-4">Остались вопросы?</h2>
            <p class="text-gray-600 mb-6">
              Наши менеджеры готовы ответить на все ваши вопросы о продажах на MarketPlace
            </p>
            <button
              class="px-6 py-3 bg-blue-600 text-white rounded-xl hover:bg-blue-700 transition-colors"
            >
              Получить консультацию
            </button>
          </div>
          <!-- <img src="/images/sell/support.jpg" alt="Поддержка" class="w-full md:w-96 rounded-xl" /> -->
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import SettingsIcon from '@/components/icons/SettingsIcon.vue'
import TruckIcon from '@/components/icons/TruckIcon.vue'
import UsersIcon from '@/components/icons/UsersIcon.vue'

const benefits = [
  {
    id: 1,
    title: 'Большая аудитория',
    description: 'Более 1 миллиона активных покупателей ежемесячно',
    icon: UsersIcon,
  },
  {
    id: 2,
    title: 'Простое управление',
    description: 'Удобный личный кабинет и API для интеграции',
    icon: SettingsIcon,
  },

  {
    id: 3,
    title: 'Логистика',
    description: 'Доставка по всему Казахстану через нашу сеть партнеров',
    icon: TruckIcon,
  },
]

const steps = [
  {
    id: 1,
    title: 'Регистрация',
    description: 'Заполните анкету и загрузите документы',
  },
  {
    id: 2,
    title: 'Модерация',
    description: 'Проверка документов занимает 1-2 рабочих дня',
  },
  {
    id: 3,
    title: 'Добавление товаров',
    description: 'Загрузите товары через личный кабинет или API',
  },
  {
    id: 4,
    title: 'Начало продаж',
    description: 'Получайте заказы и зарабатывайте',
  },
]
</script>
