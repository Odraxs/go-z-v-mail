<script setup lang="ts">
import Emails from './emails/Emails.vue'
import EmailDetails from './emails/EmailDetails.vue'
</script>

<template>
  <div class="md:grid md:grid-cols-7 justify-start items-start px-3 py-6">
    <Emails class="md:col-span-3" />
    <Transition
      enter-active-class="transition-opacity opacity-500 ease-in-out"
      leave-active-class="transition-opacity opacity-500 ease-in-out"
      enter-from-class="opacity-0"
      leave-to-class="opacity-0"
    >
      <EmailDetails class="mt-[21rem] px-8 col-span-4" />
    </Transition>
  </div>
</template>
