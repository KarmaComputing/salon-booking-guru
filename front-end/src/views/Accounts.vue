<template>
    <div class="pt-20 px-4">
        <div class="pb-4 space-y-2">
            <div class="text-xl">Accounts</div>
            <Button class="p-shadow-2">Add Account</Button>
        </div>

        <SalonGrid
            class="p-shadow-2 mb-4"
            :actionButtonConfig="actionButtonConfig"
            :gridConfig="config"
        />
    </div>
</template>

<script lang="ts">
// vue
import { defineComponent } from 'vue';

// components
import SalonGrid from '../components/SalonGrid.vue';

export default defineComponent({
    components: {
        SalonGrid,
    },
    setup() {
        // grid config
        const config = [
            {
                title: 'Name',
                field: 'firstName',
            },
            {
                title: 'Email',
                field: 'email',
            },
            {
                title: 'Mobile',
                field: 'mobileNumber',
            },
            {
                title: 'Role',
                field: 'roleId',
            },
        ];

        const editCallback = (data: string) => {
            console.log(data);
        };

        const actionButtonConfig = [
            {
                icon: 'pi pi-clock',
                route: '/account/availability',
            },
            {
                icon: 'pi pi-pencil',
                route: '/account/editor',
                callback: editCallback,
            },
            {
                icon: 'pi pi-trash',
            },
        ];
        return {
            config,
            actionButtonConfig,
        };
    },
});
</script>
