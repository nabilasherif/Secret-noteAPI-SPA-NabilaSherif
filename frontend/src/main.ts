import { createApp } from 'vue'
import Toast from "vue-toastification"
import 'vue-toastification/dist/index.css'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'
loadFonts()

createApp(App)
  .use(Toast)
  .use(router)
  .use(vuetify)
  .mount('#app')