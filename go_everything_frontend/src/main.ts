import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import { router } from './router'

import Tag from "./components/Tag.vue"

createApp(App).component("tag", Tag).use(router).mount('#app')
