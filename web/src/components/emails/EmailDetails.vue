<script setup lang="ts">
import globalState from '@/helpers/globalState'
import type { FormattedEmail } from '@/types'
import { computed, ref } from 'vue'
import EmailBasicInfo from './EmailBasicInfo.vue'
import { Button } from '@/components/ui/button'
import { Modal } from '@/components/ui/modal'

const showModal = ref(false)
const email = computed<FormattedEmail | null>(() => globalState.getEmailInfo())

function toggleModal() {
  showModal.value = !showModal.value
}
</script>

<template>
  <div>
    <EmailBasicInfo v-if="email !== null" :email="email" title="Email details">
      <template v-slot:button>
        <Button size="lg" class="bg-blue-700 text-white hover:bg-blue-900" @click="toggleModal">
          See full email
        </Button>
      </template>
    </EmailBasicInfo>
    <Teleport to=".modals" v-if="showModal">
      <Modal @close="toggleModal">
        <EmailBasicInfo
          v-if="email !== null"
          :email="email"
          title="Full Email"
          class="bg-gray-100 rounded-lg p-10"
        >
          <template v-slot:content>
            <h2 class="my-2 p-2 font-bold text-lg md:text-xl">Content:</h2>
            <article class="bg-gray-50 p-4 font-medium" v-html="email.content"></article>
          </template>
          <template v-slot:button>
            <Button size="lg" variant="destructive" class="my-4" @click="toggleModal">
              Close
            </Button>
          </template>
        </EmailBasicInfo>
      </Modal>
    </Teleport>
  </div>
</template>
