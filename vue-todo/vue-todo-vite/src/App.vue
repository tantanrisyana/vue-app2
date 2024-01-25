<!-- src/App.vue -->
<template>
  <div>
    <h1>Vue To-Do List</h1>
    <input v-model="newTask" placeholder="Tambah tugas...">
    <input v-model="newTime" placeholder="Jam...">
    <button @click="addTask" class="btn btn-primary">Tambah Tugas</button>
    
    <!-- Bootstrap Table -->
    <table class="table">
      <thead>
        <tr>
          <th scope="col">Tugas</th>
          <th scope="col">Jam</th>
          <th scope="col">Aksi</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="task in tasks" :key="task.id">
          <td>{{ task.task }}</td>
          <td>{{ task.time }}</td>
          <td>
            <button @click="editTask(task.id)" class="btn btn-warning">Edit</button>
            <button @click="deleteTask(task.id)" class="btn btn-danger">Hapus</button>
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
  },
  created() {
    this.fetchTasks();
  },
};
</script>

<style>
/* Add Bootstrap CSS link here */
@import url('https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css');

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
