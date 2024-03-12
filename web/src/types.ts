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
  highlight: HighlightContent
}

interface FormattedEmail extends Omit<Email, 'highlight'> {
  highlight: string[]
}

interface HighlightContent {
  content?: string[]
  subject?: string[]
  from?: string[]
  to?: string[]
}

export type { EmailSearchResponse, Email, HighlightContent, FormattedEmail }
