<template>
  <div>
    <AppHeader />
    <router-view />
    <AppFooter />
  </div>
</template>

<script setup lang="ts">
import AppHeader from '@/modules/public/components/AppHeader.vue'
import AppFooter from '@/components/AppFooter.vue'
</script>
