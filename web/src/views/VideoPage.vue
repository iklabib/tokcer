<script setup lang="ts">
import { ref, onMounted } from 'vue'
import Card from 'primevue/card'
import { useRoute } from 'vue-router'
import SearchBox from '@/components/SearchBox.vue'
import type { VideoInfo } from '@/types/VideoInfo'
import CardContentSkeleton from '@/components/CardContentSkeleton.vue'
import CardRelatedVideoSkeleton from '@/components/CardRelatedVideoSkeleton.vue'

import 'vidstack/player/styles/default/theme.css'
import 'vidstack/player/styles/default/layouts/video.css'
import 'vidstack/player'
import 'vidstack/player/layouts/default'
import 'vidstack/player/ui'

const API_URL = import.meta.env.VITE_API_URL

const route = useRoute()
const user = route.params.user as string
const videoId = route.params.id as string
const url = `${API_URL}/stream/${user}/${videoId}`

const videoInfo = ref<VideoInfo>()
const isLoadingManfest = ref(false)

onMounted(async () => {
  try {
    isLoadingManfest.value = true
    const url = `${API_URL}/video/${user}/${videoId}`
    const options = { method: 'GET' }
    const response = await fetch(url, options)
    videoInfo.value = await response.json()
  } catch (error) {
    console.error(error)
  } finally{
    isLoadingManfest.value = false
  }
})
</script>

<template>
  <div class="container-lg p-2">
    <SearchBox :redirect="true" />

    <div class="mt-2 grid sm:grid-cols-1 lg:grid-cols-6 gap-2">
      <div class="base:col-span-2 md:col-span-4 row-span-12">
        <Card>
          <template #header>
            <media-player viewType="video" streamType="on-demand" :src="url">
              <media-provider></media-provider>
              <media-video-layout></media-video-layout>
            </media-player>
          </template>

          <template #content>
            <CardContentSkeleton v-show="isLoadingManfest"/>
            
            <div class="flex" v-show="!isLoadingManfest">
              <div class="flex flex-col justify-center">
                <img
                  class="h-[40px]"
                  :src="videoInfo?.author.avatarThumb"
                  :alt="videoInfo?.author.uniqueId"
                />
              </div>

              <div class="ml-2">
                <b class="text-lg">{{ videoInfo?.author.uniqueId }}</b>

                <div>
                  {{ videoInfo?.author.nickname }}
                </div>
              </div>
            </div>

            <p class="mt-2">
              {{ videoInfo?.desc }}
            </p>
          </template>
        </Card>
      </div>

      <CardRelatedVideoSkeleton v-for="idx in 6" v-show="isLoadingManfest"/>

      <Card v-for="item in videoInfo?.relatedVideos">
        <template #header>
          <div
            class="relative w-full h-[250px] bg-white dark:bg-black overflow-hidden aspect-[9/16]"
          >
            <img
              :src="item.thumbnail"
              :alt="item.title"
              class="absolute inset-0 w-full object-cover object-center"
            />
          </div>
        </template>

        <template #content>
          <p class="m-0 p-0 line-clamp-4 text-ellipsis overflow-hidden">
            {{ item.title }}
          </p>
        </template>
      </Card>
    </div>
  </div>
</template>
