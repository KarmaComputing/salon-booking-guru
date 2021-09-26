// vue
import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import store from '@/store';

// axios
import axios from 'axios';

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'Home',
        component: () =>
            import(/* webpackChunkName: "core" */ '@/views/Home.vue'),
    },
    {
        path: '/accounts',
        name: 'Accounts',
        component: () =>
            import(/* webpackChunkName: "core" */ '@/views/Accounts.vue'),
    },
    {
        path: '/account/availability',
        name: 'AccountAvailability',
        component: () =>
            import(
                /* webpackChunkName: "core" */ '@/views/AccountAvailability.vue'
            ),
    },
    {
        path: '/product-categories',
        name: 'ProductCategories',
        component: () =>
            import(
                /* webpackChunkName: "core" */ '@/views/ProductCategories.vue'
            ),
    },
    {
        path: '/qualifications',
        name: 'Qualifications',
        component: () =>
            import(/* webpackChunkName: "core" */ '@/views/Qualifications.vue'),
    },
    {
        path: '/log-in',
        name: 'LogIn',
        component: () =>
            import(/* webpackChunkName: "core" */ '@/views/LogIn.vue'),
    },
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes,
});

router.beforeEach((to, from, next) => {
    store.state.cancelToken.cancel('Operation cancelled by the user');

    const allowedRoutes = ['/log-in'];

    if (store.state.authToken) {
        if (allowedRoutes.indexOf(to.path) !== -1) {
            next('/');
        }
    } else {
        if (allowedRoutes.indexOf(to.path) === -1) {
            next('/log-in');
        }
    }

    store.state.cancelToken = axios.CancelToken.source();
    next();
});

export default router;
