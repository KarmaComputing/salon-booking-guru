<template>
    <form
        class="
            login-form
            mx-auto
            bg-white
            p-shadow-2
            rounded
            p-4
            space-y-4
            mt-32
        "
        @submit.prevent="logIn"
    >
        <div class="flex flex-col">
            <label>Email</label>
            <InputText v-model="credentials.email" />
        </div>

        <div class="flex flex-col w-full">
            <label>Password</label>
            <Password
                class="w-full"
                v-model="credentials.password"
                :feedback="false"
            />
        </div>

        <Button class="w-full" label="Login" type="submit" />
    </form>
</template>

<script lang="ts">
// vue
import { defineComponent, ref } from 'vue';

// services
import { useService } from '@/api/services';

// models
import { Credentials } from '@/api/models';

// primevue
import InputText from 'primevue/inputtext';
import Password from 'primevue/password';
import Button from 'primevue/button';

export default defineComponent({
    components: {
        InputText,
        Password,
        Button,
    },
    setup() {
        // hooks
        const { authenticate } = useService();

        // reactive
        const credentials = ref<Credentials>({} as Credentials);

        // methods
        const logIn = async () => {
            await authenticate(credentials.value);
        };

        return {
            credentials,
            logIn,
        };
    },
});
</script>

<style scoped>
.login-form {
    max-width: 400px;
}
</style>
