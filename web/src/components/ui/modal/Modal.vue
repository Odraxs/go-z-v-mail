<script setup lang="ts">
const emit = defineEmits(['close'])

function closeModal() {
  emit('close')
}
</script>

<template>
  <div
    class="flex justify-center items-center fixed top-0 left-0 right-0 z-50 bg-black bg-opacity-50 w-screen md:inset-0 h-screen max-h-full overflow-x-hidden overflow-y-auto"
    @click.self="closeModal"
  >
    <div class="p-2 md:p-10 md:mx-40 md:my-auto rounded-lg">
      <slot />
    </div>
  </div>
</template>
