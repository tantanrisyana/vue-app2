<template>
    <div>
      <h2>Pokemon List</h2>
      <DxDataGrid
        :data-source="currentPageData"
        :columns="[
          { dataField: 'name', caption: 'Name' },
          { dataField: 'sprites.front_default', caption: 'Image', cellTemplate: 'imageTemplate' }
        ]"
        :show-pager="false"
      >
        <template #imageTemplate="data">
          <img v-if="data.data.sprites && data.data.sprites.front_default" :src="data.data.sprites.front_default" alt="Pokemon Image" />
          <span v-else>No image available</span>
        </template>
      </DxDataGrid>
  
      <nav aria-label="Page navigation">
        <ul class="pagination">
          <li class="page-item" :class="{ disabled: currentPage === 1 }">
            <a class="page-link" href="#" @click="changePage(currentPage - 1)" aria-label="Previous">
              <span aria-hidden="true">&laquo;</span>
            </a>
          </li>
          <li class="page-item" v-for="page in totalPages" :key="page" :class="{ active: page === currentPage }">
            <a class="page-link" href="#" @click="changePage(page)">{{ page }}</a>
          </li>
          <li class="page-item" :class="{ disabled: currentPage === totalPages }">
            <a class="page-link" href="#" @click="changePage(currentPage + 1)" aria-label="Next">
              <span aria-hidden="true">&raquo;</span>
            </a>
          </li>
        </ul>
      </nav>
    </div>
  </template>
  
  <script>
  import 'devextreme/dist/css/dx.light.css';
  import { DxDataGrid } from 'devextreme-vue';
  
  export default {
    components: {
      DxDataGrid,
    },
    props: {
      pokemonList: Array,
    },
    data() {
      return {
        currentPage: 1,
        pageSize: 10,
      };
    },
    computed: {
      totalPages() {
        return Math.ceil(this.pokemonList.length / this.pageSize);
      },
      currentPageData() {
        const start = (this.currentPage - 1) * this.pageSize;
        const end = start + this.pageSize;
        return this.pokemonList.slice(start, end);
      },
    },
    methods: {
      changePage(page) {
        if (page >= 1 && page <= this.totalPages) {
          this.currentPage = page;
        }
      },
    },
  };
  </script>
  
  <style scoped>
  /* Add any necessary styling here */
  </style>
  