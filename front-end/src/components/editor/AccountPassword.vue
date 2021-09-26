<template>
    <div class="flex items-center justify-between mb-2" v-if="rowId">
        <label>Change Password</label>
        <InputSwitch v-model="isChangePassword" />
    </div>

    <label v-if="!rowId">Password</label>
    <Password
        class="w-full p-inputtext-sm"
        :modelValue="modelValue"
        @update:modelValue="update($event)"
        v-if="isChangePassword || !rowId"
        toggleMask
    />
</template>

<script lang="ts">
// vue
import { defineComponent, ref } from 'vue';

// primevue
import Password from 'primevue/password';
import InputSwitch from 'primevue/inputswitch';

export default defineComponent({
    props: {
        modelValue: {
            type: String,
            default: '',
        },
        rowId: {
            type: Number,
            default: null,
        },
    },
    components: {
        Password,
        InputSwitch,
    },
    setup(props, { emit }) {
        // reactive
        const isChangePassword = ref(false);

        // methods
        const update = (event: any) => {
            emit('update:modelValue', event);
        };

        return {
            update,
            isChangePassword,
        };
    },
});
</script>
