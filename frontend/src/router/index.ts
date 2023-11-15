import { createRouter, createWebHistory } from "vue-router"

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      component: () => import("@/views/Home.vue"),
    },
    {
      path: "/lobby",
      component: () => import("@/views/Lobby.vue"),
    },
    {
      path: "/game",
      component: () => import("@/views/Game.vue"),
    },
  ],
})

export default router
