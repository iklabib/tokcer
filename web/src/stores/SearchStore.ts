import { defineStore } from 'pinia'
import { ref } from 'vue'
import type SearchItem from '@/types/SearchItem'
import type SearchType from '@/types/SearchType'

const API_URL = import.meta.env.VITE_API_URL
const options = [
  { name: 'Videos', code: 'videos' },
  { name: 'User', code: 'user' },
  { name: 'Tag', code: 'tag' },
]

export const useSearchStore = defineStore('search', () => {
  const items = ref<SearchItem[]>([])
  const isLoading = ref(false)
  const showLoadMore = ref(false)
  const keywords = ref('')
  const searchType = ref<SearchType>({ name: 'Videos', code: 'videos' })

  const search = async () => {
    isLoading.value = true
    showLoadMore.value = true
    try {
      const url = `${API_URL}/search`
      const options = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json', 'Accept': 'application/json' },
        body: JSON.stringify({
          keywords: keywords.value.trim(),
          search_type: searchType.value.code,
        }),
      }
      const response = await fetch(url, options)
      items.value = await response.json()
    } catch (error) {
      console.error(error)
      items.value = []
      showLoadMore.value = false
    } finally {
      isLoading.value = false
    }
  }

  const setKeywords = (keyword: string) => {
    keywords.value = keyword
  }

  const setSearchType = (type: string) => {
    for (const v of options) {
      if (v.code === type) {
        searchType.value = v
        break
      }
    }
  }

  return {
    items,
    isLoading,
    showLoadMore,
    keywords,
    searchType,
    search,
    setKeywords,
    setSearchType,
    options,
  }
})
