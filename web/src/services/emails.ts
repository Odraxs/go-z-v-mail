import type { EmailSearchResponse } from '@/types'

type AllowedFields = 'content' | 'from' | 'subject'
type AllowedSortFields = 'date' | 'from' | 'subject' | '-date' | '-from' | '-subject'

interface FormData {
  term: string
  field: AllowedFields
  sort?: string
  order?: string
  maxResults: string
}

interface RequestBody {
  term: string
  field: AllowedFields
  sort_fields?: AllowedSortFields[]
  max_results: number
}

// Change to get It from .env file
const searchEmailsEndpoint = 'http://localhost:3001/emailSearch'

async function searchEmails({ term, field, sort, order, maxResults }: FormData) {
  const requestBody: RequestBody = {
    term,
    field,
    max_results: Number(maxResults),
    sort_fields: processSortField(sort, order)
  }

  console.log(requestBody)

  return fetch(searchEmailsEndpoint, {
    method: 'POST',
    headers: {
      'Content-type': 'application/json'
    },
    body: JSON.stringify(requestBody)
  })
    .then(async (response) => {
      if (!response.ok) {
        throw new Error('Failed to search emails')
      }
      return await response.json()
    })
    .then((data) => {
      const { emails }: EmailSearchResponse = data
      return emails
    })
    .catch((error) => {
      console.error(error)
      throw error
    })
  //return emailSearchResponse.emails
}

function processSortField(sort: string | undefined, order: string | undefined) {
  if (sort === undefined) {
    return []
  }
  if (order === 'asc') {
    return [sort] as AllowedSortFields[]
  }
  if (order === 'desc') {
    return [`-${sort}`] as AllowedSortFields[]
  }
}

export { searchEmails }
