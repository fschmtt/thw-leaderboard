<script setup lang="ts">
import RankingList from "@/components/RankingList.vue";
const competitors = [
  {
    name: "Max",
    faultPoints: 100,
    offsetX: 0,
    offsetY: 0,
  },
  {
    name: "Moritz",
    faultPoints: 200,
    offsetX: 0,
    offsetY: 0,
  },
  {
    name: "Peter",
    faultPoints: 300,
    offsetX: 0,
    offsetY: 0,
  },
  {
    name: "Hans",
    faultPoints: 400,
    offsetX: 0,
    offsetY: 0,
  },
  {
    name: "Klaus",
    faultPoints: 500,
    offsetX: 0,
    offsetY: 0,
  },
];

type Competitor = {
  name: string;
  faultPoints: number;
  offsetX: number;
  offsetY: number;
};
export type { Competitor };
</script>

<template>
  <main>
    <p>Linke Seite</p>
    <div>
      <h1>Letzten Versuche</h1>
      <RankingList :competitors="competitors" />
    </div>
  </main>
</template>

<style scoped>
main {
  display: grid;
  grid-template-columns: 1fr 1fr;
}
</style>
