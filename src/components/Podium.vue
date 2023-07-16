<script setup lang="ts">
import type { Competitor } from "../views/Leaderboard.vue";

defineProps<{
  competitors: Array<Competitor>;
}>();
</script>

<template>
  <div class="podium-wrapper">
    <div
      class="podium-step"
      v-for="competitor in competitors"
      :class="`step-${competitor.rank}`"
    >
      <p class="rank">#{{ competitor.rank }}</p>
      <p class="name">{{ competitor.name ?? 'Spieler' }} ({{ competitor.identifier }})</p>
      <p class="score">{{ competitor.score.toFixed(2) }}</p>
      <span class="podium-step__background"></span>
    </div>
  </div>
</template>

<style scoped>
.podium-wrapper {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  grid-template-areas: "step-2 step-1 step-3";
  gap: 12px;
  min-height: 252px;
  margin-bottom: 12px;
}

.podium-step__background {
  background-color: white;
  border: 1px solid #e6e6e6;
  width: 100%;
  height: 100%;
}

.podium-step {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  flex-grow: 0;
}

.rank {
  margin-bottom: 8px;
  background-color: #003399;
  color: black;
  font-weight: bold;
  font-size: 24px;
  padding: 6px 10px;
  text-align: center;
  display: inline-block;
}

.step-1 {
  grid-area: step-1;
}

.step-2 {
  grid-area: step-2;
}

.step-3 {
  grid-area: step-3;
}

.step-1 .rank {
  background-color: #ffd173;
}

.step-2 .rank {
  background-color: #eaeaea;
  margin-top: 60px;
}

.step-3 .rank {
  background-color: #eccec6;
  margin-top: 100px;
}

.name {
  font-size: 24px;
  text-overflow: ellipsis;
  font-weight: bold;
  width: 100%;
  text-align: center;
}

.score {
  font-size: 18px;
  line-height: 30px;
  font-weight: medium;
  text-align: center;
  margin-bottom: 8px;
}

.score::after {
  content: " cm";
}
</style>
