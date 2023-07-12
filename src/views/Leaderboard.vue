<script lang="ts">
import axios from "axios";
import RankingList from "@/components/RankingList.vue";
import Podium from "@/components/Podium.vue";

type Competitor = {
  id: number;
  name: string;
  score: number;
  offsetX: number;
  offsetY: number;
  rank: number;
};
export type { Competitor };

export default {
  components: { RankingList, Podium },

  data() {
    return {
      competitors: [] as Competitor[],
      countdown: 0,
      scaleX: 0,
    };
  },

  computed: {
    top3Competitors(): Competitor[] {
      return [...this.competitors].slice(0, 3);
    },

    topCompetitors(): Competitor[] {
      return [...this.competitors].slice(3, 10);
    },

    latest10Competitors(): Competitor[] {
      return [...this.competitors].sort((a, b) => b.id - a.id).slice(0, 10);
    },

    loaderStyles() {
      return [
        `transform: scaleX(${this.scaleX})`,
        this.countdown == 0 ? "transition: none" : "",
      ];
    },
  },

  methods: {
    async fetchCompetitors() {
      axios.get("http://localhost:8008/api/competitor").then((response) => {
        this.competitors = response.data.competitors;
      });
    },

    timer() {
      this.scaleX = this.countdown * 0.1;
      this.countdown++;
      if (this.countdown == 10) {
        this.fetchCompetitors();
        setTimeout(() => {
          this.countdown = 0;
        }, 1000);
      }
    },

    initCountdown() {
      setInterval(() => this.timer(), 1000);
    },
  },

  created() {
    this.fetchCompetitors();
    this.initCountdown();
  },
};
</script>

<template>
  <div class="loader" :style="loaderStyles"></div>
  <main>
    <div class="section">
      <h2>Bestenliste</h2>
      <Podium :competitors="top3Competitors" />
      <RankingList :competitors="topCompetitors" />
    </div>
    <div class="section">
      <h2>Letzte Messungen</h2>
      <RankingList :competitors="latest10Competitors" />
    </div>
  </main>
</template>

<style scoped>
.loader {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 5px;
  background-color: #003399;
  z-index: 999;
  transform-origin: left;
  transition: transform 1s linear;
}

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
