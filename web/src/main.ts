import { createApp } from 'vue'
import { createPinia } from 'pinia'

import PrimeVue from 'primevue/config'
import Lara from '@primevue/themes/lara'
import router from './routes/routes'

import App from './App.vue'

const app = createApp(App)

app.use(router)
app.use(createPinia())
app.use(PrimeVue, {
  theme: {
    preset: Lara,
  },
})

app.mount('#app')
