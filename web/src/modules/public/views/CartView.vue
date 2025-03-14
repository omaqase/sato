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
                <span>{{ cartStore.itemCount }} {{ getNoun(cartStore.itemCount, 'товар', 'товара', 'товаров') }}</span>
                <button
                  v-if="cartStore.itemCount"
                  @click="clearCart"
                  class="text-white/60 hover:text-white transition-colors"
                >
                  Очистить корзину
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="max-w-7xl mx-auto px-4 -mt-8">
      <!-- Loading State -->
      <div v-if="cartStore.isLoading" class="flex justify-center py-12">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
      </div>

      <!-- Error State -->
      <div v-else-if="cartStore.error" class="max-w-2xl mx-auto bg-red-50 rounded-2xl p-8 text-center">
        <div class="w-16 h-16 mx-auto mb-4 text-red-500">
          <svg viewBox="0 0 24 24" fill="none" class="w-full h-full">
            <path d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" 
                  stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
          </svg>
        </div>
        <h3 class="text-lg font-medium text-red-800 mb-2">{{ cartStore.error }}</h3>
        <button @click="cartStore.fetchCart" 
                class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500">
          Try Again
        </button>
      </div>

      <div v-else-if="cartStore.items.length" class="flex flex-col lg:flex-row gap-8">
        <!-- Cart Items -->
        <div class="flex-1">
          <div class="bg-white rounded-2xl shadow-sm overflow-hidden">
            <div class="divide-y divide-gray-100">
              <div
                v-for="item in cartStore.items"
                :key="item.productId"
                class="p-6 hover:bg-gray-50 transition-colors"
              >
                <div class="flex gap-6">
                  <!-- Product Image -->
                  <div class="relative group">
                    <router-link :to="`/product/${item.productId}`">
                      <img
                        :src="item.product?.image || '/placeholder.jpg'"
                        :alt="item.product?.title"
                        class="w-32 h-32 rounded-xl object-cover group-hover:scale-105 transition-transform duration-300"
                      />
                      <div class="absolute inset-0 rounded-xl ring-1 ring-inset ring-black/5" />
                    </router-link>
                  </div>

                  <!-- Product Info -->
                  <div class="flex-1">
                    <router-link 
                      :to="`/product/${item.productId}`"
                      class="text-lg font-medium text-gray-900 hover:text-blue-600 transition-colors line-clamp-2"
                    >
                      {{ item.product?.title }}
                    </router-link>

                    <div class="mt-4 flex items-center gap-4">
                      <div class="flex items-center gap-2">
                        <button
                          @click="updateQuantity(item.productId, item.quantity - 1)"
                          class="p-2 rounded-lg border border-gray-200 text-gray-500 hover:bg-gray-100 transition-colors"
                          :disabled="item.quantity <= 1"
                        >
                          <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none">
                            <path d="M5 12h14" stroke="currentColor" stroke-width="2" />
                          </svg>
                        </button>
                        <span class="w-12 text-center">{{ item.quantity }}</span>
                        <button
                          @click="updateQuantity(item.productId, item.quantity + 1)"
                          class="p-2 rounded-lg border border-gray-200 text-gray-500 hover:bg-gray-100 transition-colors"
                          :disabled="item.quantity >= (item.product?.stock || 0)"
                        >
                          <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none">
                            <path d="M12 5v14m-7-7h14" stroke="currentColor" stroke-width="2" />
                          </svg>
                        </button>
                      </div>
                      <button
                        @click="removeFromCart(item.productId)"
                        class="text-sm text-red-600 hover:text-red-700"
                      >
                        Удалить
                      </button>
                    </div>
                  </div>

                  <!-- Price -->
                  <div class="text-right">
                    <div class="text-lg font-bold text-gray-900">
                      {{ formatPrice((item.product?.price || 0) * item.quantity) }} ₽
                    </div>
                    <div class="text-sm text-gray-500">
                      {{ formatPrice(item.product?.price || 0) }} ₽ за шт.
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Order Summary -->
        <div class="lg:w-96">
          <div class="bg-white rounded-2xl shadow-sm p-6">
            <h2 class="text-lg font-medium text-gray-900 mb-4">Итого</h2>
            <div class="space-y-4">
              <div class="flex items-center justify-between">
                <span class="text-gray-500">Товары ({{ cartStore.itemCount }})</span>
                <span class="font-medium">{{ formatPrice(cartStore.total) }} ₽</span>
              </div>
              <div class="flex items-center justify-between">
                <span class="text-gray-500">Доставка</span>
                <span class="font-medium">Бесплатно</span>
              </div>
              <div class="pt-4 border-t border-gray-100">
                <div class="flex items-center justify-between text-lg font-bold">
                  <span>Итого к оплате</span>
                  <span>{{ formatPrice(cartStore.total) }} ₽</span>
                </div>
              </div>
              <button
                class="w-full py-4 bg-blue-600 hover:bg-blue-700 text-white font-medium rounded-xl transition-colors"
                @click="checkout"
              >
                Оформить заказ
              </button>
            </div>
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
              d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.8 19h12.4L17 13"
              stroke="currentColor"
              stroke-width="2"
            />
          </svg>
        </div>
        <h3 class="text-2xl font-bold text-gray-900 mb-2">Корзина пуста</h3>
        <p class="text-gray-500 mb-8">
          Добавляйте товары в корзину, чтобы сделать заказ
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
import { useCartStore } from '../../../stores/useCartStore'
import { useRouter } from 'vue-router'
import { useToast } from '../../../composables/useToast'

const cartStore = useCartStore()
const router = useRouter()
const toast = useToast()

function formatPrice(price: number): string {
  return new Intl.NumberFormat('ru-RU').format(price)
}

function getNoun(number: number, one: string, two: string, five: string): string {
  const cases = [2, 0, 1, 1, 1, 2]
  return [one, two, five][
    number % 100 > 4 && number % 100 < 20 ? 2 : cases[number % 10 < 5 ? number % 10 : 5]
  ]
}

async function updateQuantity(productId: string, quantity: number) {
  await cartStore.updateQuantity(productId, quantity)
}

async function removeFromCart(productId: string) {
  await cartStore.removeFromCart(productId)
}

async function clearCart() {
  await cartStore.clearCart()
}

async function checkout() {
  // TODO: Implement checkout flow
  router.push('/checkout')
}

onMounted(() => {
  cartStore.fetchCart()
})
</script>
