import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import 'bootstrap/dist/css/bootstrap.min.css'
import PageCard from './components/Page/PageCard.vue'
import PageButton from './components/Page/PageButton.vue'

const app = createApp(App);
app.component('page-card', PageCard);
app.component('page-button', PageButton);

app.mount('#app')
