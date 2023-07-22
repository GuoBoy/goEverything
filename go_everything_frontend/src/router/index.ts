import { createRouter, createWebHashHistory } from "vue-router";
import Index from "../views/index.vue"

const routes = [
  {
    path: "/",
    component: Index
  }, {
    path: "/diskManage",
    component: () => import("../views/diskManage.vue")
  }
]

export const router = createRouter({
  history: createWebHashHistory(),
  routes
})

