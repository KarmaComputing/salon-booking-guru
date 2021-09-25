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
        <div>
            {{ selectedAccount.data }}
        </div>
        <div class="flex justify-center">
            <div class="space-y-5 w-11/12">
                <span class="p-float-label">
                    <InputText
                        class="w-full p-inputtext-sm p-shadow-2"
                        type="text"
                        v-model="value"
                    />
                    <label for="username">First name</label>
                </span>
                <span class="p-float-label">
                    <InputText
                        class="w-full p-inputtext-sm p-shadow-2"
                        type="text"
                        v-model="value"
                    />
                    <label for="username">Last name</label>
                </span>
                <span class="p-float-label">
                    <InputText
                        class="w-full p-inputtext-sm p-shadow-2"
                        type="text"
                        v-model="value"
                    />
                    <label for="username">Email</label>
                </span>
                <span class="p-float-label">
                    <InputText
                        class="w-full p-inputtext-sm p-shadow-2"
                        type="text"
                        v-model="value"
                    />
                    <label for="username">Mobile number</label>
                </span>
                <Dropdown
                    class="w-full p-inputtext-sm p-shadow-2"
                    v-model="selectedCity"
                    :options="cities"
                    optionLabel="name"
                    placeholder="Select a role"
                />
                <div class="space-x-2">
                    <Button label="Cancel" class="p-button-raised" />
                    <Button
                        label="Save"
                        class="p-button-raised p-button-danger"
                    />
                </div>
            </div>
        </div>
    </Dialog>

    <!-- delete modal -->
    <Dialog
        class="w-128 m-6"
        v-model:visible="isDeleteModalVisible"
        header="Delete Confirmation"
        :modal="true"
    >
        <div>
            Are you sure you want to delete account
            <span class="font-semibold"> {{ selectedAccount.email }} </span>?
        </div>

        <template #footer>
            <div class="flex justify-end space-x-2">
                <Button
                    label="CANCEL"
                    class="p-button-text p-button-plain"
                    @click="setIsDeleteModalVisible"
                />
                <Button
                    label="CONFIRM"
                    class="p-button-danger"
                    @click="confirmDeleteAccount"
                />
            </div>
        </template>
    </Dialog>
</template>

<script lang="ts">
// vue
import { defineComponent, onMounted, ref, computed } from 'vue';

// primevue
import Dialog from 'primevue/dialog';
import InputText from 'primevue/inputtext';
import Dropdown from 'primevue/dropdown';
import Button from 'primevue/button';

// components
import Grid from '@/components/Grid.vue';

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
        InputText,
        Dropdown,
        Button,
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

        // computed
        const selectedAccount = computed((): Account => {
            return (grid?.value as any).selectedRow as Account;
        });

        // methods
        const setIsModalVisible = (account: any) => {
            isModalVisible.value = !isModalVisible.value;
        };

        const setIsDeleteModalVisible = () => {
            isDeleteModalVisible.value = !isDeleteModalVisible.value;
        };

        const confirmDeleteAccount = async () => {
            await deleteAccount(selectedAccount.value.id);
            isDeleteModalVisible.value = false;
            accounts.value = await getAllAccount();
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
            setIsDeleteModalVisible,
            selectedAccount,
            confirmDeleteAccount,
        };
    },
});
</script>
