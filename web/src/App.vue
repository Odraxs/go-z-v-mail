<script setup lang="ts">
import EmailsVisualizer from './components/EmailsVisualizer.vue'
import Footer from './components/Footer.vue'
import Header from './components/Header.vue'
</script>

<template>
  <Header msg="Go-z-v-mail" />
  <main class="flex flex-col min-h-[60rem]">
    <EmailsVisualizer />
  </main>
  <Footer />
</template>
