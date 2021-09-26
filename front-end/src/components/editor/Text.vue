<template>
    <InputText
        class="w-full p-inputtext-sm"
        :modelValue="modelValue"
        @update:modelValue="update($event)"
    />
</template>

<script lang="ts">
// vue
import { defineComponent } from 'vue';

// primevue
import InputText from 'primevue/inputtext';

export default defineComponent({
    props: {
        modelValue: {
            type: String,
            default: '',
        },
    },
    components: {
        InputText,
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
