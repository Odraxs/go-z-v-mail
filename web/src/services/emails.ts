import type { EmailSearchResponse } from '@/types'

// Change to get It from .env file
const searchEmailsEndpoint = 'http://localhost:3001/emailSearch'

async function searchEmails(requestBody: { filter: string }) {
  // Fetch data from API
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
      console.log(data)
      const { emails }: EmailSearchResponse = data
      return emails
    })
    .catch((error) => {
      console.error(error)
      throw error
    })
}

export { searchEmails }
