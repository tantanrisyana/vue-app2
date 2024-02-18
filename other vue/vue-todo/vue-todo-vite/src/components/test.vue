<!-- src/App.vue -->
<template>
    <div id="app" class="container mt-5">
      <h1 class="mb-4">Vue To-Do List</h1>
      <button @click="showAddTaskModal" class="btn btn-primary mb-3">Tambah Tugas</button>
      
  
      <!-- Add Tugas Modal -->
      <div class="modal" :class="{ 'show': isAddTaskModalVisible }" tabindex="-1" role="dialog" style="display: none;">
        <div class="modal-dialog" role="document">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title">Tambah Tugas</h5>
              <button type="button" class="close" @click="hideAddTaskModal">
                <span>&times;</span>
              </button>
            </div>
            <div class="modal-body">
              <div class="mb-3">
                <label for="newTask" class="form-label">Tugas</label>
                <input v-model="newTask" class="form-control" id="newTask" placeholder="Tambah tugas...">
              </div>
              <div class="mb-3">
                <label for="newTime" class="form-label">Jam</label>
                <input v-model="newTime" class="form-control" id="newTime" placeholder="Jam...">
              </div>
            </div>
            <div class="modal-footer">
              <button @click="hideAddTaskModal" type="button" class="btn btn-secondary">Tutup</button>
              <button @click="addTask" type="button" class="btn btn-primary">Tambah Tugas</button>
            </div>
          </div>
        </div>
      </div>
  
      <!-- To-Do List Table -->
      <table class="table">
        <thead>
          <tr>
            <th scope="col">#</th>
            <th scope="col">Tugas</th>
            <th scope="col">Jam</th>
            <th scope="col">Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="task in tasks" :key="task.id">
            <th scope="row">{{ task.id }}</th>
            <td>{{ task.task }}</td>
            <td>{{ task.time }}</td>
            <td>
              <button @click="editTask(task.id)" class="btn btn-warning btn-sm">Edit</button>
              <button @click="deleteTask(task.id)" class="btn btn-danger btn-sm">Hapus</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        newTask: "",
        newTime: "",
        tasks: [],
        isAddTaskModalVisible: true,
      };
    },
    methods: {
      async fetchTasks() {
        const response = await fetch("http://localhost:8080/tasks");
        const data = await response.json();
        this.tasks = data;
      },
      addTask() {
        if (this.newTask.trim() !== "" && this.newTime.trim() !== "") {
          fetch("http://localhost:8080/tasks", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              task: this.newTask,
              time: this.newTime,
            }),
          })
          .then(response => {
            if (!response.ok) {
              throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.json();
          })
          .then(() => {
            this.fetchTasks();
            this.hideAddTaskModal();
            this.newTask = "";
            this.newTime = "";
          })
          .catch(error => {
            console.error("Error adding task:", error);
          });
        }
      },
      editTask(id) {
        const updatedTask = prompt("Edit tugas:");
        const updatedTime = prompt("Edit jam:");
        if (updatedTask !== null && updatedTime !== null) {
          fetch(`http://localhost:8080/tasks/${id}`, {
            method: "PUT",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              task: updatedTask,
              time: updatedTime,
            }),
          }).then(() => {
            this.fetchTasks();
          });
        }
      },
      deleteTask(id) {
        const confirmation = confirm("Apakah Anda yakin ingin menghapus tugas ini?");
        if (confirmation) {
          fetch(`http://localhost:8080/tasks/${id}`, {
            method: "DELETE",
          }).then(() => {
            this.fetchTasks();
          });
        }
      },
      showAddTaskModal() {
    console.log('Showing Add Tugas Modal');
    this.isAddTaskModalVisible = true;
  },
  
      hideAddTaskModal() {
        this.isAddTaskModalVisible = false;
      },
    },
    created() {
      this.fetchTasks();
    },
  };
  </script>
  
  <style>
  /* Your existing styles or custom styles go here */
  </style>
  