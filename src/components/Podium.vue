<script setup lang="ts">
import type { Competitor } from "../views/Leaderboard.vue";

defineProps<{
  competitors: Array<Competitor>;
}>();
</script>

<template>
  <div class="podium-wrapper" v-if="competitors.length > 0">
    <div class="podium-step step-2">
      <p class="rank">#2</p>
      <p class="name">{{ competitors[1].name }}</p>
      <p class="score">{{ competitors[1].faultPoints }}</p>
      <span class="podium-step__background"></span>
    </div>
    <div class="podium-step step-1">
      <p class="rank">#1</p>
      <p class="name">{{ competitors[0].name }}</p>
      <p class="score">{{ competitors[0].faultPoints }}</p>
      <span class="podium-step__background"></span>
    </div>
    <div class="podium-step step-3">
      <p class="rank">#3</p>
      <p class="name">{{ competitors[2].name }}</p>
      <p class="score">{{ competitors[2].faultPoints }}</p>
      <span class="podium-step__background"></span>
    </div>
  </div>
</template>

<style scoped>
.podium-wrapper {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  gap: 24px;
  height: 260px;
  margin-bottom: 16px;
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

.step-1 .rank {
  background-color: #ffd173;
}

.step-2 .rank {
  background-color: #eaeaea;
  margin-top: 50%;
}

.step-3 .rank {
  background-color: #eccec6;
  margin-top: 75%;
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
  content: " mm";
}
</style>
