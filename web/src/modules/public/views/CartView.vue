<template>
  <div class="min-h-screen bg-gradient-to-b from-gray-50 to-white pb-12">
    <!-- Hero Section -->
    <div class="bg-gradient-to-r from-blue-600 to-blue-700 text-white">
      <div class="max-w-7xl mx-auto px-4">
        <div class="py-12">
          <div class="flex items-center gap-6">
            <div
              class="w-24 h-24 rounded-2xl bg-white/10 backdrop-blur-sm flex items-center justify-center"
            >
              <svg class="w-12 h-12 text-white/90" viewBox="0 0 24 24" fill="none">
                <path
                  d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.8 19h12.4L17 13"
                  stroke="currentColor"
                  stroke-width="2"
                />
              </svg>
            </div>
            <div>
              <h1 class="text-3xl font-bold mb-1">Корзина</h1>
              <div class="flex items-center gap-4 text-white/80">
                <span
                  >{{ totalItems }} {{ getNoun(totalItems, 'товар', 'товара', 'товаров') }}</span
                >
                <span class="text-white/40">•</span>
                <span>На сумму {{ formatPrice(total) }} ₽</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="max-w-7xl mx-auto px-4 -mt-8">
      <div v-if="cartItems.length" class="flex flex-col lg:flex-row gap-8">
        <!-- Cart Items -->
        <div class="flex-1">
          <div class="bg-white rounded-2xl shadow-sm overflow-hidden">
            <div class="divide-y divide-gray-100">
              <div
                v-for="item in cartItems"
                :key="item.id"
                class="p-6 hover:bg-gray-50 transition-colors"
              >
                <div class="flex gap-6">
                  <!-- Product Image -->
                  <div class="relative group">
                    <img
                      :src="item.image"
                      :alt="item.name"
                      class="w-32 h-32 rounded-xl object-cover group-hover:scale-105 transition-transform duration-300"
                    />
                    <div class="absolute inset-0 rounded-xl ring-1 ring-inset ring-black/5" />
                  </div>

                  <!-- Product Info -->
                  <div class="flex-1 min-w-0">
                    <div class="flex justify-between">
                      <div>
                        <h3 class="font-medium text-gray-900">{{ item.name }}</h3>
                        <div class="mt-1 text-sm text-gray-500 space-y-1">
                          <p>Артикул: {{ item.sku }}</p>
                          <p>Продавец: {{ item.seller }}</p>
                        </div>
                      </div>
                      <button
                        @click="removeFromCart(item.id)"
                        class="p-2 text-gray-400 hover:text-red-500 transition-colors"
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

                    <div class="mt-4 flex items-center justify-between">
                      <div class="flex items-center gap-6">
                        <!-- Quantity -->
                        <div class="flex items-center gap-3">
                          <button
                            @click="decreaseQuantity(item.id)"
                            class="w-8 h-8 rounded-lg border border-gray-200 flex items-center justify-center text-gray-500 hover:bg-gray-50 transition-colors"
                          >
                            <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none">
                              <path
                                d="M5 12h14"
                                stroke="currentColor"
                                stroke-width="2"
                                stroke-linecap="round"
                              />
                            </svg>
                          </button>
                          <span class="w-8 text-center font-medium">{{ item.quantity }}</span>
                          <button
                            @click="increaseQuantity(item.id)"
                            class="w-8 h-8 rounded-lg border border-gray-200 flex items-center justify-center text-gray-500 hover:bg-gray-50 transition-colors"
                          >
                            <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none">
                              <path
                                d="M12 5v14m-7-7h14"
                                stroke="currentColor"
                                stroke-width="2"
                                stroke-linecap="round"
                              />
                            </svg>
                          </button>
                        </div>

                        <!-- Stock Status -->
                        <div class="text-sm text-green-600 flex items-center gap-1">
                          <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none">
                            <path
                              d="M20 6L9 17l-5-5"
                              stroke="currentColor"
                              stroke-width="2"
                              stroke-linecap="round"
                            />
                          </svg>
                          В наличии
                        </div>
                      </div>

                      <!-- Price -->
                      <div class="text-right">
                        <div class="text-sm text-gray-500 line-through">
                          {{ formatPrice(item.oldPrice) }} ₽
                        </div>
                        <div class="text-xl font-bold text-gray-900">
                          {{ formatPrice(item.price * item.quantity) }} ₽
                        </div>
                        <div class="text-sm text-red-600">Скидка {{ item.discount }}%</div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Order Summary -->
        <div class="lg:w-96">
          <div class="bg-white rounded-2xl shadow-sm p-6 sticky top-24">
            <h2 class="text-lg font-semibold mb-6">Итого</h2>

            <div class="space-y-4 mb-6">
              <div class="flex justify-between text-sm">
                <span class="text-gray-600">Товары ({{ totalItems }})</span>
                <span class="font-medium">{{ formatPrice(subtotal) }} ₽</span>
              </div>
              <div class="flex justify-between text-sm">
                <span class="text-gray-600">Скидка</span>
                <span class="text-red-600 font-medium">-{{ formatPrice(discount) }} ₽</span>
              </div>
              <div class="flex justify-between text-sm">
                <span class="text-gray-600">Доставка</span>
                <span class="font-medium">Бесплатно</span>
              </div>
            </div>

            <div class="pt-4 border-t">
              <div class="flex justify-between items-center mb-1">
                <span class="text-lg font-bold">К оплате</span>
                <span class="text-2xl font-bold text-blue-600">{{ formatPrice(total) }} ₽</span>
              </div>
              <div class="text-xs text-gray-500">Включая НДС</div>
            </div>

            <button
              class="w-full mt-6 py-4 bg-blue-600 hover:bg-blue-700 text-white font-medium rounded-xl transition-colors flex items-center justify-center gap-2 group"
            >
              <span>Оформить заказ</span>
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

            <div class="mt-6 flex items-center gap-2 text-sm text-gray-500">
              <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none">
                <path
                  d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"
                  stroke="currentColor"
                  stroke-width="2"
                />
              </svg>
              Безопасная оплата и гарантия возврата
            </div>
          </div>
        </div>
      </div>

      <!-- Empty Cart -->
      <div v-else class="max-w-2xl mx-auto bg-white rounded-2xl shadow-sm p-12 text-center">
        <div
          class="w-24 h-24 mx-auto mb-6 rounded-full bg-gray-50 flex items-center justify-center"
        >
          <svg class="w-12 h-12 text-gray-400" viewBox="0 0 24 24" fill="none">
            <path
              d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.8 19h12.4L17 13"
              stroke="currentColor"
              stroke-width="2"
            />
          </svg>
        </div>
        <h3 class="text-2xl font-bold text-gray-900 mb-2">Корзина пуста</h3>
        <p class="text-gray-500 mb-8">
          Воспользуйтесь поиском или каталогом, чтобы найти нужные товары
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
import { ref, computed } from 'vue'

interface CartItem {
  id: number
  name: string
  sku: string
  price: number
  oldPrice: number
  discount: number
  quantity: number
  image: string
  seller: string
}

// Моковые данные
const cartItems = ref<CartItem[]>([
  {
    id: 1,
    name: 'iPhone 13 Pro 256GB Графитовый',
    sku: 'IP13P256G',
    price: 89990,
    oldPrice: 99990,
    discount: 10,
    quantity: 1,
    image: 'https://via.placeholder.com/400',
    seller: 'Apple',
  },
  {
    id: 2,
    name: 'AirPods Pro 2nd generation',
    sku: 'APP2G',
    price: 22990,
    oldPrice: 24990,
    discount: 8,
    quantity: 2,
    image: 'https://via.placeholder.com/400',
    seller: 'Apple',
  },
])

// Вычисляемые свойства для итогов
const subtotal = computed(() =>
  cartItems.value.reduce((sum, item) => sum + item.oldPrice * item.quantity, 0),
)

const total = computed(() =>
  cartItems.value.reduce((sum, item) => sum + item.price * item.quantity, 0),
)

const discount = computed(() => subtotal.value - total.value)

const totalItems = computed(() => cartItems.value.reduce((sum, item) => sum + item.quantity, 0))

// Методы управления корзиной
function removeFromCart(id: number) {
  cartItems.value = cartItems.value.filter((item) => item.id !== id)
}

function increaseQuantity(id: number) {
  const item = cartItems.value.find((item) => item.id === id)
  if (item) item.quantity++
}

function decreaseQuantity(id: number) {
  const item = cartItems.value.find((item) => item.id === id)
  if (item && item.quantity > 1) item.quantity--
}

function formatPrice(price: number): string {
  return price.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ' ')
}

function getNoun(number: number, one: string, two: string, five: string): string {
  const n = Math.abs(number)
  if (n % 10 === 1 && n % 100 !== 11) {
    return one
  } else if (n % 10 >= 2 && n % 10 <= 4 && (n % 100 < 10 || n % 100 >= 20)) {
    return two
  } else {
    return five
  }
}
</script>
