<template>
  <div class="min-h-screen bg-gray-50">
    <nav class="bg-white shadow-sm">
      <div class="max-w-7xl mx-auto px-4">
        <div class="flex justify-between h-16">
          <div class="flex">
            <router-link
              to="/admin"
              class="flex items-center px-2 py-2 text-gray-900 font-medium"
            >
              Админ панель
            </router-link>

            <div class="ml-6 flex space-x-4">
              <router-link
                to="/admin/categories"
                class="px-3 py-2 rounded-md text-sm font-medium"
                :class="[
                  $route.name === 'admin-categories'
                    ? 'bg-gray-100 text-gray-900'
                    : 'text-gray-500 hover:text-gray-900',
                ]"
              >
                Категории
              </router-link>

              <router-link
                to="/admin/products"
                class="px-3 py-2 rounded-md text-sm font-medium"
                :class="[
                  $route.name === 'admin-products'
                    ? 'bg-gray-100 text-gray-900'
                    : 'text-gray-500 hover:text-gray-900',
                ]"
              >
                Товары
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </nav>

    <main class="py-10">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <router-view />
      </div>
    </main>
  </div>
</template> 