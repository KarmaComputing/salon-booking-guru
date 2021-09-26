<template>
    <Dropdown
        class="w-full"
        :modelValue="modelValue"
        @update:modelValue="update($event)"
        :options="dropdown.options"
        :optionValue="dropdown.optionValue"
        :optionLabel="dropdown.optionLabel"
        placeholder="Select"
    />
</template>

<script lang="ts">
// vue
import { defineComponent, ref, onMounted } from 'vue';

// services
import { useService } from '@/api/services';

// primevue
import Dropdown from 'primevue/dropdown';

export default defineComponent({
    props: {
        modelValue: {
            type: Number,
            default: null,
        },
        dropdown: {
            type: Object,
            default: () => ({}),
        },
    },
    components: {
        Dropdown,
    },
    setup(props, { emit }) {
        // methods
        const update = (event: any) => {
            emit('update:modelValue', event);
        };

        return {
            update,
        };
    },
});
</script>
