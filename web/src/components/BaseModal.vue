<template>
  <Transition name="modal">
    <div v-if="isOpen" class="fixed inset-0 z-50 overflow-y-auto" @click="closeModal">
      <div class="flex min-h-screen items-center justify-center p-4">
        <div class="fixed inset-0 bg-black opacity-30"></div>

        <div class="relative bg-white rounded-2xl p-8 max-w-lg w-full shadow-xl" @click.stop>
          <button
            class="absolute top-4 right-4 text-gray-400 hover:text-gray-600 transition-colors"
            @click="closeModal"
          >
            <svg class="w-6 h-6" viewBox="0 0 24 24" fill="none">
              <path
                d="M6 18L18 6M6 6l12 12"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
              />
            </svg>
          </button>

          <div class="text-center">
            <div
              class="w-16 h-16 mx-auto mb-4 rounded-full bg-blue-100 flex items-center justify-center"
            >
              <svg class="w-8 h-8 text-blue-600" viewBox="0 0 24 24" fill="none">
                <path
                  d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linecap="round"
                />
              </svg>
            </div>

            <h3 class="text-xl font-semibold text-gray-900 mb-2">
              {{ title }}
            </h3>
            <p class="text-gray-600">
              {{ message }}
            </p>

            <button
              class="mt-6 px-6 py-3 bg-blue-600 text-white rounded-xl hover:bg-blue-700 transition-colors"
              @click="closeModal"
            >
              Понятно
            </button>
          </div>
        </div>
      </div>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const isOpen = ref(false)
const title = ref('')
const message = ref('')

const openModal = (modalTitle: string, modalMessage: string) => {
  title.value = modalTitle
  message.value = modalMessage
  isOpen.value = true
}

const closeModal = () => {
  isOpen.value = false
}

defineExpose({
  openModal,
  closeModal,
})
</script>

<style scoped>
.modal-enter-active,
.modal-leave-active {
  transition: opacity 0.3s ease;
}

.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}
</style>
