import {createRouter, createWebHashHistory} from "vue-router";
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
      path: "/s3cr3t-4dm1n",
      name: "admin",
      component: Admin,
    },
  ],
});

export default router;
