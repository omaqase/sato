<template>
  <div class="relative cursor-text" @click="focusInput">
    <input
      ref="hiddenInput"
      type="text"
      inputmode="numeric"
      :maxlength="otpLength"
      class="absolute opacity-0 pointer-events-none"
      v-model="code"
      @input="handleInput"
      @keydown="handleKeydown"
      @paste="handlePaste"
      @focus="handleFocus"
      @blur="handleBlur"
    />
    <div class="flex gap-3 justify-center">
      <div
        v-for="(digit, index) in digits"
        :key="index"
        :class="[
          'relative w-12 h-12 flex items-center justify-center rounded-lg bg-white shadow-sm transition-all duration-150 ease-in-out',
          index === activeIndex && isFocused
            ? 'ring-2 ring-indigo-400 border-indigo-400'
            : 'border border-gray-200',
        ]"
      >
        <span class="text-2xl font-medium text-gray-900">{{ digit }}</span>
        <div
          v-if="!digit && index === activeIndex && isFocused"
          class="absolute bottom-2 w-8 h-0.5 bg-indigo-400 rounded-full animate-blink"
        ></div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted } from 'vue'

const props = defineProps<{
  modelValue: string
  length?: number
}>()
const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void
}>()

const otpLength = props.length || 6
const code = ref(props.modelValue || '')
watch(
  () => props.modelValue,
  (val) => {
    if (val !== code.value) {
      code.value = val
    }
  },
)
watch(code, (newVal) => {
  const filtered = newVal.replace(/\D/g, '').slice(0, otpLength)
  if (filtered !== newVal) {
    code.value = filtered
    return
  }
  emit('update:modelValue', filtered)
})

const digits = computed(() => Array.from({ length: otpLength }, (_, i) => code.value[i] || ''))

const activeIndex = computed(() =>
  code.value.length < otpLength ? code.value.length : otpLength - 1,
)

const isFocused = ref(false)
const hiddenInput = ref<HTMLInputElement | null>(null)
function focusInput() {
  hiddenInput.value?.focus()
}
function handleInput(e: Event) {
  const target = e.target as HTMLInputElement
  let newVal = target.value.replace(/\D/g, '')
  if (newVal.length > otpLength) {
    newVal = newVal.slice(0, otpLength)
  }
  code.value = newVal
}
function handleKeydown(e: KeyboardEvent) {
  if (e.key === 'Backspace') return
}
function handlePaste(e: ClipboardEvent) {
  e.preventDefault()
  const pasted = e.clipboardData?.getData('text').replace(/\D/g, '') || ''
  code.value = pasted.slice(0, otpLength)
}
function handleFocus() {
  isFocused.value = true
}
function handleBlur() {
  isFocused.value = false
}
onMounted(() => {
  focusInput()
})
</script>

<style scoped>
@keyframes blink {
  0%,
  50%,
  100% {
    opacity: 1;
  }
  25%,
  75% {
    opacity: 0;
  }
}
.animate-blink {
  animation: blink 1s linear infinite;
}
</style>
