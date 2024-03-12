import type { Email, FormattedEmail, HighlightContent } from '@/types'
import { ref } from 'vue'

interface GlobalState {
  emailInfo: FormattedEmail | null
}

const globalState = {
  state: ref<GlobalState>({
    emailInfo: null
  }),
  getEmailInfo() {
    return this.state.value.emailInfo
  },
  updateEmailInfo(emailInfo: Email) {
    this.state.value.emailInfo = formatEmailInfo(emailInfo)
  },
  resetEmailInfo() {
    this.state.value.emailInfo = null
  }
}

function formatEmailInfo(emailInfo: Email) {
  const formattedContent = emailInfo.content
    .replace(/-/g, '')
    .replace(/\n/g, '<br>')
    .replace(/\r/g, '&#x0D;')
    .replace(/\t/g, '&nbsp;')
    .replace(/(&nbsp;)+/g, '&nbsp;')
    .replace(/(&#x0D;<br>)+/g, '<br>')

  const [key] = Object.keys(emailInfo.highlight) as (keyof HighlightContent)[]

  const formattedHighlights = emailInfo.highlight[key]!.map((highlight) =>
    highlight.replace(/\n/g, '<br>').replace(/\r/g, '&#x0D;')
  )

  const formattedDate = new Date(emailInfo.date).toUTCString()

  return {
    ...emailInfo,
    content: formattedContent,
    highlight: formattedHighlights,
    date: formattedDate
  }
}

export default globalState
