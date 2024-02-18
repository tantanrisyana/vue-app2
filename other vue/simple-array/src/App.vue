<template>
  <div id="app">
    <DxDataGrid
      :data-source="customers"
      key-expr="ID"
      :columns="columns"
      :show-borders="true"
    />
  </div>
</template>

<script setup lang="ts">
import DxDataGrid from 'devextreme-vue/data-grid';
import { customers } from './data.ts';

const columns = ['CompanyName', 'City', 'State', 'Phone', 'Fax'];
</script>

<style scoped>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
