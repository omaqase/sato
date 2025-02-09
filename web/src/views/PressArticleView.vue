<template>
  <div class="max-w-4xl mx-auto px-4 py-12">
    <div class="mb-8">
      <router-link
        to="/press"
        class="inline-flex items-center gap-2 text-gray-500 hover:text-blue-600 transition-colors"
      >
        <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none">
          <path
            d="M19 12H5m7 7l-7-7 7-7"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
          />
        </svg>
        Назад к новостям
      </router-link>
    </div>

    <article class="prose prose-blue max-w-none">
      <span class="text-sm text-gray-500">15 марта 2024</span>
      <h1>MarketPlace запускает новую программу лояльности</h1>
      <!-- <img src="/images/press/loyalty.jpg" alt="Программа лояльности" class="rounded-2xl" /> -->
      <p>Полный текст статьи...</p>
    </article>
  </div>
</template>
