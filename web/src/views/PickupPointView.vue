<template>
  <div class="min-h-screen bg-gradient-to-b from-gray-50 to-white pb-12">
    <div class="bg-gradient-to-r from-blue-600 to-blue-700 text-white">
      <div class="max-w-7xl mx-auto px-4 py-16">
        <div class="max-w-2xl">
          <h1 class="text-4xl font-bold mb-6">Партнерский пункт выдачи</h1>
          <p class="text-xl text-blue-100">
            Зарабатывайте на выдаче заказов SATO в вашем помещении
          </p>
          <button
            class="mt-8 px-8 py-4 bg-white text-blue-600 rounded-xl font-medium hover:bg-blue-50 transition-colors"
          >
            Стать партнером
          </button>
        </div>
      </div>
    </div>

    <div class="max-w-7xl mx-auto px-4 py-12">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8 mb-16">
        <div
          v-for="benefit in benefits"
          :key="benefit.id"
          class="bg-white p-8 rounded-2xl hover:shadow-xl transition-all duration-300"
        >
          <div class="w-16 h-16 rounded-xl bg-blue-100 flex items-center justify-center mb-6">
            <component :is="benefit.icon" class="w-8 h-8 text-blue-600" />
          </div>
          <h3 class="text-xl font-semibold text-gray-900 mb-4">{{ benefit.title }}</h3>
          <p class="text-gray-600">{{ benefit.description }}</p>
        </div>
      </div>

      <div class="bg-white rounded-2xl p-8 mb-16">
        <h2 class="text-2xl font-bold text-gray-900 mb-8">Требования к помещению</h2>
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
            <h2 class="text-2xl font-bold text-gray-900 mb-4">Начните зарабатывать</h2>
            <p class="text-gray-600 mb-6">
              Оставьте заявку, и наш менеджер расскажет подробнее об условиях сотрудничества
            </p>
            <button
              class="px-6 py-3 bg-blue-600 text-white rounded-xl hover:bg-blue-700 transition-colors"
            >
              Оставить заявку
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import ChartHistogramIcon from '@/components/icons/ChartHistogramIcon.vue'
import CoinsIcon from '@/components/icons/CoinsIcon.vue'
import DollarIcon from '@/components/icons/DollarIcon.vue'

const benefits = [
  {
    id: 1,
    title: 'Стабильный доход',
    description: 'От 150 000 ₸ в месяц за выдачу заказов',
    icon: CoinsIcon,
  },
  {
    id: 2,
    title: 'Минимум затрат',
    description: 'Не требуется специальное оборудование',
    icon: DollarIcon,
  },
  {
    id: 3,
    title: 'Новые клиенты',
    description: 'Увеличение трафика в вашей точке',
    icon: ChartHistogramIcon,
  },
]

const requirements = [
  {
    id: 1,
    title: 'Расположение',
    description: 'В проходном месте с удобной транспортной доступностью',
    icon: 'LocationIcon',
  },
  {
    id: 2,
    title: 'Площадь',
    description: 'От 15 кв.м для хранения заказов',
    icon: 'SquareIcon',
  },
  {
    id: 3,
    title: 'График работы',
    description: 'Не менее 8 часов в день, 6 дней в неделю',
    icon: 'ClockIcon',
  },
  {
    id: 4,
    title: 'Интернет',
    description: 'Стабильное подключение к интернету',
    icon: 'WifiIcon',
  },
]
</script>
