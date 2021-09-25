// vue
import { useStore } from 'vuex';

// hooks
import { useAxios } from '@/hooks/axiosHook';

// models
import { Qualification } from '@/api/models';

export const useQualificationService = () => {
    const { apiUrl } = useStore().state;
    const axios = useAxios();

    const getAllQualification = async (): Promise<Qualification[]> => {
        const res = await axios.get(`${apiUrl}/qualificaton`);
        return res.data;
    };

    const getQualification = async (
        qualificatonId: number,
    ): Promise<Qualification> => {
        const res = await axios.get(`${apiUrl}/qualificaton/${qualificatonId}`);
        return res.data;
    };

    const createQualification = async (
        qualificaton: Qualification,
    ): Promise<Qualification> => {
        const res = await axios.post(`${apiUrl}/qualificaton`, qualificaton);
        return res.data;
    };

    const updateQualification = async (
        qualificaton: Qualification,
    ): Promise<Qualification> => {
        const res = await axios.put(`${apiUrl}/qualificaton`, qualificaton);
        return res.data;
    };

    const deleteQualification = async (
        qualificatonId: number,
    ): Promise<Qualification> => {
        const res = await axios.deleteId(
            `${apiUrl}/qualificaton/${qualificatonId}`,
        );
        return res.data;
    };

    return {
        getAllQualification,
        getQualification,
        createQualification,
        updateQualification,
        deleteQualification,
    };
};
