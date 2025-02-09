import axios from 'axios'

const API_URL = 'http://localhost:8090'

export interface SendOTPRequest {
  email: string
}

export interface SendOTPResponse {
  token: string
}

export interface VerifyOTPRequest {
  token: string
  email: string
  code: string
}

export interface VerifyOTPResponse {
  access_token: string
  refresh_token: string
  is_new_user: boolean
}

export const authApi = {
  async sendOTP(data: SendOTPRequest): Promise<SendOTPResponse> {
    const response = await axios.post(`${API_URL}/api/v1/auth/send-otp`, data)
    return response.data
  },

  async verifyOTP(data: VerifyOTPRequest): Promise<VerifyOTPResponse> {
    const response = await axios.post(`${API_URL}/api/v1/auth/verify-otp`, data)
    return response.data
  },

  async refreshToken(
    refreshToken: string,
  ): Promise<{ access_token: string; refresh_token: string }> {
    const response = await axios.get(`${API_URL}/api/v1/auth/refresh-token`, {
      params: { refresh_token: refreshToken },
    })
    return response.data
  },

  async logout(refreshToken: string): Promise<void> {
    await axios.post(`${API_URL}/api/v1/auth/logout`, null, {
      params: { refresh_token: refreshToken },
    })
  },
}
