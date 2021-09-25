// vue
import { useStore } from 'vuex';

// hooks
import { useAxios } from '@/hooks/axiosHook';

// models
import { Role } from '@/api/models';

export const useRoleService = () => {
    const { apiUrl } = useStore().state;
    const axios = useAxios();

    const getAllRole = async (): Promise<Role[]> => {
        const res = await axios.get(`${apiUrl}/role`);
        return res.data;
    };

    return {
        getAllRole,
    };
};
