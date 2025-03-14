export interface Product {
  productId: string
  title: string
  description?: string
  price: number
  oldPrice?: number
  discount?: number
  stock: number
  rating: number
  categoryId: string
  sellerId: string
  image?: string
  approved: boolean
  createdAt: string
  updatedAt?: string
}

export interface ProductReview {
  id: string
  productId: string
  userId: string
  rating: number
  comment?: string
  createdAt: string
  updatedAt: string
} 