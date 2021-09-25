<template>
    <div v-if="account">
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
                    :options="cities"
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

export default defineComponent({
    components: {
        InputText,
        Dropdown,
        Password,
        InputSwitch,
    },
    props: {
        accountId: {
            type: Number,
            default: null,
        },
    },
    setup(props) {
        // hooks
        const { getAccount, updateAccount } = useService();
        const account = ref({} as Account);
        const isChangePassword = ref(false);

        // methods
        const save = async () => {
            if (!isChangePassword.value) {
                account.value.password = '';
            }
            await updateAccount(account.value);
        };

        // lifecycle
        onMounted(async () => {
            account.value = await getAccount(props.accountId);
        });

        return {
            account,
            save,
            isChangePassword,
        };
    },
});
</script>
