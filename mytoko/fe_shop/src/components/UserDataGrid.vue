<template>
  <DxDataGrid
    :data-source="dataSource"
    key-expr="id"
    :show-borders="true"
    :column-auto-width="true"
    :width="'70%'"
    style="margin: 0 auto; text-align: center"
  >
    <DxColumn data-field="id" caption="ID" />
    <DxColumn data-field="name" caption="Name" />
    <DxColumn data-field="email" caption="Email" />

    <DxHeaderFilter :visible="true" />
    <DxGroupPanel :visible="true" />
    <DxScrolling mode="virtual" />
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

    <DxScrolling row-rendering-mode="virtual" />
    <DxPaging :page-size="5" />
    <DxPager
      :visible="true"
      :allowed-page-sizes="pageSizes"
      :display-mode="displayMode"
      :show-page-size-selector="showPageSizeSelector"
      :show-info="showInfo"
      :show-navigation-buttons="showNavButtons"
    />
  </DxDataGrid>
</template>

<script setup lang="ts">
import {
  DxDataGrid,
  DxColumn,
  DxEditing,
  DxFilterRow,
  DxHeaderFilter,
  DxGroupPanel,
  DxGrouping,
  DxScrolling,
  DxPaging,
  DxPager,
} from "devextreme-vue/data-grid";
import axios from "axios";
import "devextreme/dist/css/dx.material.blue.light.compact.css";

const url = "http://localhost:8080/users";

const dataSource = {
  key: "id",
  load: async () => {
    try {
      const response = await axios.get(url);
      return response.data;
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
</script>
