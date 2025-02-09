<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import DefaultLayout from '@/layouts/DefaultLayout.vue'
import DashboardLayout from '@/modules/dashboard/layouts/DashboardLayout.vue'

const route = useRoute()
const layout = computed(() => {
  return route.meta.layout === 'dashboard' ? DashboardLayout : DefaultLayout
})
</script>

<template>
  <component :is="layout">
    <router-view />
  </component>
</template>
