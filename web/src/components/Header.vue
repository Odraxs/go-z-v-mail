<script setup lang="ts">
defineProps<{
  msg: string
}>()
</script>

<template>
  <header class="text-3xl w-full bg-gray-900 p-3 text-white">
    <a href="" class="cursor-pointer">{{ msg }}</a>
  </header>
</template>
