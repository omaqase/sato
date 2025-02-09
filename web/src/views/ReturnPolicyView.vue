<template>
  <div class="max-w-4xl mx-auto px-4 py-12">
    <h1 class="text-3xl font-bold text-gray-900 mb-8">Возврат товара</h1>

    <div class="space-y-8">
      <section
        v-for="section in sections"
        :key="section.id"
        class="bg-white rounded-2xl p-8 shadow-sm"
      >
        <h2 class="text-xl font-semibold text-gray-900 mb-4">{{ section.title }}</h2>
        <div class="prose prose-blue max-w-none">
          <p class="text-gray-600">{{ section.description }}</p>
          <ul class="mt-4 space-y-2">
            <li v-for="item in section.items" :key="item" class="flex items-start gap-3">
              <svg class="w-5 h-5 text-blue-600 mt-1 flex-shrink-0" viewBox="0 0 24 24" fill="none">
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
            <h3 class="text-lg font-medium text-gray-900 mb-2">Остались вопросы?</h3>
            <p class="text-gray-600 mb-4">
              Свяжитесь с нашей службой поддержки, мы поможем оформить возврат
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
const sections = [
  {
    id: 1,
    title: 'Условия возврата',
    description:
      'Вы можете вернуть товар в течение 14 дней с момента получения заказа, если он вас не устроил.',
    items: [
      'Товар не был в использовании',
      'Сохранены товарный вид и потребительские свойства',
      'Сохранена оригинальная упаковка',
      'Есть документ, подтверждающий покупку',
    ],
  },
  {
    id: 2,
    title: 'Как оформить возврат',
    description: 'Процесс возврата товара состоит из нескольких простых шагов:',
    items: [
      'Заполните заявление на возврат в личном кабинете',
      'Упакуйте товар в оригинальную упаковку',
      'Отправьте товар через пункт выдачи или курьерской службой',
      'Получите возврат денежных средств в течение 10 дней',
    ],
  },
  {
    id: 3,
    title: 'Возврат денежных средств',
    description:
      'После получения и проверки возвращаемого товара мы осуществим возврат денежных средств:',
    items: [
      'На банковскую карту: 3-10 рабочих дней',
      'На электронный кошелек: 1-3 рабочих дня',
      'Наличными в пункте выдачи: моментально',
    ],
  },
]
</script>
