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
        <dxo-column data-field="no" caption="No"></dxo-column>
        <dxo-column v-if="false" data-field="id" caption="Id"></dxo-column>
        <dxo-column data-field="name" caption="Name"></dxo-column>
        <dxo-column data-field="alamat" caption="Alamat"></dxo-column>
        <dxo-column data-field="email" caption="Email"></dxo-column>
      </dx-data-grid>
    </section>

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
                      <button type="button" class="button is-danger" @click="cancelAdd">Cancel</button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </form>
        </section>
      </div>
    </div>

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
                <button type="button" class="button is-danger" @click="cancelDelete">Cancel</button>
              </div>
            </div>
          </form>
        </section>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import axios from "axios";
import { DxDataGrid, DxoColumn } from "devextreme-vue/ui/data-grid";
import "@/assets/styles.css"; // Import the styles

export default defineComponent({
  name: "App",
  setup() {
    // Inline definition of the Siswa type
    interface Siswa {
      id: number;
      name: string;
      alamat: string;
      email: string;
    }

    const rawData = ref<Siswa[]>([]);
    const filteredGridData = ref<Array<{ no: number; id: number; name: string; alamat: string; email: string }>>([]);
    const isAddModalVisible = ref(false);
    const isDeleteModalVisible = ref(false);
    const newData = ref<Siswa>({
      name: "",
      alamat: "",
      email: "",
      id: 0,
    });
    const deleteDataId = ref<number | string>("");

    const fetchData = async () => {
      try {
        const response = await axios.get<Siswa[]>("http://localhost:8080/siswa");
        rawData.value = response.data;
        filteredGridData.value = rawData.value.map((item, index) => ({
          no: index + 1,
          id: item.id,
          name: item.name,
          alamat: item.alamat,
          email: item.email,
        }));
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };

    const showAddModal = () => {
      isAddModalVisible.value = true;
    };

    const hideAddModal = () => {
      isAddModalVisible.value = false;
      newData.value = {
        name: "",
        alamat: "",
        email: "",
        id: 0,
      };
    };

    const addData = () => {
      axios
        .post("http://localhost:8080/siswa", newData.value)
        .then((response) => {
          console.log("Data added successfully:", response.data);
          fetchData();
          hideAddModal();
        })
        .catch((error) => {
          console.error("Error adding data:", error);
        });
    };

    const cancelAdd = () => {
      hideAddModal();
    };


    const showDeleteModal = () => {
      isDeleteModalVisible.value = true;
    };

    const cancelDelete = () => {
      hideDeleteModal();
    };

    const hideDeleteModal = () => {
      isDeleteModalVisible.value = false;
      deleteDataId.value = "";
    };

    const confirmDelete = () => {
      if (confirm("Are you sure you want to delete this data?")) {
        deleteData();
      }
    };

    const deleteData = () => {
      axios
        .delete(`http://localhost:8080/siswa/${deleteDataId.value}`)
        .then(() => {
          console.log("Data deleted successfully");
          fetchData();
          hideDeleteModal();
        })
        .catch((error) => {
          console.error("Error deleting data:", error);
        });
    };

    fetchData();

    return {
      rawData,
      filteredGridData,
      isAddModalVisible,
      isDeleteModalVisible,
      newData,
      deleteDataId,
      fetchData,
      showAddModal,
      hideAddModal,
      addData,
      showDeleteModal,
      hideDeleteModal,
      confirmDelete,
      deleteData,
      cancelAdd,
      cancelDelete,
    };
  },
  components: {
    DxDataGrid,
    DxoColumn,
  },
});

</script>

