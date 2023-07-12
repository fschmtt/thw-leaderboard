import {createRouter, createWebHashHistory, createWebHistory} from "vue-router";
import Admin from "../views/Admin.vue";
import Leaderboard from "../views/Leaderboard.vue";

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "leaderboard",
      component: Leaderboard,
    },
    {
      path: "/admin",
      name: "admin",
      component: Admin,
    },
  ],
});

export default router;
