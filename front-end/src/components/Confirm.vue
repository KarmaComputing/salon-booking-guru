<template>
    <Dialog
        class="w-128 m-6"
        :visible="isVisible"
        @update:visible="updateIsVisible($event)"
        :header="header"
        :modal="true"
    >
        <slot />

        <template #footer>
            <div class="flex justify-end space-x-2">
                <Button
                    label="CANCEL"
                    class="p-button-text p-button-plain"
                    @click="declineCallback"
                />
                <Button
                    label="CONFIRM"
                    class="p-button-danger"
                    @click="confirmCallback"
                />
            </div>
        </template>
    </Dialog>
</template>

<script lang="ts">
// vue
import { defineComponent } from 'vue';

// primevue
import Dialog from 'primevue/dialog';

export default defineComponent({
    emits: ['update:isVisible'],
    props: {
        header: {
            type: String,
            default: 'Confirm',
        },
        isVisible: {
            type: Boolean,
            default: false,
        },
        confirmCallback: {
            type: Function,
            default: () => ({}),
        },
        declineCallback: {
            type: Function,
            default: () => ({}),
        },
        isLoading: {
            type: Boolean,
            default: false,
        },
    },
    components: {
        Dialog,
    },
    setup(props, { emit }) {
        // methods
        const updateIsVisible = (event: any) => {
            emit('update:isVisible', event);
        };

        return {
            updateIsVisible,
        };
    },
});
</script>
