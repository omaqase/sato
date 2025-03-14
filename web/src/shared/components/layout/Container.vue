<template>
  <div class="container mx-auto px-4 sm:px-6 lg:px-8" :class="classes">
    <slot />
  </div>
</template>

<script setup lang="ts">
interface Props {
  size?: 'sm' | 'md' | 'lg' | 'xl' | '2xl' | 'full'
  fluid?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  size: 'xl',
  fluid: false
})

const classes = computed(() => ({
  'max-w-screen-sm': props.size === 'sm',
  'max-w-screen-md': props.size === 'md',
  'max-w-screen-lg': props.size === 'lg',
  'max-w-screen-xl': props.size === 'xl',
  'max-w-screen-2xl': props.size === '2xl',
  'max-w-full': props.size === 'full' || props.fluid
}))
</script>