<template>
  <div :class="['relative', wrapperClasses]">
    <!-- Левая иконка -->
    <span
      v-if="iconLeft"
      class="absolute left-3 top-1/2 transform -translate-y-1/2 text-[var(--color-gray-dark)] pointer-events-none"
    >
      <slot name="icon-left"></slot>
    </span>

    <!-- Поле ввода -->
    <input
      :class="[
        'block w-full px-4 py-3 text-base font-medium transition-all duration-300 rounded-lg border focus:outline-none',
        inputClasses,
        disabled ? 'opacity-50 cursor-not-allowed' : 'cursor-text',
        iconLeft && 'pl-10', // Отступ для левой иконки
        iconRight && 'pr-10', // Отступ для правой иконки
      ]"
      :type="type"
      :placeholder="placeholder"
      :value="modelValue"
      :disabled="disabled"
      @input="$emit('update:modelValue', ($event.target as HTMLInputElement).value)"
    />

    <!-- Правая иконка -->
    <span
      v-if="iconRight"
      class="absolute right-3 top-1/2 transform -translate-y-1/2 text-[var(--color-gray-dark)] pointer-events-none"
    >
      <slot name="icon-right"></slot>
    </span>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  modelValue?: string | number
  type?: string
  placeholder?: string
  variant?: 'primary' | 'secondary'
  size?: 'sm' | 'md' | 'lg'
  disabled?: boolean
  iconLeft?: boolean
  iconRight?: boolean
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', value: string | number): void
}>()

const sizeClasses = computed(() => {
  switch (props.size) {
    case 'sm':
      return 'px-3 py-2 text-sm'
    case 'lg':
      return 'px-6 py-4 text-lg'
    default:
      return 'px-4 py-3 text-base'
  }
})

const variantClasses = computed(() => {
  switch (props.variant) {
    case 'primary':
      return 'border-[var(--color-primary)] bg-[var(--color-gray-light)] text-[var(--color-text)] focus:border-[var(--color-primary)] focus:ring-2 focus:ring-[var(--color-primary-focus)]'
    case 'secondary':
      return 'border-[var(--color-secondary)] bg-[var(--color-gray-light)] text-[var(--color-text)] focus:border-[var(--color-secondary)] focus:ring-2 focus:ring-[var(--color-secondary-focus)]'
    default:
      return 'border-[var(--color-gray-dark)] bg-[var(--color-gray-light)] text-[var(--color-text)] focus:border-[var(--color-primary)] focus:ring-2 focus:ring-[var(--color-primary-focus)]'
  }
})

const wrapperClasses = computed(() => {
  return 'w-full'
})

const inputClasses = computed(() => {
  return `${sizeClasses.value} ${variantClasses.value}`
})
</script>
