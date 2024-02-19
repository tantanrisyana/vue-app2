<template>
 <DxDataGrid
  :columns="columns"
  :show-borders="true"
  :data-source="dataSource"
/>
</template>

<script setup lang="ts">
import DxDataGrid from 'devextreme-vue/data-grid';

const columns = ['CompanyName', 'City', 'State', 'Phone', 'Fax'];
const dataSource = '/data/customers.json';
</script>

