import axiosInstance from "./axios"

export const productsApi = {
  async getProductById(productId: string) {
    const response = await axiosInstance.get(`/api/v1/catalogue/products/${productId}`)
    return response.data
  },

  async getProductsByIds(productIds: string[]) {
    const response = await axiosInstance.get('/api/v1/catalogue/products', {
      params: { ids: productIds.join(',') }
    })
    return response.data
  }
} 