<script setup lang="ts">
import type { Competitor } from "@/views/Leaderboard.vue";

defineProps<{
  competitor: Competitor;
}>();
</script>

<template>
  <div>
    <p>{{ competitor.name }}</p>
    <p>{{ competitor.faultPoints }}</p>
  </div>
</template>

<style scoped></style>
