<template>
    <div v-if="!isLoading">
        <div class="space-y-4 w-full">
            <div class="flex flex-col w-full">
                <label>First name</label>
                <InputText
                    class="w-full p-inputtext-sm"
                    type="text"
                    v-model="account.firstName"
                />
            </div>
            <div class="flex flex-col w-full">
                <label>Last name</label>
                <InputText
                    class="w-full p-inputtext-sm"
                    type="text"
                    v-model="account.lastName"
                />
            </div>
            <div class="flex flex-col w-full">
                <label>Email</label>
                <InputText
                    class="w-full p-inputtext-sm"
                    type="text"
                    v-model="account.email"
                />
            </div>
            <div class="flex flex-col w-full">
                <label>Mobile number</label>
                <InputText
                    class="w-full p-inputtext-sm"
                    type="text"
                    v-model="account.mobileNumber"
                />
            </div>
            <div class="flex flex-col w-full">
                <label>Role</label>
                <Dropdown
                    class="w-full p-inputtext-sm"
                    v-model="account.roleId"
                    :options="roles"
                    optionValue="id"
                    optionLabel="name"
                    placeholder="Select a role"
                />
            </div>
            <div class="flex items-center justify-between">
                <label>Change Password</label>
                <InputSwitch v-model="isChangePassword" />
            </div>
            <Password
                class="w-full"
                v-if="isChangePassword"
                v-model="account.password"
                toggleMask
            />
        </div>
    </div>
    <div class="flex justify-center">
        <ProgressSpinner v-if="isLoading" />
    </div>
</template>

<script lang="ts">
// vue
import { defineComponent, onMounted, ref } from 'vue';

import { Account } from '@/api/models';

// services
import { useService } from '@/api/services';

// primevue
import InputText from 'primevue/inputtext';
import Dropdown from 'primevue/dropdown';
import Password from 'primevue/password';
import InputSwitch from 'primevue/inputswitch';
import ProgressSpinner from 'primevue/progressspinner';

export default defineComponent({
    components: {
        InputText,
        Dropdown,
        Password,
        InputSwitch,
        ProgressSpinner,
    },
    props: {
        accountId: {
            type: Number,
            default: null,
        },
        roles: {
            type: Array,
            default: () => [],
        },
    },
    setup(props) {
        // hooks
        const { getAccount, updateAccount } = useService();

        // reactive
        const account = ref({} as Account);
        const isChangePassword = ref(false);
        const isLoading = ref(true);

        // methods
        const save = async () => {
            if (!isChangePassword.value) {
                account.value.password = '';
            }
            await updateAccount(account.value);
        };

        // lifecycle
        onMounted(async () => {
            isLoading.value = true;
            account.value = await getAccount(props.accountId);
            isLoading.value = false;
        });

        return {
            account,
            save,
            isChangePassword,
            isLoading,
        };
    },
});
</script>
