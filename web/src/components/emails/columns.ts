import type { Email } from '@/types'
import type { ColumnDef } from '@tanstack/vue-table'
import { h } from 'vue'
import { Checkbox } from '@/components/ui/checkbox'
import globalState from '@/helpers/globalState'

export const columns: ColumnDef<Email>[] = [
  {
    id: 'select',
    cell: ({ row }) =>
      h(Checkbox, {
        checked: row.getIsSelected(),
        'onUpdate:checked': (value) => {
          value ? globalState.updateEmailInfo(row.original) : globalState.resetEmailInfo()
          row.toggleSelected(!!value)
        },
        ariaLabel: 'Select row',
        class: 'flex data-[state=checked]:bg-blue-400'
      })
  },
  {
    accessorKey: 'id',
    header: () => h('div', { class: 'text-left' }, 'ID'),
    cell: ({ row }) => {
      const id: number = row.getValue('id')

      return h('div', { class: 'text-left font-medium' }, id)
    }
  },
  {
    accessorKey: 'from',
    header: () => h('div', { class: 'text-left' }, 'From'),
    cell: ({ row }) => {
      const from: string = row.getValue('from')

      return h('div', { class: 'line-clamp-1 text-left font-medium' }, from)
    }
  },
  {
    accessorKey: 'to',
    header: () => h('div', { class: 'text-left' }, 'To'),
    cell: ({ row }) => {
      const to: string = row.getValue('to')

      return h('p', { class: 'line-clamp-1 text-left font-medium max-w-36' }, to)
    }
  },
  {
    accessorKey: 'date',
    header: () => h('div', { class: 'text-left' }, 'Date'),
    cell: ({ row }) => {
      const date: string = row.getValue('date')

      return h('div', { class: 'text-left font-medium' }, new Date(date).toDateString())
    }
  }
]
