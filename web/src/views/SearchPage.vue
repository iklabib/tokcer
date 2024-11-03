<script setup lang="ts">
import Card from 'primevue/card'
import Button from 'primevue/button'
import SearchBox from '../components/SearchBox.vue'
import { useSearchStore } from '../stores/SearchStore'
import router from '../routes/routes'
import { useRoute } from 'vue-router'

const searchStore = useSearchStore()
const route = useRoute()
const keywords = route.params.keywords as string

if (keywords) {
  searchStore.setKeywords(keywords)
} else {
  searchStore.initPage = true
  searchStore.setKeywords('')
}

const getVideoId = (url: string): string => {
  const parsed = URL.parse(url)
  if (parsed == null) {
    return url
  }
  return parsed.pathname.replace(/^\/+|\/+$/g, '')
}

const openVideo = (url: string) => {
  const parsed = URL.parse(url)
  if (parsed == null) {
    return
  }

  const pathname = parsed.pathname.replace(/^\/+|\/+$/g, '')
  // username/video/videoId
  const [user, , videoId] = pathname.split('/')

  router.push({
    name: 'video',
    params: {
      user: user,
      id: videoId,
    },
  })
}
</script>

<template>
  <div class="container-lg p-2">
    <SearchBox />

    <div class="flex flex-center justify-center mt-[30vh] gap-2 opacity-40" v-if="searchStore.initPage">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 24 24"
        width="1.25rem"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
        class="icon icon-tabler icons-tabler-outline icon-tabler-search dark:text-white light:text-black pb-1 w-12"
      >
        <path stroke="none" d="M0 0h24v24H0z" fill="none" />
        <path d="M10 10m-7 0a7 7 0 1 0 14 0a7 7 0 1 0 -14 0" />
        <path d="M21 21l-6 -6" />
      </svg>

      <h1 class="text-4xl">Start Searching</h1>
    </div>

    <div class="mt-4"></div>

    <div class="grid mt-4 sm:grid-cols-1 md:grid-cols-4 lg:grid-cols-6 gap-4">
      <Card
        v-for="item in searchStore.items"
        :key="getVideoId(item.url)"
        @click="openVideo(item.url)"
      >
        <template #header>
          <div
            class="relative bg-white dark:bg-black w-full overflow-hidden aspect-[9/16]"
          >
            <img
              :alt="item.coverAlt"
              :src="item.cover"
              class="absolute inset-0 w-full h-full object-cover object-center"
            />
          </div>
        </template>
        <template #title>
          <div class="flex gap-x-2">
            <img :src="item.userAvatar" :alt="item.username" width="24" />
            <div class="text-ellipsis overflow-hidden">
              {{ item.username }}
            </div>
          </div>
        </template>
        <template #content>
          <p class="m-0 line-clamp-4">{{ item.desc }}</p>
        </template>
      </Card>
    </div>
    
    <div class="flex justify-center mt-4">
      <Button
        label="Load More"
        @click="searchStore.search()"
        :loading="searchStore.isLoading"
        v-show="searchStore.showLoadMore"
      />
    </div>
  </div>
</template>
