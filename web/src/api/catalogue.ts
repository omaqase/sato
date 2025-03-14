import axiosInstance from "./axios"
import { API_CONFIG } from './config'

export interface Category {
  categoryId: string
  title: string
  slug: string
  createdAt: string
  updatedAt?: string
}

export interface Product {
  productId: string
  categoryId: string
  sellerId: string
  title: string
  stock: number
  rating: number
  approved: boolean
  createdAt: string
  updatedAt?: string
  price?: number
  oldPrice?: number
  discount?: number
  image?: string
}

export const catalogueApi = {
  // Categories
  async listCategories(params?: { limit?: number; offset?: number }) {
    try {
      const queryParams = new URLSearchParams()
      
      if (params?.limit) {
        queryParams.append('limit', params.limit.toString())
      }
      
      if (params?.offset) {
        queryParams.append('offset', params.offset.toString())
      }
      
      const url = `${API_CONFIG.ENDPOINTS.CATALOGUE.CATEGORIES}${queryParams.toString() ? `?${queryParams.toString()}` : ''}`
      const response = await axiosInstance.get(url)
      return response.data
    } catch (error) {
      console.error('Failed to fetch categories:', error)
      throw error
    }
  },

  async createCategory(data: { title: string; slug: string }) {
    try {
      const response = await axiosInstance.post(API_CONFIG.ENDPOINTS.ADMIN.CATEGORIES, data)
      return response.data
    } catch (error) {
      console.error('Failed to create category:', error)
      throw error
    }
  },

  async updateCategory(categoryId: string, data: { title: string; slug: string }) {
    try {
      const response = await axiosInstance.patch(`${API_CONFIG.ENDPOINTS.ADMIN.CATEGORIES}/${categoryId}`, data)
      return response.data
    } catch (error) {
      console.error('Failed to update category:', error)
      throw error
    }
  },

  async deleteCategory(categoryId: string) {
    try {
      const response = await axiosInstance.delete(`${API_CONFIG.ENDPOINTS.ADMIN.CATEGORIES}/${categoryId}`)
      return response.data
    } catch (error) {
      console.error('Failed to delete category:', error)
      throw error
    }
  },

  // Products
  async getProduct(productId: string): Promise<Product> {
    try {
      const response = await axiosInstance.get(`${API_CONFIG.ENDPOINTS.CATALOGUE.PRODUCTS}/${productId}`)
      return response.data
    } catch (error) {
      console.error('Failed to fetch product:', error)
      throw error
    }
  },

  async listProducts(params?: { limit?: number; offset?: number }) {
    try {
      const response = await axiosInstance.get(API_CONFIG.ENDPOINTS.CATALOGUE.PRODUCTS, { params })
      return response.data
    } catch (error) {
      console.error('Failed to fetch products:', error)
      throw error
    }
  },

  async createProduct(data: { categoryId: string; title: string; stock: number }) {
    try {
      const response = await axiosInstance.post(API_CONFIG.ENDPOINTS.SELLER.PRODUCTS, data)
      return response.data
    } catch (error) {
      console.error('Failed to create product:', error)
      throw error
    }
  },

  async updateProduct(productId: string, data: { title: string; stock: number }) {
    try {
      const response = await axiosInstance.patch(`${API_CONFIG.ENDPOINTS.SELLER.PRODUCTS}/${productId}`, data)
      return response.data
    } catch (error) {
      console.error('Failed to update product:', error)
      throw error
    }
  },

  async deleteProduct(productId: string) {
    try {
      const response = await axiosInstance.delete(`${API_CONFIG.ENDPOINTS.CATALOGUE.PRODUCTS}/${productId}`)
      return response.data
    } catch (error) {
      console.error('Failed to delete product:', error)
      throw error
    }
  },
} 