<script setup lang="ts">
import { RouterView } from 'vue-router'
import AppHeader from './modules/public/components/AppHeader.vue'
import AppFooter from './components/AppFooter.vue'
</script>

<template>
  <div class="flex flex-col min-h-screen">
    <AppHeader />
    <main class="flex-1">
      <router-view></router-view>
    </main>
    <AppFooter />
  </div>
</template>
