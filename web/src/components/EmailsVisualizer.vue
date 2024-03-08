<script setup lang="ts">
import Emails from './emails/Emails.vue'
import EmailDetails from './emails/EmailDetails.vue'

defineProps<{
  msg: string
}>()
</script>

<template>
  <h1 class="text-3xl w-full bg-black p-2 text-white">{{ msg }}</h1>
  <div class="md:grid md:grid-cols-7 justify-start items-start px-3 py-6">
    <Emails class="md:col-span-3" />
    <EmailDetails class="mt-14 px-8 col-span-4" />
  </div>
</template>
