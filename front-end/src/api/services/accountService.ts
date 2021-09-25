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

    const getAccount = async (accountId: number): Promise<Account> => {
        const res = await axios.get(`${apiUrl}/account/${accountId}`);
        return res.data;
    };

    const createAccount = async (account: Account): Promise<Account> => {
        const res = await axios.post(`${apiUrl}/account`, account);
        return res.data;
    };

    const updateAccount = async (account: Account): Promise<Account> => {
        const res = await axios.put(`${apiUrl}/account`, account);
        return res.data;
    };

    const deleteAccount = async (accountId: number): Promise<Account> => {
        const res = await axios.deleteId(`${apiUrl}/account/${accountId}`);
        return res.data;
    };

    return {
        getAllAccount,
        getAccount,
        createAccount,
        updateAccount,
        deleteAccount,
    };
};
