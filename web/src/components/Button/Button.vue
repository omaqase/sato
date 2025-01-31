<template>
  <button
    :class="[
      'flex items-center justify-center gap-2 font-medium transition-all duration-300 rounded-lg shadow-sm focus:outline-none focus-visible:ring-2 focus-visible:ring-offset-2',
      sizeClasses,
      variantClasses,
      disabled || loading ? 'opacity-50 cursor-not-allowed' : 'cursor-pointer',
    ]"
    :disabled="disabled || loading"
    @click="handleClick"
    @blur="handleBlur"
  >
    <span v-if="loading" class="animate-spin">
      <svg class="w-5 h-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
        <circle
          class="opacity-25"
          cx="12"
          cy="12"
          r="10"
          stroke="currentColor"
          stroke-width="4"
        ></circle>
        <path
          class="opacity-75"
          fill="currentColor"
          d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
        ></path>
      </svg>
    </span>
    <slot></slot>
  </button>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  variant?: 'primary' | 'secondary' | 'destructive' | 'outline' | 'ghost' | 'link' | 'icon'
  size?: 'sm' | 'md' | 'lg'
  loading?: boolean
  disabled?: boolean
}>()

const emit = defineEmits<{
  (e: 'click'): void
}>()

const sizeClasses = computed(() => {
  switch (props.size) {
    case 'sm':
      return 'px-3 py-1.5 text-sm'
    case 'lg':
      return 'px-6 py-3 text-lg'
    default:
      return 'px-4 py-2 text-base'
  }
})

const variantClasses = computed(() => {
  switch (props.variant) {
    case 'primary':
      return 'bg-[var(--color-primary)] text-white hover:bg-[var(--color-primary-hover)] focus-visible:ring-[var(--color-primary-focus)] shadow-md hover:shadow-lg'
    case 'secondary':
      return 'bg-[var(--color-secondary)] text-white hover:bg-[var(--color-secondary-hover)] focus-visible:ring-[var(--color-secondary-focus)] shadow-md hover:shadow-lg'
    case 'destructive':
      return 'bg-[var(--color-danger)] text-white hover:bg-[var(--color-danger-hover)] focus-visible:ring-[var(--color-danger-focus)] shadow-md hover:shadow-lg'
    case 'outline':
      return 'border border-[var(--color-primary)] text-[var(--color-primary)] hover:bg-[var(--color-primary-hover)] hover:text-white focus-visible:ring-[var(--color-primary-focus)]'
    case 'ghost':
      return 'text-[var(--color-primary)] hover:bg-[var(--color-primary-hover)] hover:text-white focus-visible:ring-[var(--color-primary-focus)]'
    case 'link':
      return 'text-[var(--color-primary)] underline hover:text-[var(--color-primary-hover)] focus-visible:ring-[var(--color-primary-focus)]'
    case 'icon':
      return 'p-2 bg-transparent text-[var(--color-primary)] hover:bg-[var(--color-primary-hover)] hover:text-white focus-visible:ring-[var(--color-primary-focus)]'
    default:
      return ''
  }
})

const handleClick = () => {
  if (!props.disabled && !props.loading) {
    emit('click')
  }
}

const handleBlur = (event: FocusEvent) => {
  const target = event.target as HTMLButtonElement
  target.blur()
}
</script>
