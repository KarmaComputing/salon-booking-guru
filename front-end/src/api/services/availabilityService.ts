// vue
import { useStore } from 'vuex';

// hooks
import { useAxios } from '@/hooks/axiosHook';

// models
import { Availability } from '@/api/models';

export const useAvailabilityService = () => {
    const { apiUrl } = useStore().state;
    const axios = useAxios();

    const getAllAvailabilityByAccountId = async (
        accountId: number,
    ): Promise<Availability[]> => {
        const res = await axios.get(
            `${apiUrl}/account/${accountId}/availability`,
        );
        return res.data;
    };

    const getAvailability = async (
        availabilityId: number,
    ): Promise<Availability> => {
        const res = await axios.get(`${apiUrl}/availability/${availabilityId}`);
        return res.data;
    };

    const createMultipleAvailability = async (
        availabilities: Availability[],
    ): Promise<Availability[]> => {
        const res = await axios.post(`${apiUrl}/availability`, availabilities);
        return res.data;
    };

    const updateAvailability = async (
        availability: Availability,
    ): Promise<Availability> => {
        const res = await axios.put(`${apiUrl}/availability`, availability);
        return res.data;
    };

    const deleteAvailability = async (
        availabilityId: number,
    ): Promise<Availability> => {
        const res = await axios.deleteId(
            `${apiUrl}/availability/${availabilityId}`,
        );
        return res.data;
    };

    return {
        getAllAvailabilityByAccountId,
        getAvailability,
        createMultipleAvailability,
        updateAvailability,
        deleteAvailability,
    };
};
