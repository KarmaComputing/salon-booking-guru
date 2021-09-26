// components
import Currency from '@/components/grid/Currency.vue';
import Hours from '@/components/grid/Hours.vue';

export default [
    {
        title: 'Name',
        field: 'name',
    },
    {
        title: 'Category',
        field: 'productCategoryName',
    },
    {
        title: 'Description',
        field: 'description',
    },
    {
        title: 'Price',
        field: 'price',
        cellRenderer: Currency,
    },
    {
        title: 'Deposit',
        field: 'deposit',
        cellRenderer: Currency,
    },
    {
        title: 'Duration',
        field: 'duration',
        cellRenderer: Hours,
    },
];
