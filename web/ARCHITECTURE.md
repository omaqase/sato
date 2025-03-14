# Project Architecture

## Overview
This document outlines the new architecture for the Sato project, designed to improve scalability and maintainability.

## Core Principles
1. **Module-First Architecture**: Each feature is a self-contained module
2. **Clean Architecture**: Clear separation of concerns
3. **Domain-Driven Design**: Business logic organized around domain concepts

## Directory Structure
```
src/
├── core/                 # Core application logic
│   ├── types/           # TypeScript types and interfaces
│   ├── constants/       # Application constants
│   └── utils/           # Utility functions
├── modules/             # Feature modules
│   ├── auth/            # Authentication module
│   ├── admin/           # Admin module
│   ├── products/        # Products module
│   ├── orders/          # Orders module
│   └── profile/         # User profile module
├── shared/              # Shared components and utilities
│   ├── components/      # Reusable UI components
│   ├── composables/     # Vue composables
│   ├── layouts/         # Layout components
│   └── assets/          # Static assets
├── infrastructure/      # Infrastructure layer
│   ├── api/            # API clients
│   ├── storage/        # Storage services
│   └── http/           # HTTP utilities
└── store/               # State management
    ├── modules/         # Store modules
    └── plugins/         # Store plugins
```

## Module Structure
Each feature module follows this structure:
```
module/
├── components/     # Module-specific components
├── composables/    # Module-specific composables
├── types/          # Module-specific types
├── services/       # Business logic services
├── store/          # State management
└── routes.ts       # Module routes
```

## State Management
- Use Pinia for state management
- Each module has its own store
- Shared state in root store

## Routing
- Module-based routing
- Lazy-loaded routes
- Route middleware for auth/permissions

## API Layer
- Axios for HTTP requests
- API clients per module
- Request/response interceptors
- Error handling

## Component Guidelines
1. Use composition API
2. Props validation
3. Emit typed events
4. Keep components focused
5. Use composables for logic reuse

## Testing Strategy
1. Unit tests for business logic
2. Component tests with Vue Test Utils
3. E2E tests with Cypress

## Code Style
- Follow Vue Style Guide
- Use TypeScript
- ESLint + Prettier

## Performance Considerations
1. Lazy loading routes
2. Code splitting
3. Asset optimization
4. Cache strategies

## Security
1. Input validation
2. XSS prevention
3. CSRF protection
4. Authentication/Authorization

## Deployment
1. Build optimization
2. Environment configuration
3. CI/CD pipeline

## Migration Plan
1. Create new directory structure
2. Move and refactor modules
3. Update routing system
4. Implement new state management
5. Add tests
6. Update documentation