<template>
  <div class="min-h-screen bg-gradient-to-b from-gray-50 to-white pb-12">
    <div class="bg-gradient-to-r from-blue-600 to-blue-700 text-white">
      <div class="max-w-7xl mx-auto px-4 py-4">
        <div class="flex flex-col md:flex-row items-center gap-8">
          <div class="flex-1">
            <h1 class="text-4xl font-bold mb-6">Станьте перевозчиком MarketPlace</h1>
            <p class="text-xl text-blue-100">
              Зарабатывайте на доставке заказов вместе с крупнейшим маркетплейсом России
            </p>
            <button
              class="mt-8 px-8 py-4 bg-white text-blue-600 rounded-xl font-medium hover:bg-blue-50 transition-colors"
              @click="showFeatureInDevelopment"
            >
              Стать партнером
            </button>
          </div>
          <div class="w-full md:w-1/2 lg:w-1/3">
            <img :src="deliveryPackage" alt="Доставка" class="w-full h-auto object-contain" />
          </div>
        </div>
      </div>
    </div>

    <div class="max-w-7xl mx-auto px-4 py-12">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8 mb-16">
        <div
          v-for="advantage in advantages"
          :key="advantage.id"
          class="bg-white p-8 rounded-2xl hover:shadow-xl transition-all duration-300"
        >
          <div class="w-16 h-16 rounded-xl bg-blue-100 flex items-center justify-center mb-6">
            <component :is="advantage.icon" class="w-8 h-8 text-blue-600" />
          </div>
          <h3 class="text-xl font-semibold text-gray-900 mb-4">{{ advantage.title }}</h3>
          <p class="text-gray-600">{{ advantage.description }}</p>
        </div>
      </div>

      <div class="bg-white rounded-2xl p-8 mb-16">
        <h2 class="text-2xl font-bold text-gray-900 mb-8">Требования к перевозчикам</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
          <div v-for="requirement in requirements" :key="requirement.id">
            <div class="flex items-start gap-4">
              <div class="p-3 rounded-xl bg-blue-100">
                <component :is="requirement.icon" class="w-6 h-6 text-blue-600" />
              </div>
              <div>
                <h3 class="font-medium text-gray-900 mb-2">{{ requirement.title }}</h3>
                <p class="text-sm text-gray-500">{{ requirement.description }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-blue-50 rounded-2xl p-8">
        <div class="flex flex-col md:flex-row items-center gap-8">
          <div class="flex-1">
            <h2 class="text-2xl font-bold text-gray-900 mb-4">Готовы начать?</h2>
            <p class="text-gray-600 mb-6">
              Заполните анкету, и наш менеджер свяжется с вами для обсуждения деталей сотрудничества
            </p>
            <button
              class="px-6 py-3 bg-blue-600 text-white rounded-xl hover:bg-blue-700 transition-colors"
              @click="showFeatureInDevelopment"
            >
              Оставить заявку
            </button>
          </div>
          <!-- <img src="/images/transport/delivery.jpg" alt="Доставка" class="w-full md:w-96 rounded-xl" /> -->
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import deliveryPackage from '@/assets/delivery_package.png'
import ChartHistogramIcon from '@/components/icons/ChartHistogramIcon.vue'
import CoinsIcon from '@/components/icons/CoinsIcon.vue'
import DocumentsSignedIcon from '@/components/icons/DocumentsSignedIcon.vue'
import FlameIcon from '@/components/icons/FlameIcon.vue'
import ListCheckIcon from '@/components/icons/ListCheckIcon.vue'
import TruckIcon from '@/components/icons/TruckIcon.vue'
import UserExperienceIcon from '@/components/icons/UserExperienceIcon.vue'
import { useToast } from "@/composables/useToast"

const { showFeatureInDevelopment } = useToast()

const advantages = [
  {
    id: 1,
    title: 'Стабильный поток заказов',
    description: 'Более 100 000 отправлений ежедневно по всей России',
    icon: ChartHistogramIcon,
  },
  {
    id: 2,
    title: 'Гибкие условия',
    description: 'Выбирайте удобные маршруты и график работы',
    icon: ListCheckIcon,
  },
  {
    id: 3,
    title: 'Быстрые выплаты',
    description: 'Еженедельные выплаты без задержек',
    icon: CoinsIcon,
  },
]

const requirements = [
  {
    id: 1,
    title: 'Транспорт',
    description: 'Собственный или арендованный автомобиль в исправном состоянии',
    icon: TruckIcon,
  },
  {
    id: 2,
    title: 'Документы',
    description: 'ИП или ООО, лицензия на перевозку грузов',
    icon: DocumentsSignedIcon,
  },
  {
    id: 3,
    title: 'Опыт',
    description: 'Опыт в сфере грузоперевозок от 1 года',
    icon: UserExperienceIcon,
  },
  {
    id: 4,
    title: 'Страхование',
    description: 'Страхование ответственности перевозчика',
    icon: FlameIcon,
  },
]
</script>
