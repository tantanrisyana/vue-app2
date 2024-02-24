<template>
  <div class="grid-container">
    <div class="data-grid-container">
      <h1>User</h1>
      <UserDataGrid />
    </div>
    <div class="data-grid-container">
      <h1>Product</h1>
      <ProductDataGrid />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import UserDataGrid from './components/UserDataGrid.vue';
import ProductDataGrid from './components/ProductDataGrid.vue';

export default defineComponent({
  components: {
    UserDataGrid,
    ProductDataGrid,
  },
});
</script>

<style scoped>
.grid-container {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 20px;
}

.data-grid-container {
  border: 1px solid #ddd; /* Optional: Add border for each DataGrid container */
  padding: 10px;
}
</style>
