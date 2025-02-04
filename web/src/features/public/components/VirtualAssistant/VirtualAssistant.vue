<template>
  <div>
    <!-- Кнопка для открытия чата -->
    <button
      @click="toggleChat"
      class="fixed bottom-6 right-6 w-14 h-14 bg-[#3B82F6] rounded-full flex items-center justify-center shadow-lg hover:bg-[#2563EB] transition-all duration-300 z-40 cursor-pointer"
      :class="{ 'scale-0': isOpen }"
    >
      <svg class="w-7 h-7 text-white" viewBox="0 0 24 24" fill="currentColor">
        <path
          d="M12 2C6.48 2 2 6.48 2 12c0 5.52 4.48 10 10 10s10-4.48 10-10c0-5.52-4.48-10-10-10zM8 14c-1.1 0-2-.9-2-2s.9-2 2-2 2 .9 2 2-.9 2-2 2zm4 0c-1.1 0-2-.9-2-2s.9-2 2-2 2 .9 2 2-.9 2-2 2zm4 0c-1.1 0-2-.9-2-2s.9-2 2-2 2 .9 2 2-.9 2-2 2z"
        />
      </svg>
    </button>

    <!-- Чат -->
    <Transition
      enter-active-class="transition-all duration-300 ease-out"
      enter-from-class="translate-y-full opacity-0"
      enter-to-class="translate-y-0 opacity-100"
      leave-active-class="transition-all duration-300 ease-in"
      leave-from-class="translate-y-0 opacity-100"
      leave-to-class="translate-y-full opacity-0"
    >
      <div
        v-if="isOpen"
        class="fixed bottom-0 right-0 w-full md:w-[400px] h-[100dvh] md:h-10/12 bg-white shadow-xl flex flex-col z-50 md:bottom-6 md:right-6 md:rounded-xl"
      >
        <header class="flex items-center px-4 py-3 bg-[#3B82F6] md:rounded-t-xl">
          <div class="flex items-center space-x-3">
            <div class="w-8 h-8 bg-white/10 rounded-full flex items-center justify-center">
              <svg class="w-5 h-5 text-white" viewBox="0 0 24 24" fill="none" stroke="currentColor">
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M3 18v-6a9 9 0 0118 0v6"
                />
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M21 19a2 2 0 01-2 2h-1a2 2 0 01-2-2v-3a2 2 0 012-2h3zM3 19a2 2 0 002 2h1a2 2 0 002-2v-3a2 2 0 00-2-2H3z"
                />
              </svg>
            </div>
            <span class="text-white text-lg">Поддержка Sato ✌️</span>
          </div>
          <button
            ref="closeButton"
            @click="closeChat"
            class="ml-auto p-2 hover:bg-white/10 rounded-full transition-colors cursor-pointer"
          >
            <svg class="w-5 h-5 text-white" viewBox="0 0 24 24" fill="none" stroke="currentColor">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M6 18L18 6M6 6l12 12"
              />
            </svg>
          </button>
        </header>

        <!-- История сообщений -->
        <div class="flex-1 overflow-y-auto p-4 space-y-4 scroll-smooth">
          <div
            v-for="(msg, index) in messages"
            :key="index"
            :class="[
              'space-y-2 pl-4 py-4 pr-4 rounded-xl',
              msg.sender === 'Лина' ? 'bg-gray-100' : 'bg-blue-100',
            ]"
          >
            <div class="font-bold">{{ msg.sender }}</div>
            <div class="text-gray-700 whitespace-pre-wrap">{{ msg.text }}</div>
            <div class="text-right text-xs text-gray-500">{{ formatTime(msg.time) }}</div>
          </div>
        </div>

        <!-- Поле ввода -->
        <div class="p-4 border-t border-gray-100">
          <div class="flex flex-wrap gap-2 mb-4">
            <button
              v-for="action in quickActions"
              :key="action"
              @click="sendPredefinedMessage(action)"
              class="px-4 py-2 bg-[#F5F0FF] text-[#3B82F6] rounded-full text-sm hover:bg-[#EDE5FF] transition-colors cursor-pointer"
            >
              {{ action }}
            </button>
          </div>
          <div class="flex items-center gap-3 bg-[#F8F8F8] rounded-full px-4 py-2">
            <input
              v-model="message"
              type="text"
              placeholder="Ваше сообщение..."
              class="flex-1 bg-transparent border-none focus:outline-none text-gray-700 placeholder-gray-400"
              @keyup.enter="sendMessage"
            />
            <button
              class="text-[#6F2DFF] disabled:text-gray-300 active:cursor-pointer"
              :disabled="!message.trim()"
              @click="sendMessage"
            >
              <svg class="w-5 h-5 rotate-90" viewBox="0 0 24 24" fill="none" stroke="currentColor">
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"
                />
              </svg>
            </button>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, nextTick } from 'vue'

const isOpen = ref(false)
const message = ref('')
const messages = ref<Array<{ sender: string; text: string; time: Date }>>([])
const quickActions = ['Узнать о пункте выдачи', 'Не могу зайти в профиль']

// Методы
const toggleChat = () => {
  isOpen.value = !isOpen.value
  if (isOpen.value && messages.value.length === 0) {
    messages.value.push({
      sender: 'Лина',
      text: 'Здравствуйте! Я — Лина, ваша виртуальная помощница.',
      time: new Date(),
    })
  }
}

const closeChat = () => {
  isOpen.value = false
  message.value = ''
}

const sendMessage = async () => {
  if (!message.value.trim()) return

  // Добавляем сообщение пользователя
  const userMessage = message.value.trim()
  messages.value.push({
    sender: 'Вы',
    text: userMessage,
    time: new Date(),
  })

  // Очищаем поле ввода
  message.value = ''

  // Отправляем запрос на сервер
  await fetchResponse(userMessage)
}

const sendPredefinedMessage = async (text: string) => {
  messages.value.push({
    sender: 'Вы',
    text,
    time: new Date(),
  })

  // Отправляем запрос на сервер
  await fetchResponse(text)
}

const fetchResponse = async (prompt: string) => {
  try {
    // Создаем временное сообщение для ответа ассистента
    const assistantMessage = {
      sender: 'Лина',
      text: '',
      time: new Date(),
    }
    messages.value.push(assistantMessage)

    // Отправляем запрос на сервер
    const response = await fetch('http://localhost:3000/ask', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ prompt }),
    })

    if (!response.ok) {
      throw new Error('Ошибка при отправке запроса')
    }

    // Обрабатываем потоковый ответ
    const reader = response.body?.getReader()
    const decoder = new TextDecoder()
    while (true) {
      const { done, value } = await reader!.read()
      if (done) break
      const chunk = decoder.decode(value, { stream: true })
      assistantMessage.text += chunk

      // Принудительно обновляем DOM
      await nextTick()
    }
  } catch (error) {
    console.error('Ошибка:', error)
    messages.value.push({
      sender: 'Лина',
      text: 'Извините, произошла ошибка. Попробуйте позже.',
      time: new Date(),
    })
  }
}

const formatTime = (date: Date): string => {
  return date.toLocaleTimeString('ru-RU', {
    hour: '2-digit',
    minute: '2-digit',
  })
}

// Обработка клавиш
const handleEscKey = (event: KeyboardEvent) => {
  if (event.key === 'Escape' && isOpen.value) {
    closeChat()
  }
}

onMounted(() => {
  document.addEventListener('keydown', handleEscKey)
})

onUnmounted(() => {
  document.removeEventListener('keydown', handleEscKey)
})
</script>
