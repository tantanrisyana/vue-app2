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
        <dx-data-grid :data-source="filteredGridData" key-expr="id" :column-auto-width="true">
          <!-- DataGrid columns -->
          <dxo-column data-field="Id" caption="Id"></dxo-column>
          <dxo-column data-field="Name" caption="Name"></dxo-column>
          <dxo-column data-field="Alamat" caption="Alamat"></dxo-column>
          
          <!-- Custom command column for Add, Edit, and Delete buttons -->
          <dxo-column type="buttons">
            <!-- Template for custom buttons in the column -->
            <template #cellTemplate="data">
              <!-- Edit button -->
    <button
      title="Ubah Data"
      class="me-1 btn-light-warning border border-warning btn btn-xs btn-icon"
      @click="editData(data.data.Id)"
    >
      Edit Data
    </button>
    <!-- Other buttons (Add, Delete) as needed -->
    <dx-button
      :show-delete="true"
      @on-delete="handleDeleteRow"
    ></dx-button>
            </template>
          </dxo-column>
        </dx-data-grid>
      </section>
    </div>
  </template>
  
  <script>
  // Importing necessary dependencies and styles
  
  // Importing DevExtreme components from the correct module
  import 'devextreme/dist/css/dx.light.css';
  import { DxDataGrid, DxoColumn, DxButtons, DxButton } from 'devextreme-vue/ui/data-grid';
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
        axios.get('http://localhost:8080/siswa')
          .then(response => {
            // Assigning raw data to the 'rawData' property
            this.rawData = response.data;
            // Processing raw data and assigning it to the 'filteredGridData' property
            this.filteredGridData = this.rawData.map(item => ({
              id: item.id,
              name: item.name,
              alamat: item.alamat,
            }));
          })
          .catch(error => {
            console.error('Error fetching data:', error);
          });
      },
      // Method to handle "Edit Data" button click
      editData(id) {
        // Redirect to the edit page or show a modal for editing
        // For now, let's log the ID and make a request to the edit endpoint
        console.log('Edit Data button clicked for ID:', id);
        axios.get(`http://localhost:8080/siswa/${id}`)
          .then(response => {
            // Handle the response, for example, show an edit form with the data
            console.log('Edit Data response:', response.data);
          })
          .catch(error => {
            console.error('Error fetching edit data:', error);
          });
      },
      // Method to handle "Delete" button click
      handleDeleteRow(e) {
        // Implement logic for row deletion
        // For example, you can make an API call to delete the selected row
        console.log('Delete row:', e.data);
      },
    },
    // Registering Vue components
    components: {
      DxDataGrid,
      DxoColumn,
      DxButtons,
      DxButton,
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
  
  /* Styling for the button section */
  .button-section {
    margin-bottom: 20px; /* Bottom margin for button section */
  }
  
  /* Styling for buttons */
  button {
    margin-right: 10px;  /* Right margin between buttons */
  }
  </style>
  