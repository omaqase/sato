import axios from 'axios'
import { authApi } from './auth'

const axiosInstance = axios.create({
  baseURL: 'http://localhost:8090',
})

axiosInstance.interceptors.request.use((config) => {
  const token = localStorage.getItem('access_token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

axiosInstance.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config

    if (error.response?.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true

      try {
        const refreshToken = localStorage.getItem('refresh_token')
        if (!refreshToken) {
          throw new Error('No refresh token')
        }

        const tokens = await authApi.refreshToken(refreshToken)
        localStorage.setItem('access_token', tokens.access_token)
        localStorage.setItem('refresh_token', tokens.refresh_token)

        originalRequest.headers.Authorization = `Bearer ${tokens.access_token}`
        return axios(originalRequest)
      } catch (error) {
        localStorage.removeItem('access_token')
        localStorage.removeItem('refresh_token')
        window.location.href = '/security/login'
      }
    }

    return Promise.reject(error)
  },
)

export default axiosInstance
