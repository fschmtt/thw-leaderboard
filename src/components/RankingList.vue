<script setup lang="ts">
import type { Competitor } from "../views/Leaderboard.vue";

defineProps<{
  competitors: Array<Competitor>;
}>();
</script>

<template>
  <RankingListItem v-for="competitor in competitors" />
</template>

<style scoped></style>
