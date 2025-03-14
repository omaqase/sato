<template>
  <div :class="classes">
    <slot />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  direction?: 'row' | 'col'
  gap?: number | string
  align?: 'start' | 'center' | 'end' | 'stretch' | 'baseline'
  justify?: 'start' | 'center' | 'end' | 'between' | 'around'
  wrap?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  direction: 'col',
  gap: 4,
  align: 'stretch',
  justify: 'start',
  wrap: false
})

const classes = computed(() => {
  const baseClasses = ['flex']

  // Handle direction
  baseClasses.push(props.direction === 'col' ? 'flex-col' : 'flex-row')

  // Handle gap
  if (typeof props.gap === 'number') {
    baseClasses.push(`gap-${props.gap}`)
  } else {
    baseClasses.push(`gap-${props.gap}`)
  }

  // Handle alignment
  baseClasses.push(`items-${props.align}`)
  baseClasses.push(`justify-${props.justify}`)

  // Handle wrap
  if (props.wrap) {
    baseClasses.push('flex-wrap')
  }

  return baseClasses.join(' ')
})
</script>