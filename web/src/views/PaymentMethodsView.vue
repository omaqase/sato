<template>
  <div class="max-w-4xl mx-auto px-4 py-12">
    <h1 class="text-3xl font-bold text-gray-900 mb-8">Способы оплаты</h1>

    <div class="space-y-8">
      <section
        v-for="method in paymentMethods"
        :key="method.id"
        class="bg-white rounded-2xl p-8 shadow-sm"
      >
        <div class="flex items-start gap-6">
          <div class="p-4 bg-gray-50 rounded-xl">
            <component :is="method.icon" class="w-8 h-8 text-gray-900" />
          </div>
          <div class="flex-1">
            <h2 class="text-xl font-semibold text-gray-900 mb-2">{{ method.title }}</h2>
            <p class="text-gray-600 mb-4">{{ method.description }}</p>

            <div v-if="method.options" class="grid grid-cols-2 md:grid-cols-4 gap-4">
              <div
                v-for="option in method.options"
                :key="option.name"
                class="p-4 rounded-xl bg-gray-50 flex items-center justify-center"
              >
                <img :src="option.logo" :alt="option.name" class="h-6" />
              </div>
            </div>

            <ul v-if="method.features" class="mt-4 space-y-2">
              <li
                v-for="feature in method.features"
                :key="feature"
                class="flex items-center gap-2 text-sm text-gray-600"
              >
                <svg class="w-5 h-5 text-green-500" viewBox="0 0 24 24" fill="none">
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
      </section>

      <div class="bg-blue-50 rounded-2xl p-6">
        <div class="flex items-start gap-4">
          <div class="p-3 rounded-xl bg-blue-100">
            <svg class="w-6 h-6 text-blue-600" viewBox="0 0 24 24" fill="none">
              <path
                d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"
                stroke="currentColor"
                stroke-width="2"
              />
            </svg>
          </div>
          <div>
            <h3 class="text-lg font-medium text-gray-900 mb-2">Безопасная оплата</h3>
            <p class="text-gray-600">Все платежи защищены современными технологиями шифрования</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const paymentMethods = [
  {
    id: 1,
    title: 'Банковские карты',
    description: 'Оплачивайте заказы любой банковской картой через защищенное соединение',
    options: [
      { name: 'Visa', logo: '/images/payment/visa.svg' },
      { name: 'Mastercard', logo: '/images/payment/mastercard.svg' },
      { name: 'Мир', logo: '/images/payment/mir.svg' },
    ],
    features: ['Мгновенное зачисление', 'Без комиссии', 'Безопасная оплата'],
  },
  {
    id: 2,
    title: 'Электронные кошельки',
    description: 'Используйте популярные электронные кошельки для оплаты',
    options: [
      { name: 'ЮMoney', logo: '/images/payment/ymoney.svg' },
      { name: 'QIWI', logo: '/images/payment/qiwi.svg' },
    ],
  },
  {
    id: 3,
    title: 'Оплата при получении',
    description: 'Оплатите заказ при получении наличными или картой',
    features: ['Оплата после проверки товара', 'Наличные или карта', 'Кассовый чек'],
  },
]
</script>
