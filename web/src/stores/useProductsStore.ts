import { productsApi } from '@/api/products'
import { defineStore } from 'pinia'
import { ref } from 'vue'

interface Product {
  productId: string
  title: string
  price?: number
  image?: string
  categoryId: string
  sellerId: string
  stock: number
  rating: number
}

export const useProductsStore = defineStore('products', () => {
  const products = ref<Map<string, Product>>(new Map())
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function fetchProductsByIds(productIds: string[]) {
    try {
      loading.value = true
      const uniqueIds = [...new Set(productIds)]
      const fetchedProducts = await Promise.all(
        uniqueIds.map(id => productsApi.getProduct(id))
      )
      
      fetchedProducts.forEach(product => {
        if (product) {
          products.value.set(product.productId, product)
        }
      })
    } catch (err) {
      error.value = 'Failed to fetch products'
      console.error(err)
    } finally {
      loading.value = false
    }
  }

  function getProductById(productId: string): Product | undefined {
    return products.value.get(productId)
  }

  function clearProducts() {
    products.value.clear()
  }

  return {
    products,
    loading,
    error,
    fetchProductsByIds,
    getProductById,
    clearProducts
  }
}) 