// services
import { useService } from '@/api/services';

export default [
    {
        title: 'First Name',
        field: 'firstName',
        type: 'text',
    },
    {
        title: 'Last Name',
        field: 'lastName',
        type: 'text',
    },
    {
        title: 'Email',
        field: 'email',
        type: 'text',
    },
    {
        title: 'Mobile Number',
        field: 'mobileNumber',
        type: 'text',
    },
    {
        title: 'Role',
        field: 'roleId',
        type: 'dropdown',
        dropdown: {
            getDataService: 'getAllRole',
            optionLabel: 'name',
            optionValue: 'id',
        },
    },
    {
        field: 'password',
        type: 'account-password',
    },
];
