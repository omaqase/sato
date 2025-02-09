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
                  >{{ favorites.length }}
                  {{ getNoun(favorites.length, 'товар', 'товара', 'товаров') }}</span
                >
                <button
                  v-if="favorites.length"
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
      <div
        v-if="favorites.length"
        class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6"
      >
        <div
          v-for="product in favorites"
          :key="product.id"
          class="bg-white rounded-2xl shadow-sm overflow-hidden group flex flex-col"
        >
          <!-- Product Image -->
          <div class="relative aspect-square overflow-hidden">
            <div class="w-full h-full">
              <img
                :src="product.image"
                :alt="product.name"
                class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
              />
              <div class="absolute inset-0 ring-1 ring-inset ring-black/5" />
            </div>

            <!-- Quick Actions -->
            <div class="absolute top-4 right-4 flex flex-col gap-2 z-10">
              <button
                @click="removeFromFavorites(product.id)"
                class="p-2 rounded-xl bg-white/90 backdrop-blur-sm text-gray-700 hover:text-red-600 shadow-sm transition-colors"
              >
                <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none">
                  <path
                    d="M6 18L18 6M6 6l12 12"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                  />
                </svg>
              </button>
            </div>
          </div>

          <div class="p-6 flex flex-col flex-1">
            <!-- Price Block -->
            <div class="mb-4">
              <div class="flex items-baseline gap-2">
                <span class="text-2xl font-bold text-gray-900"
                  >{{ formatPrice(product.price) }} ₽</span
                >
                <span v-if="product.oldPrice" class="text-sm text-gray-400 line-through">
                  {{ formatPrice(product.oldPrice) }} ₽
                </span>
              </div>
              <div v-if="product.discount" class="mt-1">
                <span class="text-sm font-medium text-red-600">Скидка {{ product.discount }}%</span>
              </div>
            </div>

            <!-- Title -->
            <h3
              class="text-lg font-medium text-gray-900 group-hover:text-blue-600 transition-colors line-clamp-2 min-h-[3.5rem]"
            >
              {{ product.name }}
            </h3>

            <!-- Spacer -->
            <div class="flex-1"></div>

            <!-- Rating -->
            <div class="flex items-center gap-2 mb-4">
              <div class="flex text-yellow-400">
                <svg
                  v-for="i in 5"
                  :key="i"
                  class="w-4 h-4"
                  :class="i <= product.rating ? 'fill-current' : 'stroke-current fill-none'"
                >
                  <path
                    d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"
                  />
                </svg>
              </div>
              <span class="text-sm text-gray-500">{{ product.reviewsCount }} отзывов</span>
            </div>

            <!-- Add to Cart -->
            <button
              class="w-full py-3 bg-blue-600 hover:bg-blue-700 text-white font-medium rounded-xl transition-colors flex items-center justify-center gap-2 group"
            >
              <span>В корзину</span>
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
            </button>
          </div>
        </div>
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
import { ref } from 'vue'

interface Product {
  id: number
  name: string
  price: number
  oldPrice?: number
  discount?: number
  rating: number
  reviewsCount: number
  image: string
}

// Моковые данные
const favorites = ref<Product[]>([
  {
    id: 1,
    name: 'iPhone 13 Pro 256GB Графитовый',
    price: 89990,
    oldPrice: 99990,
    discount: 10,
    rating: 5,
    reviewsCount: 128,
    image: 'https://via.placeholder.com/400x500',
  },
  {
    id: 2,
    name: 'MacBook Pro 14" M1 Pro 16/512GB Space Gray',
    price: 179990,
    oldPrice: 189990,
    discount: 5,
    rating: 4,
    reviewsCount: 64,
    image: 'https://via.placeholder.com/400x500',
  },
])

function formatPrice(price: number): string {
  return price.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ' ')
}

function removeFromFavorites(id: number) {
  favorites.value = favorites.value.filter((item) => item.id !== id)
}

function clearAll() {
  favorites.value = []
}

function getNoun(number: number, one: string, two: string, five: string): string {
  const cases = [2, 0, 1, 1, 1, 2]
  return [one, two, five][
    number % 100 > 4 && number % 100 < 20 ? 2 : cases[number % 10 < 5 ? number % 10 : 5]
  ]
}
</script>
