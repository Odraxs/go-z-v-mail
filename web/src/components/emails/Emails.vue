<script setup lang="ts">
import { type Email } from '@/types'
import { columns } from './columns'
import DataTable from './DataTable.vue'
import { toTypedSchema } from '@vee-validate/zod'
import * as z from 'zod'
import { Input } from '@/components/ui/input'
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage
} from '@/components/ui/form'
import { Button } from '@/components/ui/button'
import { MagnifyingGlass } from '@/components/ui/icons'
import { ref } from 'vue'
import { searchEmails } from '@/services/emails'

const requestState = ref({
  isLoading: false,
  error: null
})
const emails = ref<Email[]>([])
const formSchema = toTypedSchema(
  z.object({
    filter: z.string().min(3)
  })
)
async function handleSubmit(values: any) {
  console.log('Form submitted!', values)
  try {
    requestState.value = {
      isLoading: true,
      error: null
    }
    emails.value = await searchEmails(values)
  } catch (error: any) {
    requestState.value.error = error.message || 'Error searching emails'
  } finally {
    requestState.value.isLoading = false
  }
}
</script>

<template>
  <div class="grid grid-cols-1 justify-center items-start">
    <Form @submit="handleSubmit" :validation-schema="formSchema" class="px-3 py-6">
      <FormField v-slot="{ componentField }" name="filter">
        <FormItem>
          <FormLabel class="text-base ml-3">Phrase filter</FormLabel>
          <div class="flex flex-row justify-between gap-2 px-2">
            <FormControl class="flex-grow">
              <div>
                <Input type="text" placeholder="Set the phrase filter" v-bind="componentField" />
                <FormMessage class="mt-2 ml-3" />
              </div>

              <Button
                size="icon"
                variant="outline"
                type="submit"
                class="bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 dark:focus:ring-blue-900"
              >
                <MagnifyingGlass class="text-white w-6 h-6" />
              </Button>
            </FormControl>
          </div>
        </FormItem>
      </FormField>
    </Form>

    <!-- I wanted to use Suspense, however in vue, that feature is still an experimental one. -->
    <div class="px-3 py-6" v-if="emails.length > 0 || requestState.isLoading">
      <DataTable :columns="columns" :data="emails" />
    </div>
    <div v-if="requestState.isLoading && requestState.error === null">
      <strong class="text-md font-bold">Loading...</strong>
    </div>
    <div v-if="!requestState.isLoading && requestState.error !== null">
      <strong class="text-lg font-bold">{{ requestState.error }}</strong>
    </div>
  </div>
</template>
