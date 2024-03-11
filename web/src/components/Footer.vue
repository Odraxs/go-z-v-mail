<script setup lang="ts">
import { ref } from 'vue'
import contactsData from '../staticData/contacts.json'

interface ContactType {
  id: string
  contact: string
  link: string
  icon: string
}

const contacts = ref<ContactType[]>(contactsData.contacts)
</script>

<template>
  <footer className="bg-gray-900">
    <div className="mx-auto w-full max-w-screen-xl p-4 py-6 lg:py-8">
      <div>
        <div className="mb-6 md:mb-0">
          <img
            src="https://www.vectorlogo.zone/logos/vuejs/vuejs-icon.svg"
            className="h-8 me-3"
            alt="Logo"
          />
          <span className="self-center text-2xl font-semibold whitespace-nowrap text-white"
            >David Roldán</span
          >
        </div>
        <div className="grid grid-cols-[repeat(auto-fill,minmax(80px,1fr))] gap-6 mt-5">
          <div>
            <h2 className="mb-6 text-sm font-semibold uppercase text-white">Contacts</h2>
            <ul
              className="flex flex-wrap items-center mb-6 text-sm font-medium sm:mb-0 text-gray-400"
            >
              <li v-for="{ id, contact, icon, link } in contacts" :key="id">
                <a
                  :href="link"
                  target="_blank"
                  className="flex items-center mb-4 hover:underline capitalize"
                  rel="noreferrer"
                  >{{ contact }}
                  <img :src="icon" className="w-4 h-4 mx-2" />
                </a>
              </li>
            </ul>
          </div>
        </div>
      </div>
      <hr className="my-6 sm:mx-auto border-gray-700 lg:my-8" />
      <div className="sm:flex sm:items-center sm:justify-between">
        <span className="text-sm sm:text-center text-gray-400">
          © 2024 Odraxs. All Rights Reserved.
        </span>
      </div>
    </div>
  </footer>
</template>
