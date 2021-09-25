// vue
import { useStore } from 'vuex';

// hooks
import { useAxios } from '@/hooks/axiosHook';

// models
import { ProductCategory } from '@/api/models';

export const useProductCategoryService = () => {
    const { apiUrl } = useStore().state;
    const axios = useAxios();

    const getAllProductCategory = async (): Promise<ProductCategory[]> => {
        const res = await axios.get(`${apiUrl}/product-category`);
        return res.data;
    };

    const getProductCategory = async (
        productCategoryId: number,
    ): Promise<ProductCategory> => {
        const res = await axios.get(
            `${apiUrl}/product-category/${productCategoryId}`,
        );
        return res.data;
    };

    const createProductCategory = async (
        productCategory: ProductCategory,
    ): Promise<ProductCategory> => {
        const res = await axios.post(
            `${apiUrl}/product-category`,
            productCategory,
        );
        return res.data;
    };

    const updateProductCategory = async (
        productCategory: ProductCategory,
    ): Promise<ProductCategory> => {
        const res = await axios.put(
            `${apiUrl}/product-category`,
            productCategory,
        );
        return res.data;
    };

    const deleteProductCategory = async (
        productCategoryId: number,
    ): Promise<ProductCategory> => {
        const res = await axios.deleteId(
            `${apiUrl}/product-category/${productCategoryId}`,
        );
        return res.data;
    };

    return {
        getAllProductCategory,
        getProductCategory,
        createProductCategory,
        updateProductCategory,
        deleteProductCategory,
    };
};
