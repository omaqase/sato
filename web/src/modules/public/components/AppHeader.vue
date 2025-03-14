<template>
  <div class="bg-gradient-to-r from-gray-900 to-gray-800 text-white py-2">
    <div class="max-w-7xl mx-auto px-4 text-sm flex flex-col sm:flex-row justify-between items-center gap-2 sm:gap-0">
      <div class="flex items-center space-x-6">
        <span>Доставка по всему Казахстану</span>
        <span class="text-gray-400 hidden sm:inline">|</span>
        <span class="text-yellow-400">Скидки до 70%</span>
      </div>
      <div class="flex space-x-6">
        <router-link to="/seller" class="hover:text-yellow-400 transition-colors">Стать продавцом</router-link>
        <router-link to="/help" class="hover:text-yellow-400 transition-colors">Помощь</router-link>
      </div>
    </div>
  </div>

  <header class="bg-white shadow-sm sticky top-0 z-40">
    <div class="max-w-7xl mx-auto px-4">
      <div class="h-16 sm:h-20 flex items-center justify-between gap-4">
        <div class="flex items-center gap-4">
          <button
            @click="toggleSidebar"
            class="inline-flex items-center px-3 sm:px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg transition-all"
          >
            <ListIcon class="w-5 h-5" />
            <span class="ml-2 font-medium hidden sm:inline">Каталог</span>
          </button>
        </div>

        <div class="flex-1 max-w-xl hidden sm:block">
          <div class="relative">
            <input
              type="text"
              placeholder="Поиск товаров"
              class="w-full pl-10 pr-4 py-2.5 bg-gray-50 border border-gray-200 rounded-lg focus:outline-none focus:border-blue-500"
            />
            <SettingsIcon class="w-5 h-5 text-gray-400 absolute left-3 top-1/2 -translate-y-1/2" />
          </div>
        </div>

        <div class="flex items-center gap-1 sm:gap-4">
          <router-link to="/favorites" class="p-2 hover:bg-gray-100 rounded-lg hidden sm:block">
            <HeartIcon class="w-6 h-6 text-gray-600" />
          </router-link>
          <router-link to="/cart" class="p-2 hover:bg-gray-100 rounded-lg">
            <ShoppingCartIcon class="w-6 h-6 text-gray-600" />
          </router-link>
          <router-link to="/account" class="p-2 hover:bg-gray-100 rounded-lg">
            <UserIcon class="w-6 h-6 text-gray-600" />
          </router-link>
        </div>
      </div>
    </div>

    <div class="sm:hidden border-t border-gray-100 px-4 py-2">
      <div class="relative">
        <input
          type="text"
          placeholder="Поиск товаров"
          class="w-full pl-10 pr-4 py-2.5 bg-gray-50 border border-gray-200 rounded-lg focus:outline-none focus:border-blue-500"
        />
        <SettingsIcon class="w-5 h-5 text-gray-400 absolute left-3 top-1/2 -translate-y-1/2" />
      </div>
    </div>
  </header>

  <Transition name="drawer">
    <div v-if="isOpen" class="fixed inset-0 z-50">
      <div class="fixed inset-0 bg-black/50 backdrop-blur-sm" @click="closeDrawer" />
      <div class="fixed inset-y-0 left-0 w-full sm:w-[400px] bg-white shadow-xl">
        <div class="h-full flex flex-col">
          <div class="px-4 py-4 border-b border-gray-100 flex items-center justify-between">
            <h2 class="text-lg font-semibold text-gray-900">Категории</h2>
            <button @click="closeDrawer" class="p-2 hover:bg-gray-100 rounded-lg">
              <SettingsIcon class="w-5 h-5 text-gray-500" />
            </button>
          </div>

          <div class="flex-1 overflow-y-auto">
            <div class="divide-y divide-gray-100">
              <div v-for="category in categories" :key="category.name" class="group">
                <button
                  class="w-full px-4 py-3.5 flex items-center justify-between hover:bg-gray-50"
                  @click="toggleCategory(category)"
                >
                  <div class="flex items-center gap-3">
                    <component
                      :is="category.icon || 'div'"
                      class="w-5 h-5 text-gray-400 group-hover:text-blue-600"
                    />
                    <span class="font-medium text-gray-900">{{ category.name }}</span>
                  </div>
                </button>

                <div v-show="openCategories.includes(category.name)" class="bg-gray-50 px-4 py-3 space-y-6">
                  <div v-for="(items, section) in category.sections" :key="section" class="space-y-2">
                    <h3 class="text-sm font-medium text-gray-900">{{ section }}</h3>
                    <div class="space-y-1">
                      <router-link
                        v-for="item in items"
                        :key="item"
                        :to="`/category/${category.name}/${item}`"
                        class="block px-3 py-2 text-sm text-gray-600 hover:text-blue-600 rounded-lg hover:bg-white"
                        @click="closeDrawer"
                      >
                        {{ item }}
                      </router-link>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed } from 'vue'
import { useWindowSize } from '@vueuse/core'
import SettingsIcon from '@/components/icons/SettingsIcon.vue'
import HeartIcon from '@/components/icons/HeartIcon.vue'
import ShoppingCartIcon from '@/components/icons/ShoppingCartIcon.vue'
import UserIcon from '@/components/icons/UserIcon.vue'
import { categoryApi, type Category as CategoryType } from '@/api/category'

interface CategorySection {
  name: string
  icon: unknown
  sections: Record<string, string[]>
}

const { width } = useWindowSize()
const isMobile = computed(() => width.value < 640)

const isOpen = ref(false)
const openCategories = ref<string[]>([])
const categories = ref<CategorySection[]>([])

async function fetchCategories() {
  try {
    const response = await categoryApi.list()
    categories.value = response.categories.map(category => ({
      name: category.title,
      icon: '',
      sections: {
        [category.title]: [category.slug]
      }
    }))
  } catch (error) {
    console.error('Failed to fetch categories:', error)
  }
}

function toggleDrawer() {
  isOpen.value = !isOpen.value
}

function closeDrawer() {
  isOpen.value = false
  openCategories.value = []
}

function toggleCategory(category: CategorySection) {
  const index = openCategories.value.indexOf(category.name)
  if (index === -1) {
    openCategories.value.push(category.name)
  } else {
    openCategories.value.splice(index, 1)
  }
}

onMounted(() => {
  fetchCategories()
})
</script>

<style scoped>
.drawer-enter-active,
.drawer-leave-active {
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
}

.drawer-enter-from,
.drawer-leave-to {
  opacity: 0;
  transform: translateX(-10px);
}
</style>
