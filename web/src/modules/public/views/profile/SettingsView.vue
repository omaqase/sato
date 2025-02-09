<template>
  <div class="flex-1">
    <div class="bg-white rounded-2xl shadow-sm overflow-hidden">
      <div class="p-6 border-b border-gray-100">
        <h2 class="text-lg font-semibold text-gray-900">Настройки профиля</h2>
      </div>

      <div class="p-6 space-y-8">
        <!-- Personal Info -->
        <div>
          <h3 class="text-base font-medium text-gray-900 mb-4">Личные данные</h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div v-for="field in personalFields" :key="field.id">
              <label :for="field.id" class="block text-sm font-medium text-gray-700 mb-2">
                {{ field.label }}
              </label>
              <input
                :type="field.type"
                :id="field.id"
                :value="field.value"
                class="w-full px-4 py-2.5 rounded-xl border border-gray-200 focus:ring-2 focus:ring-blue-100 focus:border-blue-600 transition-colors"
                :placeholder="field.placeholder"
              />
            </div>
          </div>
        </div>

        <!-- Notifications Settings -->
        <div class="pt-8 border-t">
          <h3 class="text-base font-medium text-gray-900 mb-4">Уведомления</h3>
          <div class="space-y-4">
            <div
              v-for="setting in notificationSettings"
              :key="setting.id"
              class="flex items-center justify-between py-3"
            >
              <div>
                <div class="font-medium text-gray-900">{{ setting.title }}</div>
                <div class="text-sm text-gray-500">{{ setting.description }}</div>
              </div>
              <button
                @click="setting.enabled = !setting.enabled"
                class="relative inline-flex h-6 w-11 items-center rounded-full transition-colors"
                :class="setting.enabled ? 'bg-blue-600' : 'bg-gray-200'"
              >
                <span
                  class="inline-block h-4 w-4 transform rounded-full bg-white transition-transform"
                  :class="setting.enabled ? 'translate-x-6' : 'translate-x-1'"
                />
              </button>
            </div>
          </div>
        </div>

        <!-- Save Button -->
        <div class="pt-8 border-t">
          <button
            class="w-full py-3 bg-blue-600 hover:bg-blue-700 text-white font-medium rounded-xl transition-colors"
          >
            Сохранить изменения
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const personalFields = ref([
  {
    id: 'name',
    label: 'Имя',
    type: 'text',
    value: 'Александр',
    placeholder: 'Введите имя',
  },
  {
    id: 'surname',
    label: 'Фамилия',
    type: 'text',
    value: 'Иванов',
    placeholder: 'Введите фамилию',
  },
  {
    id: 'email',
    label: 'Email',
    type: 'email',
    value: 'alexander@example.com',
    placeholder: 'Введите email',
  },
  {
    id: 'phone',
    label: 'Телефон',
    type: 'tel',
    value: '+7 (999) 123-45-67',
    placeholder: 'Введите телефон',
  },
])

const notificationSettings = ref([
  {
    id: 1,
    title: 'Push-уведомления',
    description: 'Получать уведомления о статусе заказа',
    enabled: true,
  },
  {
    id: 2,
    title: 'Email-рассылка',
    description: 'Получать информацию о новых акциях',
    enabled: false,
  },
  {
    id: 3,
    title: 'SMS-уведомления',
    description: 'Получать SMS о статусе доставки',
    enabled: true,
  },
])
</script>
