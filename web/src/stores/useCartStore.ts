import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { cartApi } from '../api/cart'
import type { CartItem } from '../api/cart'

export const useCartStore = defineStore('cart', () => {
  const items = ref<CartItem[]>([])
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  const total = computed(() => {
    return items.value.reduce((sum: number, item: CartItem) => {
      const price = item.product?.price || 0
      return sum + price * item.quantity
    }, 0)
  })

  const itemCount = computed(() => {
    return items.value.reduce((sum: number, item: CartItem) => sum + item.quantity, 0)
  })

  async function fetchCart() {
    try {
      isLoading.value = true
      error.value = null
      const response = await cartApi.getCart()
      items.value = response.items
    } catch (err) {
      console.error('Failed to fetch cart:', err)
      error.value = 'Failed to fetch cart. Please try again later.'
    } finally {
      isLoading.value = false
    }
  }

  async function addToCart(productId: string, quantity: number = 1) {
    try {
      error.value = null
      await cartApi.addToCart(productId, quantity)
      await fetchCart()
    } catch (err) {
      console.error('Failed to add to cart:', err)
      error.value = 'Failed to add to cart. Please try again later.'
      throw err
    }
  }

  async function removeFromCart(productId: string) {
    try {
      error.value = null
      await cartApi.removeFromCart(productId)
      items.value = items.value.filter((item: CartItem) => item.productId !== productId)
    } catch (err) {
      console.error('Failed to remove from cart:', err)
      error.value = 'Failed to remove from cart. Please try again later.'
      throw err
    }
  }

  async function updateQuantity(productId: string, quantity: number) {
    try {
      error.value = null
      if (quantity <= 0) {
        await removeFromCart(productId)
        return
      }
      await cartApi.updateCartItem(productId, quantity)
      const item = items.value.find((item: CartItem) => item.productId === productId)
      if (item) {
        item.quantity = quantity
      }
    } catch (err) {
      console.error('Failed to update quantity:', err)
      error.value = 'Failed to update quantity. Please try again later.'
      throw err
    }
  }

  async function clearCart() {
    try {
      error.value = null
      await cartApi.clearCart()
      items.value = []
    } catch (err) {
      console.error('Failed to clear cart:', err)
      error.value = 'Failed to clear cart. Please try again later.'
      throw err
    }
  }

  return {
    items,
    isLoading,
    error,
    total,
    itemCount,
    fetchCart,
    addToCart,
    removeFromCart,
    updateQuantity,
    clearCart
  }
}, {
  persist: true
}) 