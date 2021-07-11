// vue
import { createApp } from 'vue';
import App from './App.vue';
import './registerServiceWorker';
import router from './router';
import store from './store';

// primevue
import 'primeicons/primeicons.css';
import 'primevue/resources/primevue.min.css';
import 'primevue/resources/themes/md-light-indigo/theme.css';
import PrimeVue from 'primevue/config';

// primeflex
import 'primeflex/primeflex.css';

// primevue components
import Button from 'primevue/button';

// create app
const app = createApp(App);

// uses
app.use(store);
app.use(router);
app.use(PrimeVue, { ripple: true });

// components
app.component('Button', Button);

// mount app
app.mount('#app');
