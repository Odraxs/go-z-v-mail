<!-- // Reusable component -->
<script setup lang="ts" generic="TData, TValue">
import type { ColumnDef, SortingState } from '@tanstack/vue-table'
import {
  FlexRender,
  getCoreRowModel,
  useVueTable,
  getPaginationRowModel,
  getSortedRowModel
} from '@tanstack/vue-table'

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow
} from '@/components/ui/table'
import { ref } from 'vue'
import { valueUpdater } from '@/lib/utils'
import DataTablePagination from './DataTablePagination.vue'

interface DataTableProps {
  columns: ColumnDef<TData, TValue>[]
  data: TData[]
}

const props = defineProps<DataTableProps>()

const sorting = ref<SortingState>([])

const table = useVueTable({
  get data() {
    return props.data
  },
  get columns() {
    return props.columns
  },
  getCoreRowModel: getCoreRowModel(),
  getPaginationRowModel: getPaginationRowModel(),
  getSortedRowModel: getSortedRowModel(),
  enableRowSelection: true,
  enableMultiRowSelection: false,
  onSortingChange: (updaterOrValue) => valueUpdater(updaterOrValue, sorting),
  state: {
    get sorting() {
      return sorting.value
    }
  }
})
</script>

<template>
  <div class="space-y-2">
    <div class="border rounded-md">
      <Table>
        <TableHeader>
          <TableRow v-for="headerGroup in table.getHeaderGroups()" :key="headerGroup.id">
            <TableHead v-for="header in headerGroup.headers" :key="header.id">
              <FlexRender
                v-if="!header.isPlaceholder"
                :render="header.column.columnDef.header"
                :props="header.getContext()"
              />
            </TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <template v-if="table.getRowModel().rows?.length">
            <TableRow
              v-for="row in table.getRowModel().rows"
              :key="row.id"
              :data-state="row.getIsSelected() ? 'selected' : undefined"
              class="data-[state=selected]:bg-blue-200"
            >
              <TableCell v-for="cell in row.getVisibleCells()" :key="cell.id">
                <FlexRender :render="cell.column.columnDef.cell" :props="cell.getContext()" />
              </TableCell>
            </TableRow>
          </template>
          <template v-else>
            <TableRow>
              <TableCell :colSpan="columns.length" class="h-24 text-center">
                No results.
              </TableCell>
            </TableRow>
          </template>
        </TableBody>
      </Table>
    </div>
    <DataTablePagination :table="table" />
  </div>
</template>
