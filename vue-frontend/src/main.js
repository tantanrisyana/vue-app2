import { createApp } from 'vue';
import DataGridComponent from './components/DataGridComponent.vue';

// Import DevExtreme components
import { DxDataGrid, DxColumn, DxPaging, DxHeaderFilter } from 'devextreme-vue/data-grid';

// Buat aplikasi Vue
const app = createApp(DataGridComponent);

// Daftarkan DevExtreme components sebagai global components
app.component('dx-data-grid', DxDataGrid);
app.component('dxc-column', DxColumn);
app.component('dxo-paging', DxPaging);
app.component('dxo-header-filter', DxHeaderFilter);

// Pasang aplikasi ke elemen dengan ID 'app'
app.mount('#app');
