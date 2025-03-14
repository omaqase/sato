<template>
  <div class="min-h-screen bg-gradient-to-b from-gray-50 to-white pb-12">
    <!-- Hero Section -->
    <div class="bg-gradient-to-r from-red-500 to-red-600 text-white">
      <div class="max-w-7xl mx-auto px-4">
        <div class="py-12">
          <div class="flex items-center gap-6">
            <div
              class="w-24 h-24 rounded-2xl bg-white/10 backdrop-blur-sm flex items-center justify-center"
            >
              <svg class="w-12 h-12 text-white/90" viewBox="0 0 24 24" fill="none">
                <path
                  d="M20.84 4.61a5.5 5.5 0 00-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 00-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 000-7.78z"
                  stroke="currentColor"
                  stroke-width="2"
                />
              </svg>
            </div>
            <div>
              <h1 class="text-3xl font-bold mb-1">Избранное</h1>
              <div class="flex items-center gap-4 text-white/80">
                <span
                  >{{ favoritesStore.favorites.length }}
                  {{ getNoun(favoritesStore.favorites.length, 'товар', 'товара', 'товаров') }}</span
                >
                <button
                  v-if="favoritesStore.favorites.length"
                  @click="clearAll"
                  class="text-white/60 hover:text-white transition-colors"
                >
                  Очистить список
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="max-w-7xl mx-auto px-4 -mt-8">
      <!-- Loading State -->
      <div v-if="favoritesStore.isLoading" class="flex justify-center py-12">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
      </div>

      <!-- Error State -->
      <div v-else-if="favoritesStore.error" class="max-w-2xl mx-auto bg-red-50 rounded-2xl p-8 text-center">
        <div class="w-16 h-16 mx-auto mb-4 text-red-500">
          <svg viewBox="0 0 24 24" fill="none" class="w-full h-full">
            <path d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" 
                  stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
          </svg>
        </div>
        <h3 class="text-lg font-medium text-red-800 mb-2">{{ favoritesStore.error }}</h3>
        <button @click="favoritesStore.fetchFavorites" 
                class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500">
          Try Again
        </button>
      </div>

      <!-- Products Grid -->
      <div v-else-if="favoritesStore.favorites.length" 
           class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
        <ProductCard v-for="product in favoritesStore.favorites" 
                    :key="product.productId" 
                    :product="product"
                    class="h-full">
          <template #actions>
            <button @click="favoritesStore.removeFromFavorites(product.productId)"
                    class="w-full flex items-center justify-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500">
              Remove from Favorites
            </button>
          </template>
        </ProductCard>
      </div>

      <!-- Empty State -->
      <div v-else class="max-w-2xl mx-auto bg-white rounded-2xl shadow-sm p-12 text-center">
        <div
          class="w-24 h-24 mx-auto mb-6 rounded-full bg-gray-50 flex items-center justify-center"
        >
          <svg class="w-12 h-12 text-gray-400" viewBox="0 0 24 24" fill="none">
            <path
              d="M20.84 4.61a5.5 5.5 0 00-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 00-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 000-7.78z"
              stroke="currentColor"
              stroke-width="2"
            />
          </svg>
        </div>
        <h3 class="text-2xl font-bold text-gray-900 mb-2">Нет избранных товаров</h3>
        <p class="text-gray-500 mb-8">
          Добавляйте товары в избранное, чтобы следить за изменениями цен и быстро находить нужные
          товары
        </p>
        <router-link
          to="/"
          class="inline-flex items-center gap-2 px-8 py-3 bg-blue-600 text-white font-medium rounded-xl hover:bg-blue-700 transition-colors group"
        >
          <span>Перейти к покупкам</span>
          <svg
            class="w-5 h-5 transition-transform group-hover:translate-x-1"
            viewBox="0 0 24 24"
            fill="none"
          >
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
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useFavoritesStore } from '@/stores/useFavoritesStore'
import ProductCard from '@/components/ProductCard.vue'

const favoritesStore = useFavoritesStore()

function getNoun(number: number, one: string, two: string, five: string): string {
  const cases = [2, 0, 1, 1, 1, 2]
  return [one, two, five][
    number % 100 > 4 && number % 100 < 20 ? 2 : cases[number % 10 < 5 ? number % 10 : 5]
  ]
}

function clearAll() {
  favoritesStore.favorites = []
}

onMounted(() => {
  favoritesStore.fetchFavorites()
})
</script>
