<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  rating: number
  maxRating?: number
  size?: 'sm' | 'md' | 'lg'
}>()

const stars = computed(() => {
  const max = props.maxRating || 5
  const fullStars = Math.floor(props.rating)
  const hasHalfStar = props.rating % 1 >= 0.5
  
  return {
    full: fullStars,
    half: hasHalfStar ? 1 : 0,
    empty: max - fullStars - (hasHalfStar ? 1 : 0)
  }
})

const sizeClass = computed(() => {
  switch (props.size) {
    case 'sm': return 'w-3 h-3'
    case 'lg': return 'w-6 h-6'
    default: return 'w-4 h-4' // md is default
  }
})
</script>

<template>
  <div class="flex">
    <!-- Full stars -->
    <svg 
      v-for="i in stars.full" 
      :key="`full-${i}`" 
      :class="sizeClass"
      viewBox="0 0 24 24" 
      fill="currentColor"
      class="text-yellow-400"
    >
      <path d="M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z" />
    </svg>
    
    <!-- Half star -->
    <svg 
      v-if="stars.half" 
      :class="sizeClass"
      viewBox="0 0 24 24" 
      class="text-yellow-400"
    >
      <defs>
        <linearGradient id="half-star" x1="0" x2="100%" y1="0" y2="0">
          <stop offset="50%" stop-color="currentColor" />
          <stop offset="50%" stop-color="#D1D5DB" />
        </linearGradient>
      </defs>
      <path d="M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z" fill="url(#half-star)" />
    </svg>
    
    <!-- Empty stars -->
    <svg 
      v-for="i in stars.empty" 
      :key="`empty-${i}`" 
      :class="sizeClass"
      viewBox="0 0 24 24" 
      fill="currentColor"
      class="text-gray-300"
    >
      <path d="M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z" />
    </svg>
  </div>
</template> 