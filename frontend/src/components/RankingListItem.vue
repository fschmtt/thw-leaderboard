<script setup lang="ts">
import type { Competitor } from "@/views/Leaderboard.vue";

defineProps<{
  competitor: Competitor;
}>();
</script>

<template>
  <div class="rankingListItem">
    <p class="rank">#{{ competitor.rank }}</p>
    <p class="name">{{ competitor.name }}</p>
    <p class="score">{{ competitor.score.toFixed(2) }}</p>
  </div>
</template>

<style scoped>
.rankingListItem {
  display: grid;
  grid-template-columns: min-content 1fr max-content;
  padding: 16px;
  border: 1px solid #e6e6e6;
  gap: 16px;
  background-color: white;
  align-items: center;
}

.rank {
  background-color: #003399;
  color: white;
  font-weight: bold;
  font-size: 24px;
  line-height: 30px;
  padding: 6px 10px;
  text-align: center;
  min-width: 80px;
}
.name {
  font-size: 24px;
  line-height: 30px;
  text-overflow: ellipsis;
  white-space: nowrap;
  font-weight: bold;
  overflow: hidden;
}

.score {
  font-size: 24px;
  line-height: 30px;
  font-weight: medium;
  text-align: right;
}

.score::after {
  content: " cm";
}
</style>
