<!-- web/src/views/FavoritesView.vue -->
<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useFavoritesStore } from '@/stores/useFavoritesStore'
import { useProductsStore } from '@/stores/useProductsStore'
import ProductCard from '@/components/ProductCard.vue'

const favoritesStore = useFavoritesStore()
const productsStore = useProductsStore()
const isLoading = ref(true)

onMounted(async () => {
  try {
    await favoritesStore.fetchFavorites()
    if (favoritesStore.favorites.length > 0) {  
      await productsStore.fetchProductsByIds(favoritesStore.favorites)
    }
  } catch (error) {
    console.error('Failed to load favorites:', error)
  } finally {
    isLoading.value = false
  }
})
</script>

<template>
  <div class="container mx-auto px-4 py-8">
    <h1 class="text-2xl font-bold mb-6">My Favorites</h1>
    
    <div v-if="isLoading || favoritesStore.loading || productsStore.loading" 
         class="text-center py-8">
      Loading...
    </div>
    
    <div v-else-if="favoritesStore.error || productsStore.error" 
         class="text-center text-red-500 py-8">
      {{ favoritesStore.error || productsStore.error }}
    </div>
    
    <div v-else-if="favoritesStore.favorites.length === 0" 
         class="text-center py-8">
      You haven't added any products to your favorites yet.
    </div>
    
    <div v-else
         class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
      <ProductCard
        v-for="productId in favoritesStore.favorites"
        :key="productId"
        :product="productsStore.getProductById(productId)"
      />
    </div>
  </div>
</template>