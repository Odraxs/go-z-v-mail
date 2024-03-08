interface EmailSearchResponse {
  time: number
  emails: Email[]
}

interface Email {
  id: string
  from: string
  to: string
  content: string
  subject: string
  date: string
  highlight: string[]
}

export type { EmailSearchResponse, Email }
