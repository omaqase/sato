<template>
  <div class="flex-1">
    <div class="bg-white rounded-2xl shadow-sm overflow-hidden">
      <div class="p-6 border-b border-gray-100">
        <div class="flex items-center justify-between">
          <h2 class="text-lg font-semibold text-gray-900">Уведомления</h2>
          <button
            @click="markAllAsRead"
            class="text-sm font-medium text-blue-600 hover:text-blue-700"
          >
            Отметить все как прочитанные
          </button>
        </div>
      </div>

      <div class="divide-y divide-gray-100">
        <div
          v-for="notification in notifications"
          :key="notification.id"
          class="p-6 hover:bg-gray-50 transition-colors"
          :class="[!notification.read && 'bg-blue-50']"
        >
          <div class="flex items-center gap-4">
            <div :class="`p-3 rounded-xl ${getNotificationStyle(notification.type).bg}`">
              <component
                :is="getNotificationStyle(notification.type).icon"
                class="w-6 h-6"
                :class="getNotificationStyle(notification.type).color"
              />
            </div>
            <div class="flex-1 min-w-0">
              <h3 class="font-medium text-gray-900">{{ notification.title }}</h3>
              <p class="mt-1 text-sm text-gray-500">{{ notification.text }}</p>
              <p class="mt-2 text-xs text-gray-400">{{ notification.date }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

interface Notification {
  id: number
  type: 'order' | 'info' | 'warning' | 'success'
  title: string
  text: string
  date: string
  read: boolean
}

const notifications = ref<Notification[]>([
  {
    id: 1,
    type: 'success',
    title: 'Заказ доставлен',
    text: 'Ваш заказ #12345 успешно доставлен',
    date: '12 марта 2024',
    read: false,
  },
  {
    id: 2,
    type: 'order',
    title: 'Заказ в пути',
    text: 'Ваш заказ #12346 передан курьеру',
    date: '11 марта 2024',
    read: false,
  },
  {
    id: 3,
    type: 'info',
    title: 'Акция',
    text: 'Скидка 20% на все товары до конца недели',
    date: '10 марта 2024',
    read: true,
  },
])

const getNotificationStyle = (type: Notification['type']) => {
  const styles = {
    order: {
      bg: 'bg-blue-50',
      color: 'text-blue-600',
      icon: 'BoxIcon',
    },
    info: {
      bg: 'bg-purple-50',
      color: 'text-purple-600',
      icon: 'InformationCircleIcon',
    },
    warning: {
      bg: 'bg-yellow-50',
      color: 'text-yellow-600',
      icon: 'ExclamationCircleIcon',
    },
    success: {
      bg: 'bg-green-50',
      color: 'text-green-600',
      icon: 'CheckCircleIcon',
    },
  }
  return styles[type]
}

const markAllAsRead = () => {
  notifications.value = notifications.value.map((n) => ({ ...n, read: true }))
}
</script>
