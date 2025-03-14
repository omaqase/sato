<script setup lang="ts">
import InstagramIcon from './icons/InstagramIcon.vue'
import WhatsappIcon from './icons/WhatsappIcon.vue'
</script>

<template>
  <footer class="bg-gray-900 text-white">
    <div class="max-w-7xl mx-auto px-4 py-12">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8 mb-8">
        <div>
          <h3 class="font-medium mb-4">О компании</h3>
          <ul class="space-y-2 text-sm text-gray-400">
            <li>
              <router-link to="/about" class="hover:text-white transition-colors"
                >О нас</router-link
              >
            </li>
            <li>
              <router-link to="/careers" class="hover:text-white transition-colors"
                >Вакансии</router-link
              >
            </li>
            <li>
              <router-link to="/requisites" class="hover:text-white transition-colors"
                >Реквизиты</router-link
              >
            </li>
            <li>
              <router-link to="/press" class="hover:text-white transition-colors"
                >Пресс-центр</router-link
              >
            </li>
          </ul>
        </div>

        <div>
          <h3 class="font-medium mb-4">Покупателям</h3>
          <ul class="space-y-2 text-sm text-gray-400">
            <li>
              <router-link to="/how-to-order" class="hover:text-white transition-colors"
                >Как сделать заказ</router-link
              >
            </li>
            <li>
              <router-link to="/payment-methods" class="hover:text-white transition-colors"
                >Способы оплаты</router-link
              >
            </li>

            <li>
              <router-link to="/delivery" class="hover:text-white transition-colors"
                >Доставка</router-link
              >
            </li>

            <li>
              <router-link to="/return-policy" class="hover:text-white transition-colors"
                >Возврат товара</router-link
              >
            </li>
          </ul>
        </div>

        <div>
          <h3 class="font-medium mb-4">Партнерам</h3>
          <ul class="space-y-2 text-sm text-gray-400">
            <li>
              <router-link to="/sell" class="hover:text-white transition-colors"
                >Продавайте на Sato</router-link
              >
            </li>
            <li>
              <router-link to="/transport" class="hover:text-white transition-colors"
                >Перевозчикам</router-link
              >
            </li>
            <li>
              <router-link to="/pickup-point" class="hover:text-white transition-colors"
                >Партнерский пункт выдачи</router-link
              >
            </li>
          </ul>
        </div>

        <div>
          <h3 class="font-medium mb-4">Контакты</h3>
          <ul class="space-y-2 text-sm text-gray-400">
            <li>+7 777 777 77 77</li>
            <li>support@sato.kz</li>
          </ul>
          <div class="flex items-center gap-4 mt-4">
            <a href="#" class="p-2 rounded-full bg-gray-800 hover:bg-gray-700 transition-colors">
              <InstagramIcon />
            </a>
            <a href="#" class="p-2 rounded-full bg-gray-800 hover:bg-gray-700 transition-colors">
              <WhatsappIcon />
            </a>
          </div>
        </div>
      </div>

      <div class="pt-8 border-t border-gray-800 text-sm text-gray-400">
        <div class="flex flex-col md:flex-row justify-between items-center gap-4">
          <div>© 2024 Sato. Все права защищены</div>
          <div class="flex items-center gap-4">
            <router-link to="/privacy-policy" class="hover:text-white transition-colors"
              >Политика конфиденциальности</router-link
            >
            <router-link to="/terms" class="hover:text-white transition-colors">
              Пользовательское соглашение
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </footer>
</template>
