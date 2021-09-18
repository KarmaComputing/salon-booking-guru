<template>
    <div class="text-2xl border-b pb-2 mb-8">Accounts</div>

    <div class="pb-4 space-y-2">
        <Button class="p-shadow-2" label="ADD ACCOUNT" icon="pi pi-plus" />
    </div>

    <Grid
        class="p-shadow-2 mb-4"
        :actionButtonConfig="actionButtonConfig"
        :gridConfig="accountGridConfig"
        :gridData="testAccounts"
    />
</template>

<script lang="ts">
// vue
import { defineComponent } from 'vue';

// components
import Grid from '@/components/Grid.vue';

// config
import accountGridConfig from '@/config/grid/accountGrid';

export default defineComponent({
    components: {
        Grid,
    },
    setup() {
        // properties
        const actionButtonConfig = [
            {
                icon: 'pi pi-clock',
                callback: () => {
                    console.log('check availability');
                },
            },
            {
                icon: 'pi pi-pencil',
                callback: () => {
                    console.log('edit');
                },
            },
            {
                icon: 'pi pi-trash',
                callback: () => {
                    console.log('delete');
                },
            },
        ];

        const testAccounts = [
            {
                name: 'Jake',
                email: 'jaketurner810@gmail.com',
                mobile: '07557140411',
                role: 'Employee',
            },
            {
                name: 'Hannah',
                email: 'hannaht95@gmail.com',
                mobile: '075271125374',
                role: 'Owner',
            },
            {
                name: 'Jamie Scollay',
                email: 'deltabrot@gmail.com',
                mobile: '075276128374',
                role: 'Employee',
            },
        ];

        return {
            accountGridConfig,
            actionButtonConfig,
            testAccounts,
        };
    },
});
</script>
