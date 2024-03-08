<script setup lang="ts">
import globalState from '@/helpers/globalState'
import type { Email } from '@/types'
import { computed } from 'vue'
import { InboxArrowDown, InboxStack } from '@/components/ui/icons'

const email = computed<Email | null>(() => globalState.getEmailInfo())
</script>

<template>
  <div class="flex flex-col justify-center items-start md:mx-6">
    <h2 class="text-2xl md:text-3xl font-semibold mb-10 md:mb-[3.2rem]">Email Details</h2>
    <section
      v-if="email !== null"
      class="dark:bg-gray-800 border border-gray-200 dark:border-gray-700 rounded-lg p-8 md:p-12 h-96 w-full max-h-96 md:h-full md:max-h-[40rem] overflow-y-auto"
    >
      <p
        class="bg-blue-100 text-blue-800 text-md font-medium inline-flex items-center px-2.5 py-0.5 rounded-md dark:bg-gray-700 dark:text-blue-400 mb-2"
      >
        ID {{ email.id }}
      </p>
      <p class="text-gray-700 dark:text-white text-sm font-semibold md:text-base my-2">
        <span class="font-bold text-gray-900">Date:</span>
        {{ email.date }}
      </p>
      <h1 class="text-gray-700 dark:text-white text-lg font-semibold md:text-xl mb-2">
        <span class="font-extrabold text-gray-900">Subject:</span>
        {{ email.subject }}
      </h1>
      <figcaption class="flex items-start mt-6 space-x-3 rtl:space-x-reverse">
        <div
          class="flex items-center divide-x-2 rtl:divide-x-reverse divide-gray-300 dark:divide-gray-700"
        >
          <cite
            class="flex gap-3 justify-center items-center pe-3 font-medium text-gray-900 dark:text-white"
            >From <InboxArrowDown class="w-5 h-5"
          /></cite>
          <cite class="ps-3 text-md text-gray-700 dark:text-gray-400">{{ email.from }}</cite>
        </div>
      </figcaption>
      <figcaption class="flex items-center mt-6 space-x-3 rtl:space-x-reverse">
        <div class="flex items-start">
          <cite
            class="flex gap-3 justify-center items-center pe-3 font-medium text-gray-900 dark:text-white mr-[0.06rem]"
            >To <InboxStack class="ml-5 w-5 h-5"
          /></cite>
          <cite
            class="ps-3 text-md text-gray-700 dark:text-gray-400 border-l-2 border-gray-300 dark:border-gray-700"
            >{{ email.to }}</cite
          >
        </div>
      </figcaption>
      <h2 class="my-2 p-2 font-bold text-base md:text-lg">Highlights:</h2>
      <blockquote
        v-for="highlight in email.highlight"
        :key="highlight"
        class="p-4 my-4 border-s-4 border-gray-100 bg-gray-50 dark:border-gray-500 dark:bg-gray-800"
      >
        <p
          class="text-base md:text-lg italic font-medium leading-relaxed text-gray-900 dark:text-white"
          v-html="highlight"
        ></p>
      </blockquote>
    </section>
  </div>
</template>
