import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';

// accounts
import Home from '../views/Home.vue';
import Accounts from '../views/Accounts.vue';
import AccountEditor from '../views/AccountEditor.vue';
import AccountAvailability from '../views/AccountAvailability.vue';

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
    {
        path: '/account/editor',
        name: 'Account editor',
        component: AccountEditor,
    },
    {
        path: '/account/availability',
        name: 'Account availability',
        component: AccountAvailability,
    },
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes,
});

export default router;
