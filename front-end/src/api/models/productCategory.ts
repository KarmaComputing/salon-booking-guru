export interface ProductCategory {
    id: number;
    name: string;
}

export const emptyProductCategory = (): ProductCategory => {
    return {
        id: 0,
        name: '',
    };
};
