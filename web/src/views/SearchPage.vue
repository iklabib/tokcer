<script setup lang="ts">
import Card from 'primevue/card'
import Button from 'primevue/button'
import SearchBox from '../components/SearchBox.vue'
import { useSearchStore } from '../stores/SearchStore'
import router from '@/routes/routes'
import { useRoute } from 'vue-router'

const searchStore = useSearchStore()

const route = useRoute()

const keywords = route.params.keywords as string
const searchType = route.params.type as string

if (keywords && searchType) {
  searchStore.setKeywords(keywords)
  searchStore.setSearchType(searchType)
} else {
  searchStore.setKeywords('')
  searchStore.setSearchType('videos')
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
    <div class="grid sm:grid-cols-1 md:grid-cols-4 lg:grid-cols-6 gap-4">
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
