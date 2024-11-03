<script setup lang="ts">
import Select from 'primevue/select'
import InputText from 'primevue/inputtext'
import IconField from 'primevue/iconfield'
import InputIcon from 'primevue/inputicon'
import ProgressBar from 'primevue/progressbar'
import { useSearchStore } from '../stores/SearchStore'
import router from '@/routes/routes'

const props = defineProps({
  redirect: Boolean,
})

const searchStore = useSearchStore()

const search = () => {
  if (!props.redirect) {
    searchStore.search()
  } else {
    const keywords = searchStore.keywords.trim()

    router.push({
      name: 'search',
      params: {
        keywords: keywords,
        type: 'videos',
      },
    })
  }
}
</script>

<template>
  <div>
    <ProgressBar
      mode="indeterminate"
      style="height: 6px"
      v-show="searchStore.isLoading"
    />
    <Toolbar style="padding: 5px">
      <template #start>
        <RouterLink to="/">Home</RouterLink>
      </template>

      <template #center>
        <div class="flex flex-center justify-center gap-2">
          <IconField>
            <InputIcon>
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
                class="icon icon-tabler icons-tabler-outline icon-tabler-search dark:text-white light:text-black pb-1"
              >
                <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                <path d="M10 10m-7 0a7 7 0 1 0 14 0a7 7 0 1 0 -14 0" />
                <path d="M21 21l-6 -6" />
              </svg>
            </InputIcon>
            <InputText
              placeholder="Search"
              class="w-[250px]"
              v-model="searchStore.keywords"
              @keyup.enter="search"
            />
          </IconField>
        </div>
      </template>
    </Toolbar>
  </div>
</template>
