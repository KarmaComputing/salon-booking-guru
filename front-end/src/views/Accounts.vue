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
        />
    </div>
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
    <Dialog
        class="w-11/12"
        v-model:visible="isDeleteModalVisible"
        header="Delete Confirmation"
        :modal="true"
    >
        <div class="mb-4">Are you sure you want to delete this account?</div>
        <div class="space-x-2">
            <Button label="No" class="p-button-raised" />
            <Button label="Yes" class="p-button-raised p-button-danger" />
        </div>
    </Dialog>
</template>

<script lang="ts">
// vue
import { defineComponent, onMounted, ref } from 'vue';

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
import { useAccountService } from '@/api/services/accountService';

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
        const { getAllAccount } = useAccountService();

        // reactive
        const isModalVisible = ref(false);
        const isDeleteModalVisible = ref(false);
        const accounts = ref();
        const selectedAccount = ref();

        // methods
        const setIsModalVisible = (account: any) => {
            isModalVisible.value = !isModalVisible.value;
            selectedAccount.value = account;
            console.log(account.data.firstName);
        };

        const setIsDeleteModalVisible = () => {
            isDeleteModalVisible.value = !isDeleteModalVisible.value;
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
            accountGridConfig,
            actionButtonConfig,
            accounts,
            isModalVisible,
            isDeleteModalVisible,
            selectedAccount,
        };
    },
});
</script>
