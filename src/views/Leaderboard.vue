<script lang="ts">
import axios from "axios";
import RankingList from "@/components/RankingList.vue";

type Competitor = {
  name: string;
  faultPoints: number;
  offsetX: number;
  offsetY: number;
};
export type {Competitor};

export default {
  components: {RankingList},

  data() {
    return {
      competitors: [] as Competitor[],
    };
  },

  mounted() {
    axios.get("http://localhost:8008/api/all")
        .then((response) => {
          this.competitors = response.data.competitors;
        });
  },
};
</script>

<template>
  <main>
    <p>Linke Seite</p>
    <div>
      <h1>Letzten Versuche</h1>
      <RankingList :competitors="competitors"/>
    </div>
  </main>
</template>

<style scoped>
main {
  display: grid;
  grid-template-columns: 1fr 1fr;
}
</style>
