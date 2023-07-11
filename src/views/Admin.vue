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
  methods: {
    async onSubmit() {
      try {
        await axios.post("http://localhost:8008/api/competitor", {
          method: "POST",
          headers: {
            Accept: "application/json",
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            name: this.name,
            offsetX: this.offsetX,
            offsetY: this.offsetY,
          }),
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
      <input type="text" name="name" v-model="name" />

      <label for="offsetX">Abweichung X</label>
      <input type="number" name="offsetX" v-model="offsetX" />

      <label for="offsetY">Abweichung Y</label>
      <input type="number" name="offsetY" v-model="offsetY" />
      <button type="submit">Absenden</button>
    </form>
  </main>
</template>
