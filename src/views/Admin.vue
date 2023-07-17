<script lang="ts">
import axios, {AxiosError} from 'axios';
import type {Competitor} from "@/views/Leaderboard.vue";
import RankingListItem from "@/components/RankingListItem.vue";
import RankingList from "@/components/RankingList.vue";

type ComponentData = {
  name: string | null;
  identifier: number | null;
  offsetX: number | null;
  offsetY: number | null;
  error: AxiosError | null;
  competitors: Array<Competitor>;
};

export default {
  components: {RankingList, RankingListItem},
  data(): ComponentData {
    return {
      name: null,
      identifier: null,
      offsetX: null,
      offsetY: null,
      error: null,
      competitors: [],
    };
  },

  computed: {
    isButtonDisabled(): boolean {
      if (this.identifier === null || this.identifier < 1000) {
        return true;
      }

      if (this.offsetX === null || this.offsetX < 0) {
        return true;
      }

      if (this.offsetY === null || this.offsetY < 0) {
        return true;
      }

      return false;
    },

    errorMessage() {
      // @ts-expect-error
      return this.error?.response?.data?.error;
    },

    latestCompetitors(): Competitor[] {
      return [...this.competitors].sort((a, b) => b.id - a.id);
    },
  },

  methods: {
    async addCompetitor() {
      try {
        await axios.post(`${import.meta.env.VITE_LEADERBOARD_API_URL}/api/competitor`, {
          name: this.name,
          identifier: this.identifier,
          offsetX: this.offsetX,
          offsetY: this.offsetY,
        });

        this.name = null;
        this.identifier = null;
        this.offsetX = null;
        this.offsetY = null;
        this.error = null;
      } catch (error) {
        this.error = error as AxiosError;
      } finally {
        await this.fetchCompetitors();
      }
    },

    async fetchCompetitors() {
      axios.get(`${import.meta.env.VITE_LEADERBOARD_API_URL}/api/competitor`).then((response) => {
        this.competitors = response.data.competitors;
      });
    },
  },

  created() {
    this.fetchCompetitors();
  }
};
</script>

<template>
  <main>
    <!-- @ts-ignore -->
    <div class="error" v-if="error"><strong>{{ error }}:</strong> {{ errorMessage }}</div>

    <form @submit.prevent="addCompetitor">
      <label for="name">Name (optional)</label>
      <input type="text" name="name" id="name" autocomplete="off" v-model="name"/>

      <label for="identifier">Spielernummer</label>
      <input type="number" name="identifier" id="identifier" min="1000" step="1" autocomplete="false"
             v-model="identifier"/>

      <label for="offsetX">Abweichung X [mm]</label>
      <input type="number" name="offsetX" id="offsetX" min="0" step="1" autocomplete="false" v-model="offsetX"/>

      <label for="offsetY">Abweichung Y [mm]</label>
      <input type="number" name="offsetY" id="offsetY" min="0" step="1" autocomplete="false" v-model="offsetY"/>

      <button type="submit" :disabled="isButtonDisabled">Hinzufügen</button>
    </form>

    <hr>

    <h2>Alle Spieler ({{ competitors.length }})</h2>

    <RankingList :competitors="latestCompetitors"/>
  </main>
</template>

<style scoped>
main {
  margin: 20px;
}

label {
  font-weight: bold;
  font-size: 16px;
}

input {
  display: block;
  width: 100%;
  margin-bottom: 16px;
  padding: 16px 8px;
  font-size: 16px;
}

button {
  display: block;
  width: 100%;
  padding: 16px 8px;
  font-size: 16px;
  background-color: #003399;
  color: white;
  font-weight: bold;
  border: none;
  cursor: pointer;
  margin-bottom: 16px;
}

button[disabled] {
  background-color: #e6e6e6;
  color: #999999;
}

.error {
  color: red;
  border: 1px solid red;
  background-color: #ffe6e6;
  padding: 8px;
  margin-bottom: 16px;
}

h2 {
  margin: 16px 0;
}
</style>