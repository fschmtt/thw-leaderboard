<script lang="ts">
export default {
  data() {
    return {
      name: null,
      offsetX: null,
      offsetY: null,
    };
  },
  methods: {
    onSubmit() {
      console.log(this.name, this.offsetX, this.offsetY);
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
