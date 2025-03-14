import axiosInstance from './axios'
import { API_CONFIG } from './config'

export interface FavoriteItem {
  id: string
  productId: string
  createdAt: string
}

export interface GetFavoritesResponse {
  items: FavoriteItem[]
  total: number
}

export const favoritesApi = {
  async getFavorites(params: { limit: number; offset: number } = { limit: 10, offset: 0 }): Promise<GetFavoritesResponse> {
    try {
      const response = await axiosInstance.get(API_CONFIG.ENDPOINTS.USER.FAVORITES, { params })
      return response.data
    } catch (error) {
      console.error('Failed to fetch favorites:', error)
      throw error
    }
  },

  async addToFavorites(productId: string): Promise<void> {
    try {
      await axiosInstance.post(API_CONFIG.ENDPOINTS.USER.FAVORITES, { productId })
    } catch (error) {
      console.error('Failed to add to favorites:', error)
      throw error
    }
  },

  async removeFromFavorites(productId: string): Promise<void> {
    try {
      await axiosInstance.delete(API_CONFIG.ENDPOINTS.USER.FAVORITES, {
        params: { productId }
      })
    } catch (error) {
      console.error('Failed to remove from favorites:', error)
      throw error
    }
  }
}