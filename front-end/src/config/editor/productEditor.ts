// services
import { useService } from '@/api/services';

export default [
    {
        title: 'Name',
        field: 'name',
        type: 'text',
    },
    {
        title: 'Description',
        field: 'description',
        type: 'text',
    },
    {
        title: 'Category',
        field: 'productCategoryId',
        type: 'dropdown',
        dropdown: {
            getDataService: 'getAllProductCategory',
            optionLabel: 'name',
            optionValue: 'id',
        },
    },
    {
        title: 'Price',
        field: 'price',
        type: 'currency',
    },
    {
        title: 'Deposit',
        field: 'deposit',
        type: 'currency',
    },
    {
        title: 'Duration',
        field: 'duration',
        type: 'number',
    },
];
