<script lang="ts">
import RankingList from "@/components/RankingList.vue";
import { ref, defineComponent, onMounted } from "vue";
import axios from "axios";

type Competitor = {
  name: string;
  faultPoints: number;
  offsetX: number;
  offsetY: number;
};
export type { Competitor };

export default defineComponent({
  setup() {
    const competitors = ref([]);
    onMounted(async () => {
      const response = await axios("http://localhost:8008/api/all");
      console.log(response);
    });
  },
  components: {
    RankingList,
  },
});
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
