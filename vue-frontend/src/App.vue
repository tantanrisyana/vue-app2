<template>
  <!-- Root element of the Vue app -->
  <div id="app">
    <!-- Header section -->
    <header>
      <!-- Main title -->
      <h1>Data Grid and Tree</h1>
    </header>
    <br/> <!-- Line break for separation -->
    <!-- Data Grid section -->
    <section class="grid-section">
      <!-- Subtitle for the Data Grid section -->
      <h2>Data Grid</h2>
      <!-- DevExtreme DataGrid component -->
      <dx-data-grid :data-source="filteredGridData" key-expr="ID" :column-auto-width="true">
        <!-- DataGrid columns -->
        <dxo-column data-field="ID" caption="ID"></dxo-column>
        <dxo-column data-field="Name" caption="Name"></dxo-column>
        <dxo-column data-field="CreatedAt" caption="Created At"></dxo-column>
      </dx-data-grid>
    </section>
  </div>
</template>

<script>
// Importing necessary dependencies and styles
import 'devextreme/dist/css/dx.light.css';
import DataGrid, { Column as DataGridColumn } from 'devextreme-vue/data-grid';
import axios from 'axios';

// Importing custom CSS file
import '@/assets/datagrid.css';

export default {
  // Component name
  name: 'App',
  // Data properties
  data() {
    return {
      rawData: [],           // Raw data from the API
      filteredGridData: [],  // Processed data for the DataGrid
    };
  },
  // Lifecycle hook - executed when the component is mounted
  mounted() {
    // Fetching data from the API when the component is mounted
    this.fetchData();
  },
  // Methods section
  methods: {
    // Method to fetch data from the API
    fetchData() {
      axios.get('http://localhost:3000/api/data')
        .then(response => {
          // Assigning raw data to the 'rawData' property
          this.rawData = response.data;
          // Processing raw data and assigning it to the 'filteredGridData' property
          this.filteredGridData = this.rawData.map(item => ({
            ID: item.ID,
            Name: item.Name,
            CreatedAt: item.CreatedAt,
          }));
        })
        .catch(error => {
          console.error('Error fetching data:', error);
        });
    },
  },
  // Registering Vue components
  components: {
    DxDataGrid: DataGrid,
    DxoColumn: DataGridColumn,
  },
};
</script>

<style>
/* Styling for the root element of the Vue app */
#app {
  text-align: center;   /* Center align text */
  color: #2c3e50;       /* Text color */
  margin-top: 20px;     /* Top margin */
}

/* Styling for the header section */
header {
  background-color: #3498db;  /* Background color */
  color: white;               /* Text color */
  padding: 10px;              /* Padding around the content */
}

/* Styling for h1 and h2 elements */
h1, h2 {
  margin-bottom: 10px;  /* Bottom margin for h1 and h2 elements */
}
</style>
