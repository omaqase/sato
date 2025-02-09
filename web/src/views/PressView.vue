<template>
  <div class="min-h-screen bg-gradient-to-b from-gray-50 to-white pb-12">
    <div class="max-w-7xl mx-auto px-4 py-12">
      <h1 class="text-3xl font-bold text-gray-900 mb-8">Пресс-центр</h1>

      <!-- Фильтры -->
      <div class="flex items-center gap-4 mb-8">
        <button
          v-for="category in categories"
          :key="category.id"
          :class="[
            'px-4 py-2 rounded-xl text-sm font-medium transition-all duration-200',
            activeCategory === category.id
              ? 'bg-blue-600 text-white'
              : 'bg-white text-gray-600 hover:bg-gray-50',
          ]"
          @click="activeCategory = category.id"
        >
          {{ category.name }}
        </button>
      </div>

      <!-- Сетка новостей -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div
          v-for="article in filteredArticles"
          :key="article.id"
          class="bg-white rounded-2xl overflow-hidden hover:shadow-xl transition-all duration-300"
        >
          <div class="aspect-video relative overflow-hidden">
            <img
              :src="article.image"
              :alt="article.title"
              class="w-full h-full object-cover hover:scale-110 transition-transform duration-500"
            />
            <div class="absolute bottom-4 left-4">
              <span class="px-3 py-1 bg-blue-600 text-white text-xs font-medium rounded-full">
                {{ article.category }}
              </span>
            </div>
          </div>
          <div class="p-6">
            <div class="text-sm text-gray-500 mb-2">{{ formatDate(article.date) }}</div>
            <h3 class="text-xl font-semibold text-gray-900 mb-3">{{ article.title }}</h3>
            <p class="text-gray-600 mb-4 line-clamp-2">{{ article.description }}</p>
            <router-link
              :to="{ name: 'press-article', params: { id: article.id } }"
              class="inline-flex items-center gap-2 text-blue-600 font-medium hover:text-blue-700 transition-colors"
            >
              Читать далее
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
import { ref, computed } from 'vue'

const categories = [
  { id: 'all', name: 'Все новости' },
  { id: 'company', name: 'О компании' },
  { id: 'products', name: 'Новые товары' },
  { id: 'events', name: 'Мероприятия' },
]

const activeCategory = ref('all')

const articles = [
  {
    id: 1,
    title: 'MarketPlace запускает новую программу лояльности',
    description:
      'Теперь покупатели могут получать больше бонусов за покупки и специальные предложения.',
    category: 'company',
    date: '2024-03-15',
    image: '/press_center_loyal_program.jpg',
  },
  // Добавьте больше статей
]

const filteredArticles = computed(() => {
  if (activeCategory.value === 'all') return articles
  return articles.filter((article) => article.category === activeCategory.value)
})

const formatDate = (date: string) => {
  return new Date(date).toLocaleDateString('ru-RU', {
    day: 'numeric',
    month: 'long',
    year: 'numeric',
  })
}
</script>
