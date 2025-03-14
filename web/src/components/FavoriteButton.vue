<!-- web/src/components/FavoriteButton.vue -->
<script setup lang="ts">
import { useFavoritesStore } from '../stores/useFavoritesStore'
import { useAuthStore } from '../stores/useAuthStore'
import { useRouter } from 'vue-router'
import { computed } from 'vue'

defineOptions({
  name: 'FavoriteButton'
})

const props = defineProps<{
  productId: string
}>()

const router = useRouter()
const favoritesStore = useFavoritesStore()
const authStore = useAuthStore()

const isFavorite = computed(() => favoritesStore.isInFavorites(props.productId))

async function toggleFavorite(e: Event) {
  e.preventDefault()
  e.stopPropagation()
  
  if (!authStore.isAuthenticated) {
    router.push('/security/login')
    return
  }

  try {
    if (isFavorite.value) {
      await favoritesStore.removeFromFavorites(props.productId)
    } else {
      await favoritesStore.addToFavorites(props.productId)
    }
  } catch (error) {
    console.error('Failed to toggle favorite:', error)
  }
}
</script>

<template>
  <button
    @click="toggleFavorite"
    class="p-2 rounded-full bg-white/80 hover:bg-white transition-colors"
    :class="{ 'text-red-500': isFavorite }"
    :title="isFavorite ? 'Remove from favorites' : 'Add to favorites'"
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-6 w-6"
      :fill="isFavorite ? 'currentColor' : 'none'"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"
      />
    </svg>
  </button>
</template>