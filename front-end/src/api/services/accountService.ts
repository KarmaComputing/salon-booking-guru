// vue
import { useStore } from 'vuex';

// hooks
import { useAxios } from '@/hooks/axiosHook';

// models
import { Account } from '@/api/models';

export const useAccountService = () => {
    const { apiUrl } = useStore().state;
    const axios = useAxios();

    const getAllAccount = async (): Promise<Account[]> => {
        const res = await axios.get(`${apiUrl}/account`);
        return res.data;
    };

    return {
        getAllAccount,
    };
};
