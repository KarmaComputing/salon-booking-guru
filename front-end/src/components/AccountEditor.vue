<template>
    <div class="flex justify-center">
        <div class="pt-2 space-y-5 w-11/12">
            <span class="p-float-label">
                <InputText class="w-full p-inputtext-sm" type="text" />
                <label for="username">First name</label>
            </span>
            <span class="p-float-label">
                <InputText class="w-full p-inputtext-sm" type="text" />
                <label for="username">Last name</label>
            </span>
            <span class="p-float-label">
                <InputText class="w-full p-inputtext-sm" type="text" />
                <label for="username">Email</label>
            </span>
            <span class="p-float-label">
                <InputText class="w-full p-inputtext-sm" type="text" />
                <label for="username">Mobile number</label>
            </span>
            <Dropdown
                class="w-full p-inputtext-sm"
                :options="cities"
                optionLabel="name"
                placeholder="Select a role"
            />
            <Password v-model="value3" toggleMask></Password>

            <div class="space-x-2">
                <Button label="Cancel" class="p-button-raised" />
                <Button label="Save" class="p-button-raised p-button-danger" />
            </div>
        </div>
    </div>
</template>

<script lang="ts">
// vue
import { defineComponent } from 'vue';

// primevue
import InputText from 'primevue/inputtext';
import Dropdown from 'primevue/dropdown';
import Button from 'primevue/button';
import Password from 'primevue/password';

export default defineComponent({
    components: {
        InputText,
        Dropdown,
        Button,
        Password,
    },
    props: {
        accountId: {
            type: Number,
            default: null,
        },
    },
    setup() {
        return {};
    },
});
</script>
