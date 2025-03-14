import { ref } from 'vue'
import { createApp } from 'vue'
import ToastNotification from '@/components/ToastNotification.vue'

interface ToastOptions {
  message: string
  duration?: number
}

export function useToast() {
  const showToast = ({ message, duration = 3000 }: ToastOptions) => {
    const container = document.createElement('div')
    document.body.appendChild(container)

    const toastApp = createApp(ToastNotification, {
      message,
      duration,
      onDestroy: () => {
        document.body.removeChild(container)
      },
    })

    toastApp.mount(container)

    setTimeout(() => {
      toastApp.unmount()
      document.body.removeChild(container)
    }, duration)
  }

  const showFeatureInDevelopment = () => {
    showToast({
      message: 'Функция находится в разработке',
      duration: 3000,
    })
  }

  return {
    showToast,
    showFeatureInDevelopment,
  }
}