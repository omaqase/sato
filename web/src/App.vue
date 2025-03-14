<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import DefaultLayout from '@/layouts/DefaultLayout.vue'
import DashboardLayout from '@/modules/dashboard/layouts/DashboardLayout.vue'
import { useRoleStore } from '@/stores/useRoleStore'

const route = useRoute()
const layout = computed(() => {
  return route.meta.layout === 'dashboard' ? DashboardLayout : DefaultLayout
})

const roleStore = useRoleStore()

onMounted(async () => {
  await roleStore.checkAdminRole()
})
</script>

<template>
  <component :is="layout">
    <router-view />
  </component>
</template>
