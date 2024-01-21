<template>
  <div id="app">
    <header>
      <h1>DATA SISWA</h1>
    </header>
    <br />

    <section class="grid-section">
      <div class="mb-3 d-flex justify-content-between align-items-center">
        <button class="btn btn-success" @click="showAddModal">Add Data</button>
        <button class="btn btn-success" @click="showDeleteModal">Delete Data</button>
      </div>

      <dx-data-grid :data-source="filteredGridData" key-expr="id" :column-auto-width="true">
        <!-- Column for sequential number (No) -->
        <dxo-column data-field="no" caption="No"></dxo-column>
        <!-- Hidden column for 'id' -->
        <dxo-column v-if="false" data-field="id" caption="Id"></dxo-column>
        <!-- Columns for visible data -->
        <dxo-column data-field="name" caption="Name"></dxo-column>
        <dxo-column data-field="alamat" caption="Alamat"></dxo-column>
        <dxo-column data-field="email" caption="Email"></dxo-column>
      </dx-data-grid>
    </section>

    <!-- Add Data Modal -->
    <div class="modal" :class="{ 'is-active': isAddModalVisible }">
      <div class="modal-background" @click="hideAddModal"></div>
      <div class="modal-card">
        <header class="modal-card-head">
          <p class="modal-card-title">Tambah Data</p>
          <button class="delete" aria-label="close" @click="hideAddModal"></button>
        </header>
        <section class="modal-card-body">
          <form @submit.prevent="addData">
            <table class="table table-borderless">
              <tbody>
                <tr>
                  <td class="label">Name</td>
                  <td>
                    <div class="control">
                      <input class="input" v-model="newData.name" required />
                    </div>
                  </td>
                </tr>
                <tr>
                  <td class="label">Alamat</td>
                  <td>
                    <div class="control">
                      <input class="input" v-model="newData.alamat" required />
                    </div>
                  </td>
                </tr>
                <tr>
                  <td class="label">Email</td>
                  <td>
                    <div class="control">
                      <input class="input" v-model="newData.email" required />
                    </div>
                  </td>
                </tr>
                <tr>
                  <td colspan="2">
                    <div class="control">
                      <button type="submit" class="button is-primary">Submit</button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </form>
        </section>
      </div>
    </div>

    <!-- Delete Data Modal -->
    <div class="modal" :class="{ 'is-active': isDeleteModalVisible }">
      <div class="modal-background" @click="hideDeleteModal"></div>
      <div class="modal-card">
        <header class="modal-card-head">
          <p class="modal-card-title">Delete Data</p>
          <button class="delete" aria-label="close" @click="hideDeleteModal"></button>
        </header>
        <section class="modal-card-body">
          <form @submit.prevent="confirmDelete">
            <div class="field">
              <label class="label">ID</label>
              <div class="control">
                <input class="input" v-model="deleteDataId" required />
              </div>
            </div>
            <div class="field">
              <div class="control">
                <button type="button" class="button is-danger" @click="confirmDelete">Delete</button>
              </div>
            </div>
          </form>
        </section>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import { DxDataGrid, DxoColumn } from "devextreme-vue/ui/data-grid";

export default {
  name: "App",
  data() {
    return {
      rawData: [],
      filteredGridData: [],
      isAddModalVisible: false,
      isDeleteModalVisible: false,
      newData: {
        name: "",
        alamat: "",
        email: "",
      },
      deleteDataId: "", // ID for the delete operation
    };
  },
  mounted() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      axios
        .get("http://localhost:8080/siswa")
        .then((response) => {
          this.rawData = response.data;
          this.filteredGridData = this.rawData.map((item, index) => ({
            no: index + 1,
            id: item.id,
            name: item.name,
            alamat: item.alamat,
            email: item.email,
          }));
        })
        .catch((error) => {
          console.error("Error fetching data:", error);
        });
    },
    showAddModal() {
      this.isAddModalVisible = true;
    },
    hideAddModal() {
      this.isAddModalVisible = false;
      this.newData = {
        name: "",
        alamat: "",
        email: "",
      };
    },
    addData() {
      axios
        .post("http://localhost:8080/siswa", this.newData)
        .then((response) => {
          console.log("Data added successfully:", response.data);
          this.fetchData();
          this.hideAddModal();
        })
        .catch((error) => {
          console.error("Error adding data:", error);
        });
    },
    editData(id) {
      console.log("Edit Data button clicked for ID:", id);
      axios
        .get(`http://localhost:8080/siswa/${id}`)
        .then((response) => {
          console.log("Edit Data response:", response.data);
        })
        .catch((error) => {
          console.error("Error fetching edit data:", error);
        });
    },
    showDeleteModal() {
      this.isDeleteModalVisible = true;
    },
    hideDeleteModal() {
      this.isDeleteModalVisible = false;
      this.deleteDataId = "";
    },
    confirmDelete() {
      // Add confirmation logic here
      if (confirm("Are you sure you want to delete this data?")) {
        this.deleteData();
      }
    },
    deleteData() {
      axios
        .delete(`http://localhost:8080/siswa/${this.deleteDataId}`)
        .then(() => {
          console.log("Data deleted successfully");
          this.fetchData();
          this.hideDeleteModal();
        })
        .catch((error) => {
          console.error("Error deleting data:", error);
        });
    },
  },
  components: {
    DxDataGrid,
    DxoColumn,
  },
};
</script>

<style>
#app {
  text-align: center;
  color: #2c3e50;
  margin-top: 20px;
}

header {
  background-color: #3498db;
  color: white;
  padding: 10px;
}

.grid-section {
  margin: 20px;
}

.modal {
  display: none;
}

.modal.is-active {
  display: flex;
  align-items: center;
  justify-content: center;
}

.modal-card {
  width: 400px;
}

.dx-datagrid {
  border: 1px solid #ccc;
}

.dx-datagrid-rowsview {
  border-top: 1px solid #ccc;
  border-bottom: 1px solid #ccc;
}

.dx-datagrid-content {
  border-bottom: 1px solid #ccc;
}

.dx-datagrid-rowsview .dx-row {
  border-bottom: 1px solid #ccc;
}
</style>
