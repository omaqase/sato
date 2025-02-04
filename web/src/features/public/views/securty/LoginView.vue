<template>
  <main class="min-h-screen bg-gray-50 flex items-center justify-center p-4">
    <article class="w-full max-w-md bg-white rounded-2xl shadow-sm">
      <div class="px-6 py-8 space-y-6">
        <div v-if="!isCodeSent" class="space-y-6">
          <h2 class="text-2xl font-medium text-center">Войти или создать профиль</h2>

          <form @submit.prevent="sendCode" class="space-y-4">
            <div class="space-y-2">
              <input
                v-model="email"
                type="email"
                name="email"
                placeholder="sayat@gmail.com"
                :class="[
                  'w-full h-12 px-4 bg-gray-50 border rounded-xl focus:outline-none focus:ring-2 transition-shadow',
                  emailError
                    ? 'border-red-500 focus:ring-red-500'
                    : 'border-gray-200 focus:ring-blue-500',
                ]"
              />
              <p v-if="emailError" class="text-sm text-red-500">{{ emailError }}</p>
            </div>

            <button
              type="submit"
              :disabled="!!emailError"
              :class="[
                'w-full h-12 px-8 rounded-2xl text-white font-medium text-base transition-colors duration-200',
                emailError
                  ? 'bg-gray-400 cursor-not-allowed'
                  : 'bg-[#3B82F6] hover:bg-[#2563EB] cursor-pointer',
              ]"
            >
              Получить код
            </button>
          </form>

          <p class="text-sm text-gray-600">
            Соглашаюсь с
            <a href="#" class="underline text-gray-900 hover:text-black">
              правилами пользования торговой площадкой
            </a>
            и
            <a href="#" class="underline text-gray-900 hover:text-black"> возврата </a>
          </p>
        </div>

        <div v-else class="space-y-6">
          <button
            @click="goBack"
            class="p-2 -ml-2 text-gray-600 hover:text-gray-900 cursor-pointer"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M15 19l-7-7 7-7"
              />
            </svg>
          </button>
          <div class="space-y-3">
            <h2 class="text-2xl font-medium text-center">Введите код из письма</h2>
            <p class="text-center text-gray-600">
              Отправили на <span class="font-bold">{{ email }}</span>
            </p>
          </div>
          <div class="flex justify-between gap-2">
            <input
              v-for="i in 6"
              :key="i"
              v-model="code[i - 1]"
              type="text"
              inputmode="numeric"
              pattern="[0-9]*"
              maxlength="1"
              @keypress="onlyNumbers"
              class="w-12 h-12 text-center text-lg font-medium bg-gray-50 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              @input="(event) => onCodeInput(event, i)"
              @keydown.backspace="(event) => onBackspace(event, i)"
            />
          </div>
          <button
            @click="requestNewCode"
            class="w-full h-12 text-gray-900 bg-gray-50 rounded-xl hover:bg-gray-100 transition-colors duration-200 cursor-pointer"
          >
            Запросите код повторно
          </button>
        </div>
      </div>
    </article>
  </main>
  <VirtualAssistant />
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { z } from 'zod'
import VirtualAssistant from '@/features/public/components/VirtualAssistant/VirtualAssistant.vue'

const emailSchema = z
  .string()
  .min(1, 'Email обязателен')
  .email('Неверный формат email')
  .refine((email) => email.includes('.'), 'Email должен содержать домен')

const email = ref<string>('')
const emailError = ref<string>('')
const isCodeSent = ref<boolean>(false)
const code = ref<string[]>(['', '', '', '', '', ''])

watch(email, (newValue) => {
  try {
    emailSchema.parse(newValue)
    emailError.value = ''
  } catch (error) {
    if (error instanceof z.ZodError) {
      emailError.value = error.errors[0].message
    }
  }
})

const sendCode = async (): Promise<void> => {
  try {
    emailSchema.parse(email.value)

    isCodeSent.value = true
  } catch (error) {
    if (error instanceof z.ZodError) {
      emailError.value = error.errors[0].message
    }
  }
}

const goBack = (): void => {
  isCodeSent.value = false
  code.value = ['', '', '', '', '', '']
}

const onlyNumbers = (event: KeyboardEvent): void => {
  if (!/[0-9]/.test(event.key)) {
    event.preventDefault()
  }
}

const onCodeInput = (event: Event, index: number): void => {
  const input = event.target as HTMLInputElement

  input.value = input.value.replace(/\D/g, '')

  if (input.value && index < 6) {
    const nextInput = input.parentElement?.children[index] as HTMLInputElement | undefined
    if (nextInput) {
      nextInput.focus()
    }
  }
}

const onBackspace = (event: KeyboardEvent, index: number): void => {
  const input = event.target as HTMLInputElement
  if (!input.value && index > 1) {
    const prevInput = input.parentElement?.children[index - 2] as HTMLInputElement | undefined
    if (prevInput) {
      prevInput.focus()
    }
  }
}

const requestNewCode = (): void => {
  code.value = ['', '', '', '', '', '']
}
</script>
