<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 p-6">
    <div class="bg-white rounded-2xl shadow-lg p-10 max-w-md w-full">
      <h2 class="text-3xl font-bold text-center text-gray-900 mb-8">Войти в систему</h2>
      <div v-if="errorMessage" class="mb-4 text-sm text-red-500 text-center">
        {{ errorMessage }}
      </div>
      <div v-if="successMessage" class="mb-4 text-sm text-green-500 text-center">
        {{ successMessage }}
      </div>
      <transition name="fade" mode="out-in">
        <template v-if="currentStep === Step.Email">
          <div key="step-email">
            <label for="email" class="block text-gray-700 mb-2">Email</label>
            <input
              id="email"
              type="email"
              v-model="formData.email"
              placeholder="Введите ваш email"
              class="w-full px-4 py-3 rounded-xl border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 transition"
            />
            <button
              @click="sendCode"
              :disabled="isLoading"
              class="w-full mt-6 bg-blue-600 hover:bg-blue-700 text-white font-semibold py-3 rounded-xl transition disabled:opacity-50"
            >
              Продолжить
            </button>
          </div>
        </template>
        <template v-else-if="currentStep === Step.Code">
          <div key="step-code">
            <label class="block text-gray-700 mb-2">6-значный код</label>
            <OTPInput v-model="formData.code" />
            <button
              @click="verifyCode"
              :disabled="isLoading"
              class="w-full mt-6 bg-blue-600 hover:bg-blue-700 text-white font-semibold py-3 rounded-xl transition disabled:opacity-50"
            >
              Войти
            </button>
          </div>
        </template>
      </transition>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { z } from 'zod'
import OTPInput from './OTPInput.vue'

// Определение шагов формы
enum Step {
  Email = 'email',
  Code = 'code',
}

const currentStep = ref<Step>(Step.Email)

// Объект для хранения данных формы
const formData = reactive({
  email: '',
  code: '',
})

// Состояния для сообщений, загрузки и обработки ошибок
const errorMessage = ref<string>('')
const successMessage = ref<string>('')
const isLoading = ref<boolean>(false)

// Схемы валидации с помощью Zod
const emailSchema = z.object({
  email: z
    .string()
    .min(1, { message: 'Введите email' })
    .email({ message: 'Неверный формат email' }),
})

const codeSchema = z.object({
  code: z.string().regex(/^\d{6}$/, { message: 'Код должен состоять из 6 цифр' }),
})

// Функция для отправки кода на почту
async function sendCode() {
  errorMessage.value = ''
  successMessage.value = ''

  const validation = emailSchema.safeParse({ email: formData.email })
  if (!validation.success) {
    errorMessage.value = validation.error.errors.map((err) => err.message).join(', ')
    return
  }

  isLoading.value = true
  try {
    await new Promise((resolve) => setTimeout(resolve, 1500))
    successMessage.value = 'Код отправлен. Проверьте почту.'
    currentStep.value = Step.Code
  } catch (error: unknown) {
    errorMessage.value = error instanceof Error ? error.message : 'Ошибка сервера'
  } finally {
    isLoading.value = false
  }
}

// Функция для проверки введённого кода
async function verifyCode() {
  errorMessage.value = ''
  successMessage.value = ''

  const validation = codeSchema.safeParse({ code: formData.code })
  if (!validation.success) {
    errorMessage.value = validation.error.errors.map((err) => err.message).join(', ')
    return
  }

  isLoading.value = true
  try {
    await new Promise((resolve) => setTimeout(resolve, 1500))
    if (formData.code !== '123456') {
      throw new Error('Неверный код')
    }
    successMessage.value = 'Добро пожаловать!'
  } catch (error: unknown) {
    errorMessage.value = error instanceof Error ? error.message : 'Ошибка сервера'
  } finally {
    isLoading.value = false
  }
}
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
