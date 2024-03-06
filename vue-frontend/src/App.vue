<template>
  <div><h2>DATA SISWA TAHUN AJARAN 2024/2025</h2></div>
  <DxDataGrid
    :data-source="dataSource"
    key-expr="id"
    :show-borders="true"
    :column-auto-width="true"
    :width="'70%'"
    style="margin: 0 auto; text-align: center"
    @row-prepared="onRowPrepared"
  >
    <DxColumn data-field="no" caption="No" />
    <DxColumn data-field="id" caption="ID" :visible="false" />
    <DxColumn data-field="name" caption="Nama" />
    <DxColumn data-field="alamat" caption="Alamat" />
    <DxColumn data-field="email" caption="Email" />
    <DxEditing
      :allow-adding="true"
      :allow-updating="true"
      :allow-deleting="true"
      mode="row"
      :use-icons="true"
      @onInitNewRow="handleInitNewRow"
      @onRowInserting="showAddModal"
      @onRowUpdating="handleRowUpdating"
      @onRowRemoving="confirmDelete"
    />

    <DxGroupPanel :visible="true" />
    <DxSearchPanel :visible="true" :highlight-case-sensitive="true" />
    <DxGrouping :auto-expand-all="false" />
    <DxPager :allowed-page-sizes="pageSizes" :show-page-size-selector="true" />
    <DxPaging :page-size="10" />
  </DxDataGrid>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import axios from "axios";
import {
  DxDataGrid,
  DxColumn,
  DxGrouping,
  DxGroupPanel,
  DxPager,
  DxPaging,
  DxSearchPanel,
  DxDataGridTypes,
  DxEditing,
} from "devextreme-vue/data-grid";
const url = "http://localhost:8080/siswa";

const dataSource = {
  key: "id",
  load: async () => {
    try {
      const response = await axios.get(url);
      const dataWithNo = response.data.map((item, index) => ({
        ...item,
        no: index + 1,
      }));
      return dataWithNo;
    } catch (error) {
      console.error("Error loading data:", error);
      throw error;
    }
  },
  insert: async (values) => {
    try {
      await axios.post(url, values);
    } catch (error) {
      console.error("Error inserting data:", error);
      throw error;
    }
  },
  update: async (key, values) => {
    try {
      await axios.put(`${url}/${key}`, values);
    } catch (error) {
      console.error("Error updating data:", error);
      throw error;
    }
  },
  remove: async (key) => {
    try {
      await axios.delete(`${url}/${key}`);
    } catch (error) {
      console.error("Error deleting data:", error);
      throw error;
    }
  },
};

const pageSizes = [3, 5, 7, 12];

let collapsed = false;

const onContentReady = (e: DxDataGridTypes.ContentReadyEvent) => {
  if (!collapsed) {
    e.component.expandRow([dataSource.value[0].id]); // Adjust based on your data
    collapsed = true;
  }
};

onMounted(async () => {
  try {
    const data = await dataSource.value.load();
    dataSource.value = data;
  } catch (error) {
    console.error("Error fetching data:", error);
  }
});
</script>
