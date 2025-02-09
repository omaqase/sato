<template>
  <div class="min-h-screen bg-gradient-to-b from-gray-50 to-white pb-12">
    <div class="max-w-7xl mx-auto px-4 py-8">
      <!-- Breadcrumbs -->
      <div class="mb-6 flex items-center gap-2 text-sm">
        <router-link to="/" class="text-gray-500 hover:text-blue-600">Главная</router-link>
        <span class="text-gray-400">/</span>
        <router-link to="/category/electronics" class="text-gray-500 hover:text-blue-600">
          Электроника
        </router-link>
        <span class="text-gray-400">/</span>
        <span class="text-gray-900">{{ product.name }}</span>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-2 gap-12">
        <!-- Product Images -->
        <div class="space-y-4">
          <div class="aspect-square bg-white rounded-2xl p-8 border-2 border-gray-100">
            <img :src="product.image" :alt="product.name" class="w-full h-full object-contain" />
          </div>
          <div class="grid grid-cols-4 gap-4">
            <button
              v-for="(image, index) in product.images"
              :key="index"
              class="aspect-square bg-white rounded-xl p-4 border-2 border-gray-100 hover:border-blue-500"
            >
              <img :src="image" :alt="product.name" class="w-full h-full object-contain" />
            </button>
          </div>
        </div>

        <!-- Product Info -->
        <div>
          <h1 class="text-3xl font-bold text-gray-900 mb-4">{{ product.name }}</h1>

          <div class="flex items-center gap-4 mb-6">
            <div class="flex items-center gap-1">
              <div class="flex text-yellow-400">
                <svg
                  v-for="i in 5"
                  :key="i"
                  class="w-5 h-5"
                  :class="i <= product.rating ? 'fill-current' : 'stroke-current fill-none'"
                >
                  <path
                    d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"
                  />
                </svg>
              </div>
              <span class="text-sm text-gray-600">{{ product.rating }}</span>
            </div>
            <span class="text-gray-400">•</span>
            <button class="text-sm text-gray-600 hover:text-blue-600">
              {{ product.reviewsCount }} отзывов
            </button>
            <span class="text-gray-400">•</span>
            <button class="text-sm text-gray-600 hover:text-blue-600">
              {{ product.ordersCount }} заказов
            </button>
          </div>

          <div class="bg-white rounded-2xl p-6 border-2 border-gray-100 mb-6">
            <div class="flex items-baseline gap-4 mb-4">
              <span class="text-3xl font-bold">{{ formatPrice(product.price) }} ₽</span>
              <span class="text-xl text-gray-400 line-through">
                {{ formatPrice(product.oldPrice) }} ₽
              </span>
              <span class="px-2 py-1 bg-red-100 text-red-600 text-sm font-medium rounded">
                -{{ product.discount }}%
              </span>
            </div>

            <div class="space-y-4">
              <button
                class="w-full py-4 bg-blue-600 hover:bg-blue-700 text-white font-medium rounded-xl transition-colors"
              >
                Добавить в корзину
              </button>
              <button
                class="w-full py-4 bg-gray-100 hover:bg-gray-200 text-gray-900 font-medium rounded-xl transition-colors"
              >
                В избранное
              </button>
            </div>
          </div>

          <!-- Product Details -->
          <div class="space-y-6">
            <div v-for="(section, index) in product.details" :key="index">
              <h3 class="font-medium text-gray-900 mb-2">{{ section.title }}</h3>
              <div class="space-y-2">
                <div
                  v-for="(spec, key) in section.specs"
                  :key="key"
                  class="flex items-center gap-4 text-sm"
                >
                  <span class="text-gray-500">{{ key }}</span>
                  <div class="flex-1 border-b border-dotted border-gray-300"></div>
                  <span class="text-gray-900">{{ spec }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Reviews Section -->
      <div class="mt-16">
        <div class="flex items-center justify-between mb-8">
          <h2 class="text-2xl font-bold text-gray-900">
            Отзывы покупателей ({{ product.reviewsCount }})
          </h2>
          <button
            @click="showReviewModal = true"
            class="px-6 py-3 bg-blue-600 hover:bg-blue-700 text-white font-medium rounded-xl transition-colors"
          >
            Написать отзыв
          </button>
        </div>

        <!-- Reviews List -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-8">
          <div
            v-for="review in displayedReviews"
            :key="review.id"
            class="bg-white rounded-xl p-4 border border-gray-100 hover:border-gray-200 transition-all"
          >
            <div class="flex items-center justify-between mb-3">
              <span class="font-medium text-gray-900">{{ review.author }}</span>
              <span class="text-sm text-gray-400">{{ review.date }}</span>
            </div>

            <div class="flex text-yellow-400 mb-2">
              <svg
                v-for="i in 5"
                :key="i"
                class="w-4 h-4"
                :class="i <= review.rating ? 'fill-current' : 'stroke-current fill-none'"
              >
                <path
                  d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"
                />
              </svg>
            </div>

            <p class="text-sm text-gray-600 line-clamp-3 mb-2">{{ review.text }}</p>

            <div class="flex justify-between items-center">
              <button class="text-sm text-blue-600 hover:text-blue-700">Читать полностью</button>
              <div class="flex items-center gap-1">
                <button class="text-gray-400 hover:text-blue-600">
                  <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none">
                    <path
                      d="M14 9V5a3 3 0 0 0-3-3l-4 9v11h11.28a2 2 0 0 0 2-1.7l1.38-9a2 2 0 0 0-2-2.3zM7 22H4a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h3"
                      stroke="currentColor"
                      stroke-width="2"
                    />
                  </svg>
                </button>
                <span class="text-sm text-gray-500">{{ review.likes }}</span>
              </div>
            </div>
          </div>
        </div>

        <!-- Show More Button -->
        <div class="text-center">
          <button
            @click="showAllReviews"
            class="px-8 py-3 bg-gray-100 hover:bg-gray-200 text-gray-900 font-medium rounded-xl transition-colors"
          >
            Смотреть все отзывы
          </button>
        </div>

        <!-- Review Modal -->
        <Transition name="modal">
          <div v-if="showReviewModal" class="fixed inset-0 z-50">
            <div
              class="absolute inset-0 bg-black/60 backdrop-blur-sm"
              @click="showReviewModal = false"
            ></div>
            <div class="absolute inset-0 flex items-center justify-center p-4">
              <div
                class="bg-white rounded-2xl w-full max-w-lg transform transition-all"
                :class="showReviewModal ? 'scale-100 opacity-100' : 'scale-95 opacity-0'"
              >
                <div class="p-6 border-b border-gray-100">
                  <div class="flex items-center justify-between">
                    <h3 class="text-xl font-bold text-gray-900">Написать отзыв</h3>
                    <button
                      @click="showReviewModal = false"
                      class="p-2 text-gray-400 hover:text-gray-600 rounded-lg hover:bg-gray-100"
                    >
                      <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none">
                        <path
                          d="M6 18L18 6M6 6l12 12"
                          stroke="currentColor"
                          stroke-width="2"
                          stroke-linecap="round"
                        />
                      </svg>
                    </button>
                  </div>
                </div>

                <div class="p-6 space-y-6">
                  <!-- Rating -->
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-2"
                      >Оценка товара</label
                    >
                    <div class="flex items-center gap-2">
                      <div class="flex gap-1">
                        <button
                          v-for="i in 5"
                          :key="i"
                          @mouseenter="hoverRating = i"
                          @mouseleave="hoverRating = 0"
                          @click="newReview.rating = i"
                          class="p-1 transition-transform hover:scale-110"
                        >
                          <svg
                            class="w-8 h-8"
                            :class="[
                              hoverRating >= i || newReview.rating >= i
                                ? 'text-yellow-400 fill-current'
                                : 'text-gray-300 fill-current',
                            ]"
                            viewBox="0 0 24 24"
                          >
                            <path
                              d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"
                            />
                          </svg>
                        </button>
                      </div>
                      <span class="text-sm text-gray-500">
                        {{ ratingLabels[newReview.rating - 1] || 'Выберите оценку' }}
                      </span>
                    </div>
                    <p v-if="showErrors && !newReview.rating" class="mt-1 text-sm text-red-500">
                      Пожалуйста, поставьте оценку
                    </p>
                  </div>

                  <!-- Review Text -->
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-2">
                      Текст отзыва <span class="text-red-500">*</span>
                    </label>
                    <textarea
                      v-model="newReview.text"
                      rows="4"
                      :class="[
                        'w-full px-4 py-3 rounded-xl border-2 transition-colors',
                        showErrors && !newReview.text
                          ? 'border-red-300 focus:border-red-500'
                          : 'border-gray-100 focus:border-blue-500',
                      ]"
                      placeholder="Поделитесь своими впечатлениями о товаре"
                    ></textarea>
                    <p v-if="showErrors && !newReview.text" class="mt-1 text-sm text-red-500">
                      Пожалуйста, напишите отзыв
                    </p>
                  </div>

                  <!-- Pros & Cons -->
                  <div class="grid grid-cols-2 gap-4">
                    <div>
                      <label class="block text-sm font-medium text-gray-700 mb-2"
                        >Достоинства</label
                      >
                      <input
                        v-model="newReview.pros"
                        type="text"
                        class="w-full px-4 py-3 rounded-xl border-2 border-gray-100 focus:border-blue-500"
                        placeholder="Что понравилось?"
                      />
                    </div>
                    <div>
                      <label class="block text-sm font-medium text-gray-700 mb-2">Недостатки</label>
                      <input
                        v-model="newReview.cons"
                        type="text"
                        class="w-full px-4 py-3 rounded-xl border-2 border-gray-100 focus:border-blue-500"
                        placeholder="Что не понравилось?"
                      />
                    </div>
                  </div>
                </div>

                <div class="p-6 border-t border-gray-100 flex justify-end gap-3">
                  <button
                    @click="showReviewModal = false"
                    class="px-4 py-2 text-gray-700 font-medium hover:bg-gray-100 rounded-lg transition-colors"
                  >
                    Отмена
                  </button>
                  <button
                    @click="addReview"
                    class="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white font-medium rounded-lg transition-colors"
                  >
                    Опубликовать
                  </button>
                </div>
              </div>
            </div>
          </div>
        </Transition>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'

const product = ref({
  id: 1,
  name: 'Apple iPhone 13 Pro 256GB',
  price: 89990,
  oldPrice: 99990,
  discount: 10,
  rating: 4.9,
  reviewsCount: 256,
  ordersCount: 1205,
  image: '/products/iphone.jpg',
  images: [
    '/products/iphone-1.jpg',
    '/products/iphone-2.jpg',
    '/products/iphone-3.jpg',
    '/products/iphone-4.jpg',
  ],
  details: [
    {
      title: 'Характеристики',
      specs: {
        Экран: '6.1", OLED',
        Процессор: 'Apple A15 Bionic',
        Память: '256 ГБ',
        Камера: '12 MP + 12 MP + 12 MP',
        Аккумулятор: '3095 мАч',
      },
    },
    {
      title: 'Общие характеристики',
      specs: {
        Гарантия: '1 год',
        'Страна производства': 'Китай',
        'Год выпуска': '2021',
      },
    },
  ],
})

const formatPrice = (price: number) => {
  return new Intl.NumberFormat('ru-RU').format(price)
}

interface Review {
  id: number
  author: string
  date: string
  rating: number
  text: string
  pros?: string
  cons?: string
  likes: number
  photos?: string[]
}

const reviews = ref<Review[]>([
  {
    id: 1,
    author: 'Александр В.',
    date: '15 марта 2024',
    rating: 5,
    text: 'Отличный телефон, пользуюсь уже месяц. Камера супер, батарея держит долго.',
    pros: 'Камера, батарея, производительность, дисплей',
    cons: 'Высокая цена',
    likes: 24,
    photos: ['/reviews/photo1.jpg', '/reviews/photo2.jpg'],
  },
  {
    id: 2,
    author: 'Мария К.',
    date: '10 марта 2024',
    rating: 4,
    text: 'В целом доволен покупкой, но есть небольшие недочеты.',
    pros: 'Качество сборки, экран, звук',
    cons: 'Греется при длительных играх',
    likes: 15,
  },
])

const router = useRouter()
const showReviewModal = ref(false)
const displayedReviews = computed(() => reviews.value.slice(0, 4))

const newReview = ref({
  rating: 5,
  text: '',
  pros: '',
  cons: '',
})

const hoverRating = ref(0)
const showErrors = ref(false)
const ratingLabels = ['Ужасно', 'Плохо', 'Нормально', 'Хорошо', 'Отлично']

function showAllReviews() {
  router.push({ name: 'reviews', params: { id: product.value.id } })
}

function addReview() {
  showErrors.value = true

  if (!newReview.value.rating || !newReview.value.text) {
    return
  }

  const review: Review = {
    id: reviews.value.length + 1,
    author: 'Пользователь',
    date: new Date().toLocaleDateString('ru-RU'),
    rating: newReview.value.rating,
    text: newReview.value.text,
    pros: newReview.value.pros,
    cons: newReview.value.cons,
    likes: 0,
  }

  reviews.value.unshift(review)
  showReviewModal.value = false
  newReview.value = { rating: 5, text: '', pros: '', cons: '' }
  showErrors.value = false
}
</script>

<style scoped>
.modal-enter-active,
.modal-leave-active {
  transition: all 0.3s ease;
}

.modal-enter-from,
.modal-leave-to {
  opacity: 0;
  transform: scale(0.95);
}
</style>
