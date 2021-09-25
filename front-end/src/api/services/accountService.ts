// axios
import axios, { AxiosError } from 'axios';
import { Account } from '../models/account';

export const useAccountService = () => {
    const getAllAccount = async (): Promise<Account[]> => {
        const res = await axios.get('http://127.0.0.1:8085/v1/account', {
            headers: {
                Authorization:
                    'Bearer KjzaAHGT+WSMPgXSSyXdXxiFERriP2Jw8+xYWFOwAqaxxi4s5ZwRxc6cXElJ54vE5TYkoaZekCA1jU1dUi/VVg==',
            },
        });
        return res.data;
    };

    return {
        getAllAccount,
    };
};
