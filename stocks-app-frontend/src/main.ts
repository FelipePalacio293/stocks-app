import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import { createI18n } from 'vue-i18n'

import es from './locales/es.json'
import en from './locales/en.json'

const app = createApp(App)

const i18n = createI18n({
    locale: 'en',
    fallbackLocale: 'en',
    messages: {
        es,
        en
    }
})

app.use(createPinia())
app.use(router)
app.use(i18n)

app.mount('#app')
