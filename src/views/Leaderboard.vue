<script lang="ts">
import axios from "axios";
import RankingList from "@/components/RankingList.vue";
import Podium from "@/components/Podium.vue";

type Competitor = {
  id: number;
  name: string;
  identifier: number;
  target: number;
  measurement: number;
  offset: number;
  rank: number;
};
export type { Competitor };

export default {
  components: { RankingList, Podium },

  data() {
    return {
      competitors: [] as Competitor[],
    };
  },

  methods: {
    async fetchCompetitors() {
      axios.get(`${import.meta.env.VITE_LEADERBOARD_API_URL}/api/competitor`).then((response) => {
        this.competitors = response.data.competitors;
      });
    },
  },

  async created() {
    await this.fetchCompetitors();
  },
};
</script>

<template>
  <main>
    <div class="section">
      <h2>Bestenliste ({{ competitors.length }})</h2>
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
  min-height: 100vh;
}

h2 {
  margin-bottom: 40px;
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
  padding: 48px 72px;
}

.section:nth-child(2) {
  border-left: 3px dashed #003399;
}

@media screen and (max-width: 1100px) {
  main {
    grid-template-columns: 1fr;
  }
  .section:nth-child(2) {
    border-left: none;
  }
  .section {
    padding: 5%;
  }

  h2::before {
    display: none;
  }

  h2 {
    margin-bottom: 40px;
  }
}
</style>
