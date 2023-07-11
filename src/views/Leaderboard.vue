<script lang="ts">
import axios from "axios";
import RankingList from "@/components/RankingList.vue";
import Podium from "@/components/Podium.vue";

type Competitor = {
  name: string;
  faultPoints: number;
  offsetX: number;
  offsetY: number;
};
export type { Competitor };

export default {
  components: { RankingList, Podium },

  data() {
    return {
      competitors: [] as Competitor[],
    };
  },

  mounted() {
    axios.get("http://localhost:8008/api/all").then((response) => {
      this.competitors = response.data.competitors;
    });
  },
};
</script>

<template>
  <main>
    <div class="section">
      <h2>Bestenliste</h2>
      <Podium :competitors="competitors" />
      <RankingList :competitors="competitors" />
    </div>
    <div class="section">
      <h2>Letzte Versuche</h2>
      <RankingList :competitors="competitors" />
    </div>
  </main>
</template>

<style scoped>
main {
  display: grid;
  grid-template-columns: 1fr 1fr;
  font-family: Arial, Helvetica, sans-serif;
  background-color: #eff1f5;
}

h2 {
  margin-bottom: 96px;
  font-size: 56px;
  line-height: 88px;
  color: #003399;
  position: relative;
  font-weight: bold;
}

h2::before {
  content: "";
  background-color: #fcec4f;
  width: 48px;
  height: 88px;
  display: block;
  position: absolute;
  left: -72px;
}
.section {
  padding: 72px;
}

.section:nth-child(2) {
  border-left: 3px dashed #003399;
}
</style>
