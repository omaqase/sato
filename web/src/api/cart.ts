import axiosInstance from './axios'
import { API_CONFIG } from './config'
import type { Product } from './catalogue'

export interface CartItem {
  id: string
  productId: string
  quantity: number
  createdAt: string
  updatedAt: string
  product?: Product
}

export interface GetCartResponse {
  items: CartItem[]
  total: number
}

export const cartApi = {
  async getCart(): Promise<GetCartResponse> {
    try {
      const response = await axiosInstance.get(API_CONFIG.ENDPOINTS.USER.CART)
      return response.data
    } catch (error) {
      console.error('Failed to fetch cart:', error)
      throw error
    }
  },

  async addToCart(productId: string, quantity: number = 1): Promise<void> {
    try {
      await axiosInstance.post(API_CONFIG.ENDPOINTS.USER.CART, { productId, quantity })
    } catch (error) {
      console.error('Failed to add to cart:', error)
      throw error
    }
  },

  async removeFromCart(productId: string): Promise<void> {
    try {
      await axiosInstance.delete(`${API_CONFIG.ENDPOINTS.USER.CART}/${productId}`)
    } catch (error) {
      console.error('Failed to remove from cart:', error)
      throw error
    }
  },

  async updateCartItem(productId: string, quantity: number): Promise<void> {
    try {
      await axiosInstance.patch(`${API_CONFIG.ENDPOINTS.USER.CART}/${productId}`, { quantity })
    } catch (error) {
      console.error('Failed to update cart item:', error)
      throw error
    }
  },

  async clearCart(): Promise<void> {
    try {
      await axiosInstance.delete(API_CONFIG.ENDPOINTS.USER.CART)
    } catch (error) {
      console.error('Failed to clear cart:', error)
      throw error
    }
  }
} 