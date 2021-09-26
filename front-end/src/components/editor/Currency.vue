<template>
    <InputNumber
        class="w-full p-inputtext-sm"
        :modelValue="modelValue"
        @update:modelValue="update($event)"
        mode="currency"
        currency="GBP"
        locale="en-GB"
        :useGrouping="false"
        @input="update($event.value)"
    />
</template>

<script lang="ts">
// vue
import { defineComponent } from 'vue';

// primevue
import InputNumber from 'primevue/inputnumber';

export default defineComponent({
    props: {
        modelValue: {
            type: Number,
            default: null,
        },
    },
    components: {
        InputNumber,
    },
    setup(props, { emit }) {
        const update = (event: any) => {
            emit('update:modelValue', event);
        };

        return {
            update,
        };
    },
});
</script>
