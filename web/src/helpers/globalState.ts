import type { Email } from '@/types'
import { ref } from 'vue'

interface GlobalState {
  emailInfo: Email | null
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

  const formattedHighlights = emailInfo.highlight.map((highlight) =>
    highlight.replace(/\n/g, '<br>').replace(/\r/g, '&#x0D;')
  )

  return { ...emailInfo, content: formattedContent, highlight: formattedHighlights }
}

export default globalState
