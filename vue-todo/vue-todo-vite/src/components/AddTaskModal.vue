<!-- src/components/AddTaskModal.vue -->
<template>
    <!-- The modal container -->
    <div class="modal" :class="{ 'show': isAddTaskModalVisible }" tabindex="-1" role="dialog">
      <!-- Modal dialog -->
      <div class="modal-dialog" role="document">
        <!-- Modal content -->
        <div class="modal-content">
          <!-- Modal header -->
          <div class="modal-header">
            <h5 class="modal-title">Tambah Tugas</h5>
            <!-- Close button to hide the modal -->
            <button type="button" class="close" @click="hideAddTaskModal">
              <span>&times;</span>
            </button>
          </div>
          <!-- Modal body -->
          <div class="modal-body">
            <!-- Form fields for new task and time -->
            <div class="mb-3">
              <label for="newTask" class="form-label">Tugas</label>
              <input v-model="newTask" class="form-control" id="newTask" placeholder="Tambah tugas...">
            </div>
            <div class="mb-3">
              <label for="newTime" class="form-label">Jam</label>
              <input v-model="newTime" class="form-control" id="newTime" placeholder="Jam...">
            </div>
          </div>
          <!-- Modal footer with buttons -->
          <div class="modal-footer">
            <!-- Button to close the modal -->
            <button @click="hideAddTaskModal" type="button" class="btn btn-secondary">Tutup</button>
            <!-- Button to add a new task -->
            <button @click="addTask" type="button" class="btn btn-primary">Tambah Tugas</button>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    props: ['isAddTaskModalVisible'],
    data() {
      return {
        // Data properties for the form fields
        newTask: '',
        newTime: '',
      };
    },
    methods: {
      // Method to hide the modal
      hideAddTaskModal() {
        // Emit a custom event to notify the parent component to hide the modal
        this.$emit('hide');
      },
      // Method to add a new task
      addTask() {
        // Perform any necessary logic to add the new task
        // For example, emit an event to notify the parent component about the new task details
        this.$emit('add-task', {
          task: this.newTask,
          time: this.newTime,
        });
  
        // Clear the form fields
        this.newTask = '';
        this.newTime = '';
  
        // Optionally, hide the modal after adding the task
        this.hideAddTaskModal();
      },
      // ... (any other methods you need for the modal functionality)
    },
  };
  </script>
  
  <style scoped>
  /* Add any scoped styles for the modal component */
  </style>
  