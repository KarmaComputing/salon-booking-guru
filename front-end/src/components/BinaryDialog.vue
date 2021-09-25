<template>
    <Dialog
        class="p-dialog-sm w-full m-6"
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
                    v-if="!isLoading"
                />
                <Button
                    :label="confirmLabel"
                    :class="confirmClass"
                    @click="confirmCallback"
                    v-if="!isLoading"
                />
                <ProgressSpinner v-if="isLoading" />
            </div>
        </template>
    </Dialog>
</template>

<script lang="ts">
// vue
import { defineComponent } from 'vue';

// primevue
import Dialog from 'primevue/dialog';
import ProgressSpinner from 'primevue/progressspinner';

export default defineComponent({
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
        confirmLabel: {
            type: String,
            default: 'CONFIRM',
        },
        confirmClass: {
            type: String,
            default: 'p-button-danger',
        },
    },
    components: {
        Dialog,
        ProgressSpinner,
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

<style>
.p-progress-spinner {
    width: 40px;
    height: 40px;
}

.p-dialog-sm {
    max-width: 500px;
}
</style>
