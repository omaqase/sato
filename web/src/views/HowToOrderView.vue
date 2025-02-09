<template>
  <div class="max-w-4xl mx-auto px-4 py-12">
    <h1 class="text-3xl font-bold text-gray-900 mb-8">Как сделать заказ</h1>

    <div class="space-y-12">
      <div v-for="(step, index) in steps" :key="step.id" class="flex gap-8 items-start">
        <div
          class="flex-shrink-0 w-12 h-12 rounded-full bg-blue-100 flex items-center justify-center"
        >
          <span class="text-xl font-bold text-blue-600">{{ index + 1 }}</span>
        </div>
        <div class="flex-1">
          <h3 class="text-xl font-semibold text-gray-900 mb-4">{{ step.title }}</h3>
          <p class="text-gray-600 mb-6">{{ step.description }}</p>
          <div class=
          "bg-white rounded-2xl overflow-hidden shadow-sm">
            <img :src="step.image" :alt="step.title" class="w-full h-auto" />
          </div>
        </div>
      </div>

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
            <h3 class="text-lg font-medium text-gray-900 mb-2">Остались вопросы?</h3>
            <p class="text-gray-600 mb-4">
              Наша служба поддержки готова помочь вам с оформлением заказа
            </p>
            <button
              class="inline-flex items-center gap-2 text-blue-600 font-medium hover:text-blue-700"
            >
              Написать в поддержку
              <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none">
                <path
                  d="M5 12h14m-6-6l6 6-6 6"
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
const steps = [
  {
    id: 1,
    title: 'Выберите товары',
    description:
      'Добавьте нужные товары в корзину. Вы можете сравнивать цены, читать отзывы и смотреть рейтинги продавцов.',
    image: '/images/how-to-order/step1.jpg',
  },
  {
    id: 2,
    title: 'Перейдите в корзину',
    description: 'Проверьте выбранные товары, их количество и общую стоимость заказа.',
    image: '/images/how-to-order/step2.jpg',
  },
  {
    id: 3,
    title: 'Укажите адрес доставки',
    description: 'Выберите удобный способ доставки и введите адрес.',
    image: '/images/how-to-order/step3.jpg',
  },
  {
    id: 4,
    title: 'Выберите способ оплаты',
    description:
      'Оплатите заказ удобным способом: картой, электронным кошельком или при получении.',
    image: '/images/how-to-order/step4.jpg',
  },
]
</script>
