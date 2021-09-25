// vue
import { useStore } from 'vuex';

// hooks
import { useAxios } from '@/hooks/axiosHook';

// models
import { Product } from '@/api/models';

export const useProductService = () => {
    const { apiUrl } = useStore().state;
    const axios = useAxios();

    const getAllProduct = async (): Promise<Product[]> => {
        const res = await axios.get(`${apiUrl}/product`);
        return res.data;
    };

    const getProduct = async (productId: number): Promise<Product> => {
        const res = await axios.get(`${apiUrl}/product/${productId}`);
        return res.data;
    };

    const getProductQualificationNames = async (
        productId: number,
    ): Promise<string[]> => {
        const res = await axios.get(
            `${apiUrl}/product/${productId}/qualification`,
        );
        return res.data;
    };

    const createProduct = async (product: Product): Promise<Product> => {
        const res = await axios.post(`${apiUrl}/product`, product);
        return res.data;
    };

    const updateProduct = async (product: Product): Promise<Product> => {
        const res = await axios.put(`${apiUrl}/product`, product);
        return res.data;
    };

    const upsertProductQualifications = async (
        productId: number,
        qualificationIds: number[],
    ): Promise<string[]> => {
        const res = await axios.put(
            `${apiUrl}/product/${productId}/qualification`,
            qualificationIds,
        );
        return res.data;
    };

    const deleteProduct = async (productId: number): Promise<Product> => {
        const res = await axios.deleteId(`${apiUrl}/product/${productId}`);
        return res.data;
    };

    return {
        getAllProduct,
        getProduct,
        getProductQualificationNames,
        createProduct,
        updateProduct,
        upsertProductQualifications,
        deleteProduct,
    };
};
