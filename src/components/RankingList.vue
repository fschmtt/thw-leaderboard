<script setup lang="ts">
import type { Competitor } from "../views/Leaderboard.vue";
import RankingListItem from "./RankingListItem.vue";

defineProps<{
  competitors: Array<Competitor>;
}>();
</script>

<template>
  <div class="rankingList">
    <RankingListItem
      v-for="competitor in competitors"
      :competitor="competitor"
    />
  </div>
</template>

<style scoped>
.rankingList {
  display: grid;
  grid-template-columns: 1fr;
  gap: 16px;
}
</style>
