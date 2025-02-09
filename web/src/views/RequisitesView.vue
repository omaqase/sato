<template>
  <div class="max-w-4xl mx-auto px-4 py-12">
    <h1 class="text-3xl font-bold text-gray-900 mb-8">Реквизиты компании</h1>

    <div class="bg-white rounded-2xl shadow-sm p-8">
      <div class="space-y-8">
        <section>
          <h2 class="text-xl font-semibold text-gray-900 mb-4">Основная информация</h2>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <p class="text-sm text-gray-500 mb-1">Полное наименование</p>
              <p class="text-gray-900">Общество с ограниченной ответственностью "МаркетПлейс"</p>
            </div>
            <div>
              <p class="text-sm text-gray-500 mb-1">Сокращенное наименование</p>
              <p class="text-gray-900">ООО "МаркетПлейс"</p>
            </div>
            <div>
              <p class="text-sm text-gray-500 mb-1">ИНН</p>
              <p class="text-gray-900">7731347089</p>
            </div>
            <div>
              <p class="text-sm text-gray-500 mb-1">КПП</p>
              <p class="text-gray-900">773101001</p>
            </div>
            <div>
              <p class="text-sm text-gray-500 mb-1">ОГРН</p>
              <p class="text-gray-900">1157746644270</p>
            </div>
            <div>
              <p class="text-sm text-gray-500 mb-1">Дата регистрации</p>
              <p class="text-gray-900">17.07.2015</p>
            </div>
          </div>
        </section>

        <section>
          <h2 class="text-xl font-semibold text-gray-900 mb-4">Банковские реквизиты</h2>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <p class="text-sm text-gray-500 mb-1">Расчетный счет</p>
              <p class="text-gray-900">40702810638000123456</p>
            </div>
            <div>
              <p class="text-sm text-gray-500 mb-1">Банк</p>
              <p class="text-gray-900">ПАО СБЕРБАНК</p>
            </div>
            <div>
              <p class="text-sm text-gray-500 mb-1">БИК</p>
              <p class="text-gray-900">044525225</p>
            </div>
            <div>
              <p class="text-sm text-gray-500 mb-1">Корр. счет</p>
              <p class="text-gray-900">30101810400000000225</p>
            </div>
          </div>
        </section>

        <section>
          <h2 class="text-xl font-semibold text-gray-900 mb-4">Юридический адрес</h2>
          <p class="text-gray-900">123456, г. Москва, ул. Примерная, д. 1, офис 123</p>
        </section>

        <div class="flex items-center gap-2 text-sm text-gray-500">
          <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none">
            <path
              d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
            />
          </svg>
          <span>Информация обновлена 01.03.2024</span>
        </div>
      </div>
    </div>
  </div>
</template>
