<script setup lang="ts">
type Competitor = {
  name: string;
  score: number;
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
      <p>Rechte Seite</p>
    </div>
  </main>
</template>

<style scoped>
main {
  display: grid;
  grid-template-columns: 1fr 1fr;
}
</style>
