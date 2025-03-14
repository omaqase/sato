<template>
  <div :class="classes">
    <slot />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  cols?: number | { [key: string]: number }
  gap?: number | string
  flow?: 'row' | 'col'
  align?: 'start' | 'center' | 'end' | 'stretch'
  justify?: 'start' | 'center' | 'end' | 'between' | 'around'
}

const props = withDefaults(defineProps<Props>(), {
  cols: 1,
  gap: 4,
  flow: 'row',
  align: 'stretch',
  justify: 'start'
})

const classes = computed(() => {
  const baseClasses = ['grid']
  
  // Handle gap
  if (typeof props.gap === 'number') {
    baseClasses.push(`gap-${props.gap}`)
  } else {
    baseClasses.push(`gap-${props.gap}`)
  }

  // Handle columns
  if (typeof props.cols === 'number') {
    baseClasses.push(`grid-cols-${props.cols}`)
  } else {
    Object.entries(props.cols).forEach(([breakpoint, cols]) => {
      if (breakpoint === 'default') {
        baseClasses.push(`grid-cols-${cols}`)
      } else {
        baseClasses.push(`${breakpoint}:grid-cols-${cols}`)
      }
    })
  }

  // Handle alignment
  baseClasses.push(`items-${props.align}`)
  baseClasses.push(`justify-${props.justify}`)

  // Handle flow direction
  if (props.flow === 'col') {
    baseClasses.push('grid-flow-col')
  }

  return baseClasses.join(' ')
})
</script>