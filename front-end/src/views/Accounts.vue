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
            :gridData="accountSummaries"
            ref="grid"
        />
    </div>

    <!-- editor modal -->
    <BinaryDialog
        header="Edit Account"
        v-model:isVisible="isEditorVisible"
        :confirmCallback="saveAccount"
        :declineCallback="() => setIsEditorVisible(false)"
        confirmLabel="SAVE"
        confirmClass="p-button-success"
    >
        <AccountEditor
            ref="accountEditor"
            :accountId="selectedAccount.id"
            :roles="roles"
        />
    </BinaryDialog>

    <!-- delete modal -->
    <BinaryDialog
        header="Delete Confirmation"
        v-model:isVisible="isDeleteVisible"
        :confirmCallback="confirmDeleteAccount"
        :declineCallback="() => setIsDeleteVisible(false)"
        :isLoading="isDeleteLoading"
    >
        <div>
            Are you sure you want to delete account
            <span class="font-semibold"> {{ selectedAccount.email }} </span>?
        </div>
    </BinaryDialog>
</template>

<script lang="ts">
// vue
import { defineComponent, onMounted, ref, computed } from 'vue';

// primevue
import Button from 'primevue/button';

// components
import Grid from '@/components/Grid.vue';
import AccountEditor from '@/components/AccountEditor.vue';
import BinaryDialog from '@/components/BinaryDialog.vue';

// config
import accountGridConfig from '@/config/grid/accountGrid';

// services
import { useService } from '@/api/services';

// models
import { Account } from '@/api/models';

export default defineComponent({
    components: {
        Grid,
        Button,
        AccountEditor,
        BinaryDialog,
    },
    setup() {
        // hooks
        const { getAllAccountSummary, deleteAccount, getAllRole } =
            useService();

        // refs
        const grid = ref<InstanceType<typeof Grid>>();
        const accountEditor = ref<InstanceType<typeof AccountEditor>>();

        // reactive
        const accountSummaries = ref();
        const roles = ref();
        const isEditorVisible = ref(false);
        const isDeleteVisible = ref(false);
        const isDeleteLoading = ref(false);

        // computed
        const selectedAccount = computed((): Account => {
            return grid?.value?.selectedRow as Account;
        });

        // methods
        const refreshGrid = async () => {
            accountSummaries.value = await getAllAccountSummary();
        };

        const setIsEditorVisible = (value: boolean) => {
            isEditorVisible.value = value;
        };

        const setIsDeleteVisible = (value: boolean) => {
            isDeleteVisible.value = value;
        };

        const confirmDeleteAccount = async () => {
            isDeleteLoading.value = true;
            await deleteAccount(selectedAccount.value.id);
            isDeleteLoading.value = false;
            isDeleteVisible.value = false;
            refreshGrid();
        };

        const saveAccount = async () => {
            await accountEditor?.value?.save();
            refreshGrid();
            setIsEditorVisible(false);
        };

        // lifecycle
        onMounted(async () => {
            refreshGrid();
            roles.value = await getAllRole();
        });

        const actionButtonConfig = [
            {
                icon: 'pi pi-clock',
                route: '/account/availability',
            },
            {
                icon: 'pi pi-pencil',
                callback: () => setIsEditorVisible(true),
            },
            {
                icon: 'pi pi-trash',
                callback: () => setIsDeleteVisible(true),
            },
        ];

        return {
            grid,
            accountEditor,
            accountSummaries,
            accountGridConfig,
            actionButtonConfig,
            isEditorVisible,
            isDeleteVisible,
            isDeleteLoading,
            setIsDeleteVisible,
            setIsEditorVisible,
            selectedAccount,
            confirmDeleteAccount,
            saveAccount,
            roles,
        };
    },
});
</script>
