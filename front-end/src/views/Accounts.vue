<template>
    <div>
        <div class="text-2xl border-b pb-2 mb-8">Accounts</div>

        <div class="pb-4 space-y-2">
            <Button class="p-shadow-2" label="ADD ACCOUNT" icon="pi pi-plus" />
        </div>

        <Grid
            class="p-shadow-2 mb-4"
            :actionButtonConfig="actionButtonConfig"
            :gridConfig="accountGridConfig"
            :gridData="accounts"
        />
    </div>
</template>

<script lang="ts">
// vue
import { defineComponent, onMounted, ref } from 'vue';

// components
import Grid from '@/components/Grid.vue';

// config
import accountGridConfig from '@/config/grid/accountGrid';

// services
import { useAccountService } from '@/api/services/accountService';

export default defineComponent({
    components: {
        Grid,
    },
    setup() {
        // hooks
        const { getAllAccount } = useAccountService();

        // reactive
        const accounts = ref();

        // methods
        const editCallback = (data: string) => {
            console.log(data);
        };

        // lifecycle
        onMounted(async () => {
            accounts.value = await getAllAccount();
        });

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
                callback: () => {
                    console.log('delete');
                },
            },
        ];

        return {
            accountGridConfig,
            actionButtonConfig,
            accounts,
        };
    },
});
</script>
