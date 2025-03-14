import { HttpClient } from '@/core/infrastructure/http/client'

export interface Category {
  category_id: string
  title: string
  slug: string
  created_at: string
  updated_at: string
}

export interface ListCategoryResponse {
  categories: Category[]
}

export const categoryApi = {
  list: () => http.get<ListCategoryResponse>('/api/v1/category'),
  getById: (id: string) => http.get<Category>(`/api/v1/category/${id}`),
  getBySlug: (slug: string) => http.get<Category>(`/api/v1/category/${slug}`)
}