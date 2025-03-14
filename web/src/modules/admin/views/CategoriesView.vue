<template>
  <div>
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-2xl font-semibold text-gray-900">Категории</h1>
      <button
        @click="showCreateModal = true"
        class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
      >
        Добавить категорию
      </button>
    </div>

    <!-- Таблица категорий -->
    <div class="bg-white shadow-sm rounded-lg overflow-hidden">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">ID</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
              Название
            </th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Slug</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
              Создано
            </th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase">
              Действия
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="category in categories" :key="category.categoryId">
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ category.categoryId }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
              {{ category.title }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ category.slug }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ new Date(category.createdAt).toLocaleDateString() }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
              <button
                @click="editCategory(category)"
                class="text-blue-600 hover:text-blue-900 mr-4"
              >
                Редактировать
              </button>
              <button
                @click="deleteCategory(category.categoryId)"
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
    <Modal v-model="showCreateModal" :title="isEditing ? 'Редактировать категорию' : 'Новая категория'">
      <form @submit.prevent="saveCategory" class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700">Название</label>
          <input
            v-model="form.title"
            type="text"
            class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
          />
        </div>
        <div>
          <label class="block text-sm font-medium text-gray-700">Slug</label>
          <input
            v-model="form.slug"
            type="text"
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
import { catalogueApi, type Category } from '@/api/catalogue'
import Modal from '@/components/Modal.vue'

const categories = ref<Category[]>([])
const showCreateModal = ref(false)
const isEditing = ref(false)
const editingId = ref<string | null>(null)

const form = ref({
  title: '',
  slug: '',
})

async function loadCategories() {
  try {
    const response = await catalogueApi.listCategories({ limit: 10, offset: 0 })
    categories.value = response.categories
  } catch (error) {
    console.error('Error loading categories:', error)
  }
}

function editCategory(category: Category) {
  isEditing.value = true
  editingId.value = category.categoryId
  form.value = {
    title: category.title,
    slug: category.slug,
  }
  showCreateModal.value = true
}

async function saveCategory() {
  try {
    if (isEditing.value && editingId.value) {
      await catalogueApi.updateCategory(editingId.value, form.value)
    } else {
      await catalogueApi.createCategory(form.value)
    }
    showCreateModal.value = false
    loadCategories()
    resetForm()
  } catch (error) {
    console.error('Error saving category:', error)
  }
}

async function deleteCategory(id: string) {
  if (confirm('Вы уверены, что хотите удалить эту категорию?')) {
    try {
      await catalogueApi.deleteCategory(id)
      loadCategories()
    } catch (error) {
      console.error('Error deleting category:', error)
    }
  }
}

function resetForm() {
  form.value = {
    title: '',
    slug: '',
  }
  isEditing.value = false
  editingId.value = null
}

onMounted(loadCategories)
</script> 