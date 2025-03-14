<template>
  <div class="min-h-screen bg-gradient-to-b from-gray-50 to-white">
    <!-- Hero Banner -->
    <div class="bg-gradient-to-r from-blue-600 to-blue-700">
      <div class="max-w-7xl mx-auto px-4">
        <div class="relative py-16">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center">
            <div class="text-white space-y-8">
              <div>
                <span
                  class="inline-flex items-center gap-2 px-4 py-2 rounded-full bg-white/10 backdrop-blur-sm text-sm font-medium mb-4"
                >
                  <span class="w-2 h-2 rounded-full bg-yellow-400 animate-pulse"></span>
                  Специальное предложение
                </span>
                <h1 class="text-5xl font-bold leading-tight">
                  Весенняя распродажа
                  <br />
                  <span class="text-yellow-400">Скидки до 70%</span>
                </h1>
              </div>
              <p class="text-white/80 text-xl">
                Успейте купить товары известных брендов
                <br />
                по самым выгодным ценам
              </p>
              <div class="flex items-center gap-4">
                <button
                  class="px-8 py-4 bg-white text-blue-600 rounded-xl font-medium hover:bg-white/90 transition-all duration-300 hover:shadow-lg hover:scale-105"
                >
                  Перейти к товарам
                </button>
                <button
                  class="px-8 py-4 bg-white/10 text-white rounded-xl font-medium hover:bg-white/20 transition-all duration-300"
                >
                  Подробнее
                </button>
              </div>
            </div>
            <div class="relative">
              <div
                class="absolute inset-0 bg-gradient-to-r from-blue-600/90 to-blue-700/90 rounded-2xl"
              />
              <div class="w-full h-[500px] rounded-2xl bg-gray-200 animate-pulse" />
              <!-- Декоративные элементы -->
              <div
                class="absolute -top-8 -right-8 w-32 h-32 bg-yellow-400/20 rounded-full blur-3xl"
              ></div>
              <div
                class="absolute -bottom-8 -left-8 w-40 h-40 bg-blue-400/20 rounded-full blur-3xl"
              ></div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Categories Grid -->
    <div class="max-w-7xl mx-auto px-4 py-16">
      <div class="flex items-center justify-between mb-8">
        <h2 class="text-2xl font-bold text-gray-900">Популярные категории</h2>
        <button class="text-blue-600 hover:text-blue-700 font-medium transition-colors">
          Все категории
        </button>
      </div>
      <div class="grid grid-cols-2 md:grid-cols-4 gap-6">
        <div
          v-for="category in categories"
          :key="category.id"
          class="group relative h-48 bg-gradient-to-br rounded-2xl overflow-hidden cursor-pointer transform hover:scale-105 transition-all duration-300 hover:shadow-xl"
          :class="category.gradient"
        >
          <img
            :src="category.image"
            :alt="category.name"
            class="absolute inset-0 w-full h-full object-cover mix-blend-overlay opacity-20"
          />
          <div
            class="absolute inset-0 p-6 flex flex-col justify-between bg-gradient-to-t from-black/50 via-black/0 to-transparent"
          >
            <h3 class="text-xl font-medium text-white">{{ category.name }}</h3>
            <p class="text-sm text-white/80">от {{ category.price }} ₽</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Products Grid -->
    <div class="max-w-7xl mx-auto px-4 py-16">
      <div class="flex items-center justify-between mb-8">
        <div>
          <h2 class="text-2xl font-bold text-gray-900 mb-2">Популярные товары</h2>
          <p class="text-gray-500">Самые популярные товары на основе продаж</p>
        </div>
        <div class="flex items-center gap-4">
          <button
            class="p-3 rounded-xl border-2 border-gray-200 hover:border-blue-600 hover:bg-blue-50 transition-all duration-300"
          >
            <svg class="w-5 h-5 text-gray-400 hover:text-blue-600" viewBox="0 0 24 24" fill="none">
              <path
                d="M15 6l-6 6 6 6"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
              />
            </svg>
          </button>
          <button
            class="p-3 rounded-xl border-2 border-gray-200 hover:border-blue-600 hover:bg-blue-50 transition-all duration-300"
          >
            <svg class="w-5 h-5 text-gray-400 hover:text-blue-600" viewBox="0 0 24 24" fill="none">
              <path
                d="M9 6l6 6-6 6"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
              />
            </svg>
          </button>
        </div>
      </div>

      <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-5 gap-6">
        <div v-for="product in products" :key="product.productId">
          <ProductCard :product="product">
            <div class="relative aspect-square mb-4 overflow-hidden rounded-xl">
              <img
                :src="product.image || '/products/iphone.jpg'"
                :alt="product.title"
                class="w-full h-full object-cover group-hover:scale-110 transition-transform duration-500"
              />
              <div class="absolute top-2 right-2 flex flex-col gap-2">
                <button
                  class="p-2 rounded-xl bg-white/80 backdrop-blur-sm hover:bg-white transition-colors shadow-lg"
                >
                  <svg
                    class="w-5 h-5 text-gray-400 hover:text-red-500 transition-colors"
                    viewBox="0 0 24 24"
                    fill="none"
                  >
                    <path
                      d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"
                      stroke="currentColor"
                      stroke-width="2"
                    />
                  </svg>
                </button>
                <button
                  class="p-2 rounded-xl bg-white/80 backdrop-blur-sm hover:bg-white transition-colors shadow-lg"
                >
                  <svg
                    class="w-5 h-5 text-gray-400 hover:text-blue-500 transition-colors"
                    viewBox="0 0 24 24"
                    fill="none"
                  >
                    <path
                      d="M21 11.5a8.38 8.38 0 01-.9 3.8 8.5 8.5 0 01-7.6 4.7 8.38 8.38 0 01-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 01-.9-3.8 8.5 8.5 0 014.7-7.6 8.38 8.38 0 013.8-.9h.5a8.48 8.48 0 018 8v.5z"
                      stroke="currentColor"
                      stroke-width="2"
                    />
                  </svg>
                </button>
              </div>
            </div>
            <div>
              <div class="flex items-center gap-1 mb-2">
                <div class="flex">
                  <svg
                    v-for="i in 5"
                    :key="i"
                    class="w-4 h-4"
                    :class="i <= Math.floor(product.rating || 0) ? 'text-yellow-400' : 'text-gray-200'"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                  >
                    <path
                      d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"
                  />
                  </svg>
                </div>
              </div>
              <h3
                class="font-medium text-gray-900 mb-2 line-clamp-2 group-hover:text-blue-600 transition-colors"
              >
                {{ product.title }}
              </h3>
              <div class="flex items-center justify-between">
                <div>
                  <div class="text-lg font-bold text-gray-900">
                    {{ formatPrice(product.price || 0) }} ₽
                  </div>
                </div>
                <button
                  class="p-3 rounded-xl bg-blue-600 text-white hover:bg-blue-700 transition-colors shadow-lg hover:shadow-blue-100"
                >
                  <svg class="w-6 h-6" viewBox="0 0 24 24" fill="none">
                    <path
                      d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.8 19h12.4L17 13"
                      stroke="currentColor"
                      stroke-width="2"
                    />
                  </svg>
                </button>
              </div>
            </div>
          </ProductCard>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { catalogueApi } from '@/api/catalogue'
import ProductCard from '@/components/ProductCard.vue'

const categories = ref([
  {
    id: 1,
    name: 'Электроника',
    price: '999',
    image: '/categories/electronics.jpg',
    gradient: 'from-blue-600 to-blue-800',
  },
  {
    id: 2,
    name: 'Одежда',
    price: '499',
    image: '/categories/clothing.jpg',
    gradient: 'from-purple-600 to-purple-800',
  },
  {
    id: 3,
    name: 'Обувь',
    price: '1 499',
    image: '/categories/shoes.jpg',
    gradient: 'from-red-600 to-red-800',
  },
  {
    id: 4,
    name: 'Аксессуары',
    price: '299',
    image: '/categories/accessories.jpg',
    gradient: 'from-green-600 to-green-800',
  },
])

const products = ref([])

const formatPrice = (price: number) => {
  return new Intl.NumberFormat('ru-RU').format(price)
}

onMounted(async () => {
  try {
    const response = await catalogueApi.listProducts({ limit: 10, offset: 0 })
    products.value = response.products
  } catch (error) {
    console.error('Failed to fetch products:', error)
  }
})
</script>