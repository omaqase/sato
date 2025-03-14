<template>
  <div class="relative group bg-white rounded-2xl shadow-sm overflow-hidden">
    <FavoriteButton
      :product-id="product.productId"
      class="absolute top-2 right-2 z-10"
    />
    
    <router-link :to="`/product/${product.productId}`" class="block">
      <!-- Product Image -->
      <div class="relative aspect-square">
        <img
          :src="product.image || '/placeholder.jpg'"
          :alt="product.title"
          class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
        />
        <div class="absolute inset-0 ring-1 ring-inset ring-black/5" />
      </div>

      <!-- Product Info -->
      <div class="p-4">
        <h3 class="text-lg font-medium text-gray-900 line-clamp-2 group-hover:text-blue-600 transition-colors">
          {{ product.title }}
        </h3>
        
        <div class="mt-2 flex items-center justify-between">
          <div class="text-lg font-bold text-gray-900">
            {{ formatPrice(product.price || 0) }} ₽
          </div>
          <div class="text-sm text-gray-500">
            {{ product.stock }} шт.
          </div>
        </div>
      </div>
    </router-link>

    <!-- Actions Slot -->
    <div class="p-4 pt-0">
      <slot name="actions">
        <button
          @click="addToCart"
          class="w-full flex items-center justify-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          :disabled="isAddingToCart"
        >
          <template v-if="isAddingToCart">
            <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            Добавление...
          </template>
          <template v-else>
            В корзину
          </template>
        </button>
      </slot>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { useFavoritesStore } from '../stores/useFavoritesStore'
import { useAuthStore } from '../stores/useAuthStore'
import { useCartStore } from '../stores/useCartStore'
import { useRouter } from 'vue-router'
import { useToast } from '../composables/useToast'
import FavoriteButton from './FavoriteButton.vue'

interface Product {
  productId: string
  title: string
  price?: number
  image?: string
  categoryId: string
  sellerId: string
  stock: number
  rating: number
}

const props = defineProps<{
  product: Product
}>()

const router = useRouter()
const favoritesStore = useFavoritesStore()
const authStore = useAuthStore()
const cartStore = useCartStore()
const toast = useToast()

const isAddingToCart = ref(false)

const isInFavorites = computed(() => favoritesStore.isInFavorites(props.product.productId))

async function addToCart() {
  if (!authStore.isAuthenticated) {
    router.push('/security/login')
    return
  }

  try {
    isAddingToCart.value = true
    await cartStore.addToCart(props.product.productId)
    toast.showToast({ message: 'Товар добавлен в корзину' })
  } catch (error) {
    console.error('Failed to add to cart:', error)
    toast.showToast({ message: 'Не удалось добавить товар в корзину' })
  } finally {
    isAddingToCart.value = false
  }
}

function formatPrice(price: number): string {
  return new Intl.NumberFormat('ru-RU').format(price)
}
</script>