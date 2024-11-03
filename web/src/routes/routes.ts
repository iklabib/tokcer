import { createRouter, createWebHistory } from 'vue-router'

import SearchPage from '@/views/SearchPage.vue'
import VideoPage from '@/views/VideoPage.vue'

const routes = [
  {
    path: '/',
    name: 'search',
    component: SearchPage,
  },
  {
    path: '/video/:user/:id',
    name: 'video',
    component: VideoPage,
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
