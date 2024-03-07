<script setup lang="ts">
defineProps<{
  msg: string
}>()
</script>

<template>
  <div>
    <h1 class="text-3xl">{{ msg }}</h1>
  </div>
</template>
