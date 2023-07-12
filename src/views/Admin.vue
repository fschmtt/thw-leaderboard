<script lang="ts">
import axios from "axios";

export default {
  data() {
    return {
      name: null,
      offsetX: null,
      offsetY: null,
    };
  },

  computed: {
    isButtonDisabled(): boolean {
      return !this.name || !this.offsetX || !this.offsetY;
    },
  },

  methods: {
    async onSubmit() {
      try {
        await axios.post(`${import.meta.env.VITE_LEADERBOARD_API_URL}/api/competitor`, {
          name: this.name,
          offsetX: this.offsetX,
          offsetY: this.offsetY,
        });

        this.name = null;
        this.offsetX = null;
        this.offsetY = null;
      } catch (error) {
        console.error(error);
      }
    },
  },
};
</script>

<template>
  <main>
    <form @submit.prevent="onSubmit">
      <label for="name">Name</label>
      <input type="text" name="name" id="name" v-model="name" />

      <label for="offsetX">Abweichung X</label>
      <input type="number" name="offsetX" id="offsetX" v-model="offsetX" />

      <label for="offsetY">Abweichung Y</label>
      <input type="number" name="offsetY" id="offsetY" v-model="offsetY" />

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
}

button[disabled] {
  background-color: #e6e6e6;
  color: #999999;
}
</style>