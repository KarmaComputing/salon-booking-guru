export interface Product {
    id: number;
    productCategoryId: number;
    name: string;
    description: string;
    price: number;
    deposit: number;
    duration: number;
}

export const emptyProduct = (): Product => {
    return {
        id: 0,
        productCategoryId: 0,
        name: '',
        description: '',
        price: 0,
        deposit: 0,
        duration: 0,
    };
};
