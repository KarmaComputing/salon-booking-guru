<template>
    <div>
        <div class="text-2xl border-b pb-2 mb-4">Accounts</div>

        <div class="pb-4 space-y-2">
            <Button class="p-shadow-2" label="ADD ACCOUNT" icon="pi pi-plus" />
        </div>

        <Grid
            class="p-shadow-2 mb-4"
            :actionButtonConfig="actionButtonConfig"
            :gridConfig="accountGridConfig"
            :gridData="accounts"
            ref="grid"
        />
    </div>

    <!-- editor modal -->
    <Dialog
        class="w-11/12"
        v-model:visible="isModalVisible"
        header="Edit Account"
        :modal="true"
    >
        <AccountEditor :accountId="selectedAccount.id" />
    </Dialog>

    <!-- delete modal -->
    <ConfirmDialog
        header="Delete ConfirmDialogation"
        v-model:isVisible="isDeleteModalVisible"
        :confirmCallback="confirmDeleteAccount"
        :declineCallback="setIsDeleteModalVisible"
        :isLoading="isDeleteLoading"
    >
        <div>
            Are you sure you want to delete account
            <span class="font-semibold"> {{ selectedAccount.email }} </span>?
        </div>
    </ConfirmDialog>
</template>

<script lang="ts">
// vue
import { defineComponent, onMounted, ref, computed } from 'vue';

// primevue
import Dialog from 'primevue/dialog';
import Button from 'primevue/button';

// components
import Grid from '@/components/Grid.vue';
import AccountEditor from '@/components/AccountEditor.vue';
import ConfirmDialog from '@/components/ConfirmDialog.vue';

// config
import accountGridConfig from '@/config/grid/accountGrid';

// services
import { useService } from '@/api/services';

// models
import { Account } from '@/api/models';

export default defineComponent({
    components: {
        Grid,
        Dialog,
        Button,
        AccountEditor,
        ConfirmDialog,
    },
    setup() {
        // hooks
        const { getAllAccount, deleteAccount } = useService();

        // refs
        const grid = ref(null);

        // reactive
        const isModalVisible = ref(false);
        const isDeleteModalVisible = ref(false);
        const accounts = ref();
        const isDeleteLoading = ref(false);

        // computed
        const selectedAccount = computed((): Account => {
            return (grid?.value as any).selectedRow as Account;
        });

        // methods
        const refreshGrid = async () => {
            accounts.value = await getAllAccount();
        };

        const setIsModalVisible = (account: any) => {
            isModalVisible.value = !isModalVisible.value;
        };

        const setIsDeleteModalVisible = () => {
            isDeleteModalVisible.value = !isDeleteModalVisible.value;
        };

        const confirmDeleteAccount = async () => {
            isDeleteLoading.value = true;
            await deleteAccount(selectedAccount.value.id);
            isDeleteLoading.value = false;
            isDeleteModalVisible.value = false;
            refreshGrid();
        };

        // lifecycle
        onMounted(async () => {
            refreshGrid();
        });

        const actionButtonConfig = [
            {
                icon: 'pi pi-clock',
                route: '/account/availability',
            },
            {
                icon: 'pi pi-pencil',
                callback: setIsModalVisible,
            },
            {
                icon: 'pi pi-trash',
                callback: setIsDeleteModalVisible,
            },
        ];

        return {
            grid,
            accounts,
            accountGridConfig,
            actionButtonConfig,
            isModalVisible,
            isDeleteModalVisible,
            isDeleteLoading,
            setIsDeleteModalVisible,
            selectedAccount,
            confirmDeleteAccount,
        };
    },
});
</script>
