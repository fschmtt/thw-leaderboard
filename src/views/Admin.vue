<script lang="ts">
import axios, {AxiosError} from 'axios';

type ComponentData = {
  name: string | null;
  identifier: number | null;
  offsetX: number | null;
  offsetY: number | null;
  error: AxiosError | null;
};

export default {
  data(): ComponentData {
    return {
      name: null,
      identifier: null,
      offsetX: null,
      offsetY: null,
      error: null,
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
    }
  },

  methods: {
    async onSubmit() {
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
      }
    },
  },
};
</script>

<template>
  <main>
    <!-- @ts-ignore -->
    <div class="error" v-if="error"><strong>{{ error }}:</strong> {{ errorMessage }}</div>

    <form @submit.prevent="onSubmit">
      <label for="name">Name (optional)</label>
      <input type="text" name="name" id="name" autocomplete="off" v-model="name" />

      <label for="identifier">Spielernummer</label>
      <input type="number" name="identifier" id="identifier" min="1000" step="1" autocomplete="false" v-model="identifier" />

      <label for="offsetX">Abweichung X [mm]</label>
      <input type="number" name="offsetX" id="offsetX" min="0" step="1" autocomplete="false" v-model="offsetX" />

      <label for="offsetY">Abweichung Y [mm]</label>
      <input type="number" name="offsetY" id="offsetY" min="0" step="1" autocomplete="false" v-model="offsetY" />

      <button type="submit" :disabled="isButtonDisabled">Hinzufügen</button>
    </form>
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
</style>