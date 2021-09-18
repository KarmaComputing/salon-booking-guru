import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';

// accounts
import Home from '../views/Home.vue';
import Accounts from '../views/Accounts.vue';

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'Home',
        component: Home,
    },
    {
        path: '/accounts',
        name: 'Accounts',
        component: Accounts,
    },
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes,
});

export default router;
