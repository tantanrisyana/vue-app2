<template>
    <div>
      <div>
        <h2>Daftar Siswa</h2>
        <dx-list :items="siswaList" displayExpr="nama" />
      </div>
      <div>
        <h2>Tambah Siswa Baru</h2>
        <dx-text-box v-model="newSiswa.nama" :placeholder="'Nama'" />
        <dx-number-box v-model="newSiswa.umur" :placeholder="'Umur'" />
        <dx-text-box v-model="newSiswa.alamat" :placeholder="'Alamat'" />
        <dx-button text="Simpan" @click="tambahSiswa" />
      </div>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent } from 'vue';
  import { DxList, DxTextBox, DxNumberBox, DxButton } from 'devextreme-vue';
  import axios from 'axios';
  
  export default defineComponent({
    components: {
      DxList,
      DxTextBox,
      DxNumberBox,
      DxButton,
    },
    data() {
      return {
        siswaList: [] as any[],
        newSiswa: {
          nama: '',
          umur: null,
          alamat: '',
        },
      };
    },
    mounted() {
      this.loadSiswa();
    },
    methods: {
      async loadSiswa() {
        try {
          const response = await axios.get('/api/v1/siswa');
          this.siswaList = response.data;
        } catch (error) {
          console.error('Error loading siswa:', error);
        }
      },
      async tambahSiswa() {
        try {
          await axios.post('/api/v1/siswa', this.newSiswa);
          this.loadSiswa();
          this.newSiswa = { nama: '', umur: null, alamat: '' };
        } catch (error) {
          console.error('Error adding siswa:', error);
        }
      },
    },
  });
  </script>
  