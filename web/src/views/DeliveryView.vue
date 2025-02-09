<template>
  <div class="max-w-4xl mx-auto px-4 py-12">
    <h1 class="text-3xl font-bold text-gray-900 mb-8">Доставка</h1>

    <div class="space-y-8">
      <section
        v-for="method in deliveryMethods"
        :key="method.id"
        class="bg-white rounded-2xl p-8 shadow-sm"
      >
        <div class="flex items-start gap-6">
          <div class="p-4 bg-gray-50 rounded-xl">
            <component :is="method.icon" class="w-8 h-8 text-gray-900" />
          </div>
          <div class="flex-1">
            <div class="flex items-center justify-between mb-4">
              <div>
                <h2 class="text-xl font-semibold text-gray-900">{{ method.title }}</h2>
                <p class="text-gray-600">{{ method.description }}</p>
              </div>
              <div class="text-right">
                <div class="text-2xl font-bold text-gray-900">{{ method.price }}</div>
                <div class="text-sm text-gray-500">{{ method.time }}</div>
              </div>
            </div>

            <div class="bg-gray-50 rounded-xl p-4">
              <ul class="space-y-3">
                <li
                  v-for="feature in method.features"
                  :key="feature"
                  class="flex items-center gap-3 text-sm text-gray-600"
                >
                  <svg class="w-5 h-5 text-green-500 flex-shrink-0" viewBox="0 0 24 24" fill="none">
                    <path
                      d="M20 6L9 17l-5-5"
                      stroke="currentColor"
                      stroke-width="2"
                      stroke-linecap="round"
                    />
                  </svg>
                  {{ feature }}
                </li>
              </ul>
            </div>
          </div>
        </div>
      </section>

      <div class="bg-blue-50 rounded-2xl p-6">
        <div class="flex items-start gap-4">
          <div class="p-3 rounded-xl bg-blue-100">
            <svg class="w-6 h-6 text-blue-600" viewBox="0 0 24 24" fill="none">
              <path
                d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
              />
            </svg>
          </div>
          <div>
            <h3 class="text-lg font-medium text-gray-900 mb-2">Отслеживание заказа</h3>
            <p class="text-gray-600 mb-4">
              После оформления заказа вы получите трек-номер для отслеживания
            </p>
            <router-link
              to="/profile/orders"
              class="inline-flex items-center gap-2 text-blue-600 font-medium hover:text-blue-700"
            >
              Мои заказы
              <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none">
                <path
                  d="M5 12h14m-6-6l6 6-6 6"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linecap="round"
                />
              </svg>
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const deliveryMethods = [
  {
    id: 1,
    title: 'Курьерская доставка',
    description: 'Доставка до двери в удобное время',
    price: 'от 299 ₽',
    time: '1-2 дня',
    features: ['Доставка до двери', 'Выбор удобного времени', 'SMS-уведомления о статусе'],
  },
  {
    id: 2,
    title: 'Пункты выдачи',
    description: 'Получение заказа в ближайшем пункте выдачи',
    price: 'от 199 ₽',
    time: '1-3 дня',
    features: ['Более 1000 пунктов выдачи', 'Примерка перед покупкой', 'Хранение заказа 7 дней'],
  },
  {
    id: 3,
    title: 'Почта России',
    description: 'Доставка в любую точку России',
    price: 'от 299 ₽',
    time: '3-14 дней',
    features: ['Доставка по всей России', 'Наложенный платеж', 'Трекинг отправления'],
  },
]
</script>
