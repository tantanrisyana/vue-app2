<template>
  <div id="app">
    <SiswaComponent />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import SiswaComponent from './components/SiswaComponent.vue';

export default defineComponent({
  components: {
    SiswaComponent,
  },
});
</script>
