<template>
  <!-- Top Bar -->
  <div class="bg-gradient-to-r from-gray-900 to-gray-800 text-white py-2">
    <div
      class="max-w-7xl mx-auto px-4 text-sm flex flex-col sm:flex-row justify-between items-center gap-2 sm:gap-0"
    >
      <div class="flex items-center space-x-6">
        <span>Доставка по всему Казахстану</span>
        <span class="text-gray-400 hidden sm:inline">|</span>
        <span class="text-yellow-400">Скидки до 70%</span>
      </div>
      <div class="flex space-x-6">
        <a href="#" class="hover:text-yellow-400 transition-colors">Стать продавцом</a>
        <a href="#" class="hover:text-yellow-400 transition-colors">Помощь</a>
      </div>
    </div>
  </div>

  <!-- Main Header -->
  <header class="bg-white shadow-lg sticky top-0 z-40">
    <div class="max-w-7xl mx-auto px-4">
      <div class="h-20 flex items-center justify-between gap-4 lg:gap-8">
        <!-- Logo & Catalog -->
        <div class="flex items-center gap-4 lg:gap-8">
          <!-- <img src="/logo.svg" alt="Logo" class="h-8 lg:h-10" /> -->
          <button
            @click="toggleSidebar"
            class="group flex items-center px-4 lg:px-6 py-2.5 lg:py-3 bg-gradient-to-r from-blue-600 to-blue-700 hover:from-blue-700 hover:to-blue-800 text-white rounded-xl transition-all duration-200 shadow-md hover:shadow-lg"
          >
            <svg
              class="w-5 h-5 mr-2 transition-transform group-hover:scale-110"
              viewBox="0 0 24 24"
              fill="none"
            >
              <path
                d="M4 6h16M4 12h16M4 18h16"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
              />
            </svg>
            <span class="font-medium hidden sm:inline">Каталог</span>
          </button>
        </div>

        <!-- Search -->
        <div class="flex-1 max-w-2xl hidden sm:block">
          <div class="relative group">
            <input
              type="text"
              placeholder="Поиск товаров"
              class="w-full pl-12 pr-4 py-3.5 bg-gray-50 border-2 border-gray-100 group-hover:border-blue-100 rounded-xl focus:bg-white focus:border-blue-500 focus:outline-none transition-all duration-200"
            />
            <svg
              class="w-5 h-5 text-gray-400 group-hover:text-gray-600 absolute left-4 top-1/2 -translate-y-1/2 transition-colors duration-200"
              viewBox="0 0 24 24"
              fill="none"
            >
              <path
                d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
              />
            </svg>
          </div>
        </div>

        <!-- Actions -->
        <div class="flex items-center gap-2 lg:gap-6">
          <button class="group flex flex-col items-center">
            <div class="p-2 rounded-full group-hover:bg-blue-50 transition-all duration-200">
              <HeartIcon />
            </div>
            <span
              class="text-xs mt-1 group-hover:text-blue-600 transition-colors duration-200 hidden sm:inline"
              >Избранное</span
            >
          </button>

          <button class="group flex flex-col items-center">
            <div class="p-2 rounded-full group-hover:bg-blue-50 transition-all duration-200">
              <ShoppingCartIcon />
            </div>
            <span
              class="text-xs mt-1 group-hover:text-blue-600 transition-colors duration-200 hidden sm:inline"
              >Корзина</span
            >
          </button>

          <button class="group flex flex-col items-center">
            <div class="p-2 rounded-full group-hover:bg-blue-50 transition-all duration-200">
              <UserIcon />
            </div>
            <span
              class="text-xs mt-1 group-hover:text-blue-600 transition-colors duration-200 hidden sm:inline"
              >Профиль</span
            >
          </button>
        </div>
      </div>
    </div>
  </header>

  <!-- Mobile Search (показывается только на мобильных) -->
  <div class="sm:hidden bg-white border-t border-gray-100 p-4">
    <div class="relative">
      <input
        type="text"
        placeholder="Поиск товаров"
        class="w-full pl-12 pr-4 py-3 bg-gray-50 border-2 border-gray-100 rounded-xl focus:bg-white focus:border-blue-500 focus:outline-none"
      />
      <svg
        class="w-5 h-5 text-gray-400 absolute left-4 top-1/2 -translate-y-1/2"
        viewBox="0 0 24 24"
        fill="none"
      >
        <path
          d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
        />
      </svg>
    </div>
  </div>

  <!-- Catalog Sidebar -->
  <transition name="sidebar">
    <div v-if="sidebarOpen" class="fixed inset-0 z-50">
      <!-- Backdrop -->
      <div class="absolute inset-0 bg-black/60 backdrop-blur-sm" @click="closeSidebar"></div>

      <!-- Sidebar Content -->
      <div class="relative w-full sm:w-[800px] h-full bg-white flex">
        <!-- Mobile Back Button -->
        <div
          v-if="activeCategory && isMobile"
          class="absolute top-5 left-4 z-10"
          @click="activeCategory = null"
        >
          <div class="flex items-center gap-2 text-gray-600">
            <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none">
              <path
                d="M15 18l-6-6 6-6"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
              />
            </svg>
            <span>Назад</span>
          </div>
        </div>

        <!-- Primary Categories -->
        <div
          v-show="!activeCategory || !isMobile"
          class="w-full sm:w-[280px] h-full border-r border-gray-100 flex flex-col"
        >
          <div class="px-6 py-5 border-b border-gray-100">
            <h3 class="text-xl font-semibold text-gray-900">Каталог товаров</h3>
          </div>

          <div class="flex-1 overflow-y-auto">
            <div
              v-for="cat in categories"
              :key="cat.name"
              @click="setActiveCategory(cat)"
              class="px-6 py-3.5 flex items-center justify-between cursor-pointer transition-all duration-200"
              :class="{ 'bg-blue-50 text-blue-600': activeCategoryName === cat.name }"
            >
              <div class="flex items-center gap-3">
                <component
                  :is="cat.icon"
                  class="w-6 h-6"
                  :class="activeCategoryName === cat.name ? 'text-blue-600' : 'text-gray-400'"
                />
                <span class="font-medium">{{ cat.name }}</span>
              </div>
              <svg
                class="w-5 h-5"
                :class="activeCategoryName === cat.name ? 'text-blue-600' : 'text-gray-400'"
                viewBox="0 0 24 24"
                fill="none"
              >
                <path
                  d="M9 18l6-6-6-6"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linecap="round"
                />
              </svg>
            </div>
          </div>
        </div>

        <!-- Subcategories Panel -->
        <div
          v-if="activeCategory"
          v-show="activeCategory && (!isMobile || (isMobile && activeCategory))"
          class="flex-1 h-full flex flex-col"
        >
          <div class="px-6 py-5 border-b border-gray-100 flex justify-between items-center">
            <h3 class="text-xl font-semibold text-gray-900">{{ activeCategory.name }}</h3>
            <button
              @click="closeSidebar"
              class="p-2 hover:bg-gray-100 rounded-lg transition-colors"
            >
              <svg class="w-5 h-5 text-gray-500" viewBox="0 0 24 24" fill="none">
                <path
                  d="M6 18L18 6M6 6l12 12"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linecap="round"
                />
              </svg>
            </button>
          </div>

          <div class="flex-1 overflow-y-auto p-6">
            <div class="grid grid-cols-2 gap-x-12 gap-y-8">
              <div
                v-for="(subcategories, section) in activeCategory.sections"
                :key="section"
                class="space-y-4"
              >
                <h4 class="font-medium text-gray-900">{{ section }}</h4>
                <div class="space-y-3">
                  <router-link
                    v-for="sub in subcategories"
                    :key="sub"
                    :to="`/category/${encodeURIComponent(activeCategory.name)}/${encodeURIComponent(sub)}`"
                    class="block text-sm text-gray-600 hover:text-blue-600 transition-colors"
                    @click="closeSidebar"
                  >
                    {{ sub }}
                  </router-link>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed } from 'vue'
import { useWindowSize } from '@vueuse/core'
import ShoppingCartIcon from '@/components/icons/ShoppingCartIcon.vue'
import UserIcon from '@/components/icons/UserIcon.vue'
import HeartIcon from '@/components/icons/HeartIcon.vue'

interface Category {
  name: string
  icon: any
  sections: Record<string, string[]>
}

const categories: Category[] = [
  {
    name: 'Электроника',
    icon: '',
    sections: {
      'Смартфоны и планшеты': ['Смартфоны', 'Планшеты', 'Электронные книги'],

      'Ноутбуки и ПК': ['Ноутбуки', 'Компьютеры', 'Мониторы'],
      Аксессуары: ['Чехлы', 'Зарядные устройства', 'Защитные стекла'],
    },
  },
  {
    name: 'Одежда',
    icon: '',
    sections: {
      Женщинам: ['Платья', 'Блузки', 'Джинсы', 'Юбки'],
      Мужчинам: ['Футболки', 'Рубашки', 'Брюки', 'Джинсы'],
      Детям: ['Для девочек', 'Для мальчиков', 'Для малышей'],
    },
  },
  {
    name: 'Обувь',
    icon: '',
    sections: {
      Мужская: ['Туфли', 'Ботинки', 'Кроссовки'],
      Женская: ['Туфли', 'Ботинки', 'Кроссовки'],
      Детская: ['Туфли', 'Ботинки', 'Кроссовки'],
    },
  },
  {
    name: 'Дом и сад',
    icon: '',
    sections: {
      Мебель: ['Кухня', 'Гостиная', 'Спальня', 'Детская'],

      Текстиль: ['Подушки', 'Одеяла', 'Покрывала', 'Ковры'],
      Освещение: ['Люстры', 'Бра', 'Светильники', 'Лампы'],
      Декор: ['Картины', 'Фотографии', 'Рисунки', 'Декоративные элементы'],
    },
  },
  {
    name: 'Красота',
    icon: '',
    sections: {
      Макияж: ['Тональный крем', 'Помада', 'Тушь для ресниц', 'Контурный карандаш'],

      'Уход за кожей': ['Мыло', 'Тонизирующий скраб', 'Маска', 'Тоник'],
      Парфюмерия: ['Духи', 'Парфюмированные воды', 'Наборы', 'Декоративные средства'],
    },
  },
  {
    name: 'Спорт',
    icon: '',
    sections: {
      Тренажеры: ['Тренажеры для дома', 'Тренажеры для тренировок', 'Тренажеры для фитнеса'],

      Одежда: ['Футболки', 'Шорты', 'Майки', 'Тренировочные костюмы'],
      Питание: ['Протеины', 'Белковые коктейли', 'Гейнеры', 'Биоактивные добавки'],
      Аксессуары: [
        'Спортивные браслеты',
        'Спортивные часы',
        'Спортивные сумки',
        'Спортивные носки',
      ],
    },
  },
]

const { width } = useWindowSize()
const isMobile = computed(() => width.value < 640)

const sidebarOpen = ref(false)
const activeCategory = ref<Category | null>(null)
const activeCategoryName = ref<string | null>(null)

function setActiveCategory(category: Category) {
  activeCategory.value = category
  activeCategoryName.value = category.name
}

function toggleSidebar() {
  sidebarOpen.value = !sidebarOpen.value
  if (!sidebarOpen.value) {
    activeCategory.value = null
    activeCategoryName.value = null
  }
}

function closeSidebar() {
  sidebarOpen.value = false
  activeCategory.value = null
  activeCategoryName.value = null
}

// Обработчик нажатия ESC
function handleEscKey(event: KeyboardEvent) {
  if (event.key === 'Escape' && sidebarOpen.value) {
    closeSidebar()
  }
}

onMounted(() => {
  document.addEventListener('keydown', handleEscKey)
})

onUnmounted(() => {
  document.removeEventListener('keydown', handleEscKey)
})
</script>

<style scoped>
.sidebar-enter-active,
.sidebar-leave-active {
  transition: all 0.3s ease-out;
}

.sidebar-enter-from,
.sidebar-leave-to {
  opacity: 0;
  transform: translateX(-100%);
}
</style>
