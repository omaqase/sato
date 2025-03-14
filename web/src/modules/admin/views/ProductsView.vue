<template>
  <div>
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-2xl font-semibold text-gray-900">Товары</h1>
      <button
        @click="showCreateModal = true"
        class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
      >
        Добавить товар
      </button>
    </div>

    <!-- Таблица товаров -->
    <div class="bg-white shadow-sm rounded-lg overflow-hidden">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">ID</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
              Название
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
              Категория
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
              Наличие
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
              Рейтинг
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">
              Действия
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="product in products" :key="product.productId">
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ product.productId }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              {{ product.title }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ getCategoryTitle(product.categoryId) }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ product.stock }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ product.rating }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
              <button
                @click="editProduct(product)"
                class="text-blue-600 hover:text-blue-900 mr-4"
              >
                Редактировать
              </button>
              <button
                @click="deleteProduct(product.productId)"
                class="text-red-600 hover:text-red-900"
              >
                Удалить
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Модальное окно создания/редактирования -->
    <Modal v-model="showCreateModal" :title="isEditing ? 'Редактировать товар' : 'Новый товар'">
      <form @submit.prevent="saveProduct" class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700">Название</label>
          <input
            v-model="form.title"
            type="text"
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700">Категория</label>
          <select
            v-model="form.categoryId"
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
          >
            <option v-for="category in categories" :key="category.categoryId" :value="category.categoryId">
              {{ category.title }}
            </option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700">Количество</label>
          <input
            v-model.number="form.stock"
            type="number"
            min="0"
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
          />
        </div>
        <div class="flex justify-end gap-4">
          <button
            type="button"
            @click="showCreateModal = false"
            class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50"
          >
            Отмена
          </button>
          <button
            type="submit"
            class="px-4 py-2 text-sm font-medium text-white bg-blue-600 border border-transparent rounded-md hover:bg-blue-700"
          >
            {{ isEditing ? 'Сохранить' : 'Создать' }}
          </button>
        </div>
      </form>
    </Modal>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { catalogueApi, type Product, type Category } from '@/api/catalogue'
import Modal from '@/components/Modal.vue'

const products = ref<Product[]>([])
const categories = ref<Category[]>([])
const showCreateModal = ref(false)
const isEditing = ref(false)
const editingId = ref<string | null>(null)

const form = ref({
  title: '',
  categoryId: '',
  stock: 0,
})

async function loadData() {
  try {
    const [productsResponse, categoriesResponse] = await Promise.all([
      catalogueApi.listProducts({ limit: 10, offset: 0 }),
      catalogueApi.listCategories({ limit: 10, offset: 0 }),
    ])
    products.value = productsResponse.products
    categories.value = categoriesResponse.categories
  } catch (error) {
    console.error('Error loading data:', error)
  }
}

function getCategoryTitle(categoryId: string) {
  return categories.value.find((c) => c.categoryId === categoryId)?.title || 'Неизвестно'
}

function editProduct(product: Product) {
  isEditing.value = true
  editingId.value = product.productId
  form.value = {
    title: product.title,
    categoryId: product.categoryId,
    stock: product.stock,
  }
  showCreateModal.value = true
}

async function saveProduct() {
  try {
    if (isEditing.value && editingId.value) {
      await catalogueApi.updateProduct(editingId.value, {
        title: form.value.title,
        stock: form.value.stock,
      })
    } else {
      await catalogueApi.createProduct(form.value)
    }
    showCreateModal.value = false
    loadData()
    resetForm()
  } catch (error) {
    console.error('Error saving product:', error)
  }
}

async function deleteProduct(id: string) {
  if (confirm('Вы уверены, что хотите удалить этот товар?')) {
    try {
      await catalogueApi.deleteProduct(id)
      loadData()
    } catch (error) {
      console.error('Error deleting product:', error)
    }
  }
}

function resetForm() {
  form.value = {
    title: '',
    categoryId: '',
    stock: 0,
  }
  isEditing.value = false
  editingId.value = null
}

onMounted(loadData)
</script> 