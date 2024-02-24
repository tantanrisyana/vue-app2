import { createApp } from 'vue';
import DataGridComponent from './components/DataGridComponent.vue';

// Import DevExtreme components
import {
        DxDataGrid,
        DxoColumn,
        DxEditing,
        DxFilterRow,
        DxHeaderFilter,
        DxGroupPanel,
        DxScrolling,
    
} from 'devextreme-vue/data-grid';
import {DxButton} from 'devextreme-vue/button';

// Buat aplikasi Vue
const app = createApp(DataGridComponent);

// Daftarkan DevExtreme components sebagai global components
app.component(
        DxDataGrid,
        DxoColumn,
        DxEditing,
        DxFilterRow,
        DxHeaderFilter,
        DxGroupPanel,
        DxScrolling,
        DxButton,
    );


// Pasang aplikasi ke elemen dengan ID 'app'
app.mount('#app');
