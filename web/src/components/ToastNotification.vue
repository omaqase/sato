<template>
  <Transition
    enter-active-class="transition-all duration-300 ease-out"
    enter-from-class="translate-y-4 opacity-0"
    enter-to-class="translate-y-0 opacity-100"
    leave-active-class="transition-all duration-300 ease-in"
    leave-from-class="translate-y-0 opacity-100"
    leave-to-class="translate-y-4 opacity-0"
  >
    <div
      v-if="isVisible"
      class="fixed bottom-4 right-4 bg-white rounded-xl shadow-lg border border-gray-100 p-4 hover:shadow-xl"
    >
      <div class="flex items-center gap-3">
        <div class="p-2 bg-blue-50 rounded-lg">
          <svg
            class="w-5 h-5 text-blue-600"
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            />
          </svg>
        </div>
        <p class="text-gray-700">{{ message }}</p>
      </div>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'

const props = defineProps<{
  message: string
  duration?: number
}>()

const isVisible = ref(false)
let timeout: NodeJS.Timeout

onMounted(() => {
  isVisible.value = true
  timeout = setTimeout(() => {
    isVisible.value = false
  }, props.duration || 3000)
})

onUnmounted(() => {
  clearTimeout(timeout)
})
</script>