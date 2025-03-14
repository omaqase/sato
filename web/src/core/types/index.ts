// Core types and interfaces

// Common response type for API calls
export interface ApiResponse<T> {
  data: T
  status: number
  message?: string
}

// Common error type
export interface ApiError {
  code: string
  message: string
  details?: Record<string, any>
}

// Pagination types
export interface PaginationParams {
  page: number
  limit: number
}

export interface PaginatedResponse<T> {
  items: T[]
  total: number
  page: number
  totalPages: number
}

// Common entity interfaces
export interface BaseEntity {
  id: string | number
  createdAt: string
  updatedAt: string
}

// User related types
export interface User extends BaseEntity {
  email: string
  name: string
  role: UserRole
}

export enum UserRole {
  ADMIN = 'admin',
  USER = 'user'
}

// Common status types
export enum Status {
  ACTIVE = 'active',
  INACTIVE = 'inactive',
  PENDING = 'pending',
  DELETED = 'deleted'
}