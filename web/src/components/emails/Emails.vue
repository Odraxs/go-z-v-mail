<script setup lang="ts">
import type { Email } from '@/types'
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
import {
  Select,
  SelectTrigger,
  SelectValue,
  SelectContent,
  SelectGroup,
  SelectItem
} from '@/components/ui/select'
import TableSkeleton from './TableSkeleton.vue'
import globalState from '@/helpers/globalState'

const requestState = ref({
  isLoading: false,
  error: null
})

const allowedFields = ['content', 'subject', 'from'] as const
const allowedSortFields = ['date', 'from', 'subject'] as const
const orders = ['asc', 'desc'] as const
const maxResultsOptions = ['20', '50', '100', '500', '1000'] as const
const defaultField = allowedFields[0]
const defaultMaxResult = maxResultsOptions[0]
const defaultOrder = orders[0]
const emails = ref<Email[] | null>(null)
const needsOrder = ref<string>()
const maxResults = ref(defaultMaxResult)

const formSchema = toTypedSchema(
  z.object({
    term: z.string().min(3),
    field: z.enum(allowedFields).default(defaultField),
    sort: z.optional(z.enum(allowedSortFields)),
    order: z.enum(orders).default(defaultOrder),
    maxResults: z.enum(maxResultsOptions).default(defaultMaxResult)
  })
)

async function handleSubmit(values: any) {
  try {
    requestState.value = {
      isLoading: true,
      error: null
    }
    emails.value = await searchEmails(values)
  } catch (error: any) {
    requestState.value.error = error.message || 'Error searching emails'
  } finally {
    globalState.resetEmailInfo()
    requestState.value.isLoading = false
  }
}
</script>

<template>
  <div class="grid grid-cols-1 justify-center items-start">
    <Form @submit="handleSubmit" :validation-schema="formSchema" class="grid gap-3 px-3 py-6">
      <div class="grid grid-cols-1 lg:grid-cols-2 lg:gap-6">
        <FormField v-slot="{ componentField }" name="field">
          <FormItem class="my-3">
            <FormLabel class="text-base ml-3">Email field to search</FormLabel>

            <Select v-bind="componentField">
              <FormControl>
                <SelectTrigger>
                  <SelectValue class="capitalize" :placeholder="`${defaultField}`" />
                </SelectTrigger>
              </FormControl>
              <SelectContent>
                <SelectGroup>
                  <SelectItem
                    v-for="field in allowedFields"
                    :value="field"
                    :key="field"
                    class="capitalize"
                  >
                    {{ field }}
                  </SelectItem>
                </SelectGroup>
              </SelectContent>
            </Select>
            <FormMessage class="mt-2 ml-3" />
          </FormItem>
        </FormField>
        <FormField v-slot="{ componentField }" name="maxResults">
          <FormItem class="my-3">
            <FormLabel class="text-base ml-3">Max amount of result</FormLabel>

            <Select v-bind="componentField">
              <FormControl>
                <SelectTrigger>
                  <SelectValue class="capitalize" :placeholder="maxResults" />
                </SelectTrigger>
              </FormControl>
              <SelectContent>
                <SelectGroup>
                  <SelectItem
                    v-for="options in maxResultsOptions"
                    :value="options"
                    :key="options"
                    class="capitalize"
                  >
                    {{ options || `All` }}
                  </SelectItem>
                </SelectGroup>
              </SelectContent>
            </Select>
            <FormMessage class="mt-2 ml-3" />
          </FormItem>
        </FormField>
      </div>
      <div class="flex flex-row gap-2 lg:gap-6">
        <FormField v-slot="{ componentField }" name="sort" v-model="needsOrder">
          <FormItem class="w-full my-3">
            <FormLabel class="text-base ml-3">Filed to order</FormLabel>

            <Select v-bind="componentField">
              <FormControl>
                <SelectTrigger>
                  <SelectValue class="capitalize" :placeholder="`Select an order`" />
                </SelectTrigger>
              </FormControl>
              <SelectContent>
                <SelectGroup>
                  <SelectItem
                    v-for="field in allowedSortFields"
                    :value="field"
                    :key="field"
                    class="capitalize"
                  >
                    {{ field }}
                  </SelectItem>
                </SelectGroup>
              </SelectContent>
            </Select>
            <FormMessage class="mt-2 ml-3" />
          </FormItem>
        </FormField>
        <FormField v-slot="{ componentField }" name="order">
          <FormItem class="w-28 my-3">
            <FormLabel class="text-base ml-3">Order</FormLabel>

            <Select v-bind="componentField" :disabled="needsOrder === undefined">
              <FormControl>
                <SelectTrigger>
                  <SelectValue class="capitalize" :placeholder="defaultOrder" />
                </SelectTrigger>
              </FormControl>
              <SelectContent>
                <SelectGroup>
                  <SelectItem
                    v-for="field in orders"
                    :value="field"
                    :key="field"
                    class="capitalize"
                  >
                    {{ field }}
                  </SelectItem>
                </SelectGroup>
              </SelectContent>
            </Select>
            <FormMessage class="mt-2 ml-3" />
          </FormItem>
        </FormField>
      </div>

      <FormField v-slot="{ componentField }" name="term">
        <FormItem>
          <FormLabel class="text-base ml-3">Keywords</FormLabel>
          <div class="flex flex-row justify-between gap-2">
            <FormControl class="flex-grow">
              <div>
                <Input
                  type="text"
                  placeholder="Set the Keywords you want to look to"
                  v-bind="componentField"
                />
                <FormMessage class="mt-2 ml-3" />
              </div>
            </FormControl>
          </div>
        </FormItem>
      </FormField>
      <Button
        size="lg"
        variant="outline"
        type="submit"
        class="bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 dark:focus:ring-blue-900 text-white hover:text-white w-40 mt-2"
      >
        Search <MagnifyingGlass class="ml-2 w-6 h-6" />
      </Button>
    </Form>

    <div v-if="requestState.isLoading && requestState.error === null" class="px-3 py-6">
      <TableSkeleton />
    </div>
    <div v-if="!requestState.isLoading && requestState.error !== null">
      <strong class="text-lg font-bold">{{ requestState.error }}</strong>
    </div>
    <!-- I wanted to use Suspense, however in vue, that feature is still an experimental one. -->
    <div
      class="px-3 py-6"
      v-if="emails !== null && !requestState.isLoading && requestState.error === null"
    >
      <DataTable :columns="columns" :data="emails" />
    </div>
  </div>
</template>
