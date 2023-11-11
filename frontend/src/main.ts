import { createApp } from "vue"
import App from "./App.vue"
import "@/assets/main.scss"

import "vuetify/styles"
import { createVuetify } from "vuetify"
import * as components from "vuetify/components"
import * as directives from "vuetify/directives"
import "@mdi/font/css/materialdesignicons.css"
import "vuetify/styles"
import { aliases, mdi } from "vuetify/iconsets/mdi"

import router from "@/router"

const vuetify = createVuetify({
  components,
  directives,
  icons: {
    defaultSet: "mdi",
    aliases,
    sets: {
      mdi,
    },
  },
})

createApp(App).use(vuetify).use(router).mount("#app")
