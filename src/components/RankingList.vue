<script setup lang="ts">
import type { Competitor } from "../views/Leaderboard.vue";
import RankingListItem from "./RankingListItem.vue";

defineProps<{
  competitors: Array<Competitor>;
}>();
</script>

<template>
  <RankingListItem v-for="competitor in competitors" :competitor="competitor" />
</template>

<style scoped></style>
