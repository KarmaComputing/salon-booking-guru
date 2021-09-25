<template>
    <div class="pt-20 flex justify-center">
        <div class="space-y-5 w-10/12">
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
            <Password
                class="p-inputtext-sm p-shadow-2"
                v-model="value3"
                toggleMask
            ></Password>
        </div>
    </div>
</template>

<script lang="ts">
// vue
import { defineComponent } from 'vue';

//primevue
import InputText from 'primevue/inputtext';
import Dropdown from 'primevue/dropdown';
import Password from 'primevue/password';

export default defineComponent({
    components: {
        InputText,
        Dropdown,
        Password,
    },
    props: {
        account: {
            type: Object,
            required: true,
        },
    },
    setup() {
        return {};
    },
});
</script>
