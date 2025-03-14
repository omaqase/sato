import { defineStore } from 'pinia'
import { ref } from 'vue'
import { favoritesApi } from '@/api/favorites'
import { catalogueApi } from '@/api/catalogue'
import type { FavoriteItem } from '@/api/favorites'

export interface FavoriteProduct {
  productId: string
  title: string
  price?: number
  image?: string
  categoryId: string
  sellerId: string
  stock: number
  rating: number
}

export const useFavoritesStore = defineStore('favorites', () => {
  const favorites = ref<FavoriteProduct[]>([])
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  async function fetchFavorites() {
    try {
      isLoading.value = true
      error.value = null
      
      // Get favorite IDs
      const response = await favoritesApi.getFavorites()
      
      if (!response.items?.length) {
        favorites.value = []
        return
      }

      // Fetch product details for each favorite
      const productPromises = response.items.map((item: FavoriteItem) => 
        catalogueApi.getProduct(item.productId)
          .catch(err => {
            console.error(`Failed to fetch product ${item.productId}:`, err)
            return null
          })
      )

      const products = await Promise.all(productPromises)
      favorites.value = products
        .filter((product): product is FavoriteProduct => product !== null)
        .map(product => ({
          ...product,
          price: product.price || 0
        }))
    } catch (err) {
      console.error('Failed to fetch favorites:', err)
      error.value = 'Failed to fetch favorites. Please try again later.'
      favorites.value = []
    } finally {
      isLoading.value = false
    }
  }

  async function addToFavorites(productId: string) {
    try {
      error.value = null
      await favoritesApi.addToFavorites(productId)
      await fetchFavorites()
    } catch (err) {
      console.error('Failed to add to favorites:', err)
      error.value = 'Failed to add to favorites. Please try again later.'
    }
  }

  async function removeFromFavorites(productId: string) {
    try {
      error.value = null
      await favoritesApi.removeFromFavorites(productId)
      favorites.value = favorites.value.filter(f => f.productId !== productId)
    } catch (err) {
      console.error('Failed to remove from favorites:', err)
      error.value = 'Failed to remove from favorites. Please try again later.'
    }
  }

  function isInFavorites(productId: string): boolean {
    return favorites.value.some(f => f.productId === productId)
  }

  return {
    favorites,
    isLoading,
    error,
    fetchFavorites,
    addToFavorites,
    removeFromFavorites,
    isInFavorites
  }
})