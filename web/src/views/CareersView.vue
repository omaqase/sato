<template>
  <div class="min-h-screen bg-gradient-to-b from-gray-50 to-white pb-12">
    <div class="max-w-7xl mx-auto px-4 py-12">
      <div class="text-center mb-12">
        <h1 class="text-4xl font-bold text-gray-900 mb-4">Карьера в MarketPlace</h1>
        <p class="text-xl text-gray-600">Присоединяйтесь к нашей команде</p>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-12">
        <div v-for="department in departments" :key="department.id" class="text-center">
          <div
            class="w-16 h-16 mx-auto mb-4 rounded-2xl bg-blue-100 flex items-center justify-center"
          >
            <component :is="department.icon" class="w-8 h-8 text-blue-600" />
          </div>
          <h3 class="font-medium text-gray-900 mb-2">{{ department.name }}</h3>
          <p class="text-sm text-gray-500">{{ department.count }} вакансий</p>
        </div>
      </div>

      <div class="space-y-4">
        <div
          v-for="job in jobs"
          :key="job.id"
          class="bg-white rounded-2xl p-6 hover:shadow-xl transition-all duration-300"
        >
          <div class="flex items-center justify-between mb-4">
            <div>
              <h3 class="text-xl font-medium text-gray-900">{{ job.title }}</h3>
              <p class="text-gray-500">{{ job.department }}</p>
            </div>
            <button
              class="px-6 py-2 bg-blue-600 text-white rounded-xl hover:bg-blue-700 transition-colors"
              @click="showFeatureModal"
            >
              Откликнуться
            </button>
          </div>
          <div class="flex items-center gap-6 text-sm text-gray-500">
            <div class="flex items-center gap-2">
              <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none">
                <path
                  d="M21 10c0 7-9 13-9 13s-9-6-9-13a9 9 0 0118 0z"
                  stroke="currentColor"
                  stroke-width="2"
                />
                <circle cx="12" cy="10" r="3" stroke="currentColor" stroke-width="2" />
              </svg>
              {{ job.location }}
            </div>
            <div class="flex items-center gap-2">
              <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none">
                <path
                  d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linecap="round"
                />
              </svg>
              {{ job.type }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <BaseModal ref="modal" />
</template>

<script setup lang="ts">
import BaseModal from '@/components/BaseModal.vue'
import CodeWindowIcon from '@/components/icons/CodeWindowIcon.vue'
import MegaphoneAnnouncementIcon from '@/components/icons/MegaphoneAnnouncementIcon.vue'
import ShoppingCartIcon from '@/components/icons/ShoppingCartIcon.vue'
import { ref } from 'vue'

const departments = [
  { id: 1, name: 'Разработка', count: 5, icon: CodeWindowIcon },
  { id: 2, name: 'Маркетинг', count: 3, icon: MegaphoneAnnouncementIcon },
  { id: 3, name: 'Продажи', count: 4, icon: ShoppingCartIcon },
]

const jobs = [
  {
    id: 1,
    title: 'Frontend Developer',
    department: 'Разработка',
    location: 'Астана',
    type: 'Полный день',
  },
  {
    id: 2,
    title: 'Marketing Manager',
    department: 'Маркетинг',
    location: 'Удаленно',
    type: 'Полный день',
  },
]

const modal = ref<InstanceType<typeof BaseModal> | null>(null)

const showFeatureModal = () => {
  modal.value?.openModal(
    'Функция в разработке',
    'Мы работаем над этой функцией. Она будет доступна в ближайшее время!',
  )
}
</script>
