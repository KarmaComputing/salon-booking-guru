// vue
import { useStore } from 'vuex';
import { useRouter } from 'vue-router';

// hooks
import { useAxios } from '@/hooks/axiosHook';

// models
import { Credentials, AuthenticateResponse } from '@/api/models';

export const useAuthenticateService = () => {
    // hooks
    const store = useStore();
    const axios = useAxios();
    const router = useRouter();

    // properties
    const { apiUrl } = store.state;

    const authenticate = async (
        credentials: Credentials,
    ): Promise<AuthenticateResponse> => {
        const res = await axios.post(`${apiUrl}/authenticate`, credentials);
        store.commit('authToken', res.data.token);
        store.commit('accountInfo', res.data.accountInfo);
        router.push({ path: '/' });

        return res.data;
    };

    const logOut = async () => {
        store.commit('authToken');
        store.commit('accountInfo');
        router.push({ path: '/log-in' });
    };

    return {
        authenticate,
        logOut,
    };
};
