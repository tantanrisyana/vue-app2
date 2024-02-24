import { createApp } from 'vue';
import App from './App.vue';
import axios from 'axios';

const app = createApp(App);

// Menambahkan global property $axios ke aplikasi Vue
app.config.globalProperties.$axios = axios;

app.mount('#app');
