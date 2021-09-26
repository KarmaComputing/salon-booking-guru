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
        const res = await axios.get(`${apiUrl}/qualification`);
        return res.data;
    };

    const getQualification = async (
        qualificationId: number,
    ): Promise<Qualification> => {
        const res = await axios.get(
            `${apiUrl}/qualification/${qualificationId}`,
        );
        return res.data;
    };

    const createQualification = async (
        qualification: Qualification,
    ): Promise<Qualification> => {
        const res = await axios.post(`${apiUrl}/qualification`, qualification);
        return res.data;
    };

    const updateQualification = async (
        qualification: Qualification,
    ): Promise<Qualification> => {
        const res = await axios.put(`${apiUrl}/qualification`, qualification);
        return res.data;
    };

    const deleteQualification = async (
        qualificationId: number,
    ): Promise<Qualification> => {
        const res = await axios.deleteId(
            `${apiUrl}/qualification/${qualificationId}`,
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
