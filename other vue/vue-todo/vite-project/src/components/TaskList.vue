<template>
    <div>
      <h1>Vue To-Do List</h1>
      <input v-model="newTask" placeholder="Tambah tugas...">
      <input v-model="newTime" placeholder="Jam...">
      <button @click="addTask">Tambah Tugas</button>
      <ul>
        <li v-for="task in tasks" :key="task.id">
          <span>{{ task.task }} - {{ task.time }}</span>
          <button @click="editTask(task.id)">Edit</button>
          <button @click="deleteTask(task.id)">Hapus</button>
        </li>
      </ul>
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
          }).then(() => {
            this.fetchTasks();
            this.newTask = "";
            this.newTime = "";
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
  