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
import { useRouter } from 'vue-router'
import { z } from 'zod'
import { authApi } from '@/api/auth'
import OTPInput from './OTPInput.vue'
import { useAuthStore } from '@/stores/useAuthStore'

// Определение шагов формы
enum Step {
  Email = 'email',
  Code = 'code',
}

const router = useRouter()
const currentStep = ref<Step>(Step.Email)
const token = ref<string>('')

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
    errorMessage.value = validation.error.errors[0].message
    return
  }

  isLoading.value = true
  try {
    const response = await authApi.sendOTP({ email: formData.email })
    token.value = response.token
    successMessage.value = 'Код отправлен. Проверьте почту.'
    currentStep.value = Step.Code
  } catch (error: any) {
    errorMessage.value = error.response?.data?.message || 'Ошибка при отправке кода'
  } finally {
    isLoading.value = false
  }
}

// Функция для проверки введённого кода
async function verifyCode() {
  errorMessage.value = ''

  const validation = codeSchema.safeParse({ code: formData.code })
  if (!validation.success) {
    errorMessage.value = validation.error.errors[0].message
    return
  }

  isLoading.value = true
  try {
    const response = await authApi.verifyOTP({
      token: token.value,
      email: formData.email,
      code: formData.code,
    })

    // Сохраняем токены
    localStorage.setItem('access_token', response.access_token)
    localStorage.setItem('refresh_token', response.refresh_token)

    // Сохраняем данные пользователя
    const authStore = useAuthStore()
    authStore.setUser({
      email: formData.email,
    })

    // Редиректим на нужную страницу
    if (response.is_new_user) {
      router.push('/account/settings')
    } else {
      router.push('/account/overview')
    }
  } catch (error: any) {
    errorMessage.value = error.response?.data?.message || 'Неверный код'
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
