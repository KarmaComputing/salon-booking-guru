<template>
    <div>
        <div class="text-2xl border-b pb-2 mb-4">{{ title }}</div>
        <div class="pb-4 space-y-2">
            <Button
                class="p-shadow-2"
                :label="addButtonLabel"
                icon="pi pi-plus"
                @click="
                    () => {
                        selectedRow = {};
                        setIsEditorVisible(true);
                    }
                "
            />
        </div>
        <Grid
            class="p-shadow-2 mb-4"
            :actionButtonConfig="actionButtonConfig"
            :gridConfig="gridConfig"
            :gridData="gridData"
            :isLoading="isGridLoading"
            ref="grid"
        />
    </div>

    <!-- editor modal -->
    <BinaryDialog
        :header="editorHeader"
        v-model:isVisible="isEditorVisible"
        :confirmCallback="saveRow"
        :declineCallback="() => setIsEditorVisible(false)"
        confirmLabel="SAVE"
        confirmClass="p-button-success"
        :isLoading="isEditorLoading"
    >
        <component ref="editor" :is="editorComponent" :id="selectedRow.id" />
    </BinaryDialog>

    <!-- delete modal -->
    <BinaryDialog
        header="Delete Confirmation"
        v-model:isVisible="isDeleteVisible"
        :confirmCallback="confirmDeleteRow"
        :declineCallback="() => setIsDeleteVisible(false)"
        :isLoading="isDeleteLoading"
    >
        <div>
            Are you sure you want to delete {{ rowName }}
            <span class="font-semibold">
                {{ selectedRow[deleteRowValue] }} </span
            >?
        </div>
    </BinaryDialog>
</template>

<script lang="ts">
// vue
import { defineComponent, onMounted, ref, computed } from 'vue';

// primevue
import Button from 'primevue/button';

// components
import Grid from '@/components/Grid.vue';
import BinaryDialog from '@/components/BinaryDialog.vue';

export default defineComponent({
    props: {
        title: {
            type: String,
            default: 'Default Grid Editor',
        },
        addButtonLabel: {
            type: String,
            default: 'ADD ROW',
        },
        editorHeader: {
            type: String,
            default: 'Editor',
        },
        rowName: {
            type: String,
            default: 'row',
        },
        deleteRowValue: {
            type: String,
            default: 'name',
        },
        gridConfig: {
            type: Object,
            default: () => ({}),
        },
        getData: {
            type: Function,
            default: () => ({}),
        },
        deleteData: {
            type: Function,
            default: () => ({}),
        },
        editorComponent: {
            type: Object,
            default: () => ({}),
        },
    },
    components: {
        Grid,
        Button,
        BinaryDialog,
    },
    setup(props) {
        // refs
        const grid = ref<InstanceType<typeof Grid>>();
        // const editor = ref<InstanceType<typeof AccountEditor>>();
        const editor = ref<InstanceType<typeof props.editorComponent>>();

        // reactive
        const gridData = ref([]);

        const isEditorVisible = ref(false);
        const isDeleteVisible = ref(false);

        const isGridLoading = ref(true);
        const isEditorLoading = ref(false);
        const isDeleteLoading = ref(false);

        // computed
        const selectedRow = computed({
            get: () => {
                return grid?.value?.selectedRow;
            },
            set: (value) => {
                if (grid.value) {
                    grid.value.selectedRow = value;
                }
            },
        });

        // methods
        const refreshGrid = async () => {
            isGridLoading.value = true;
            gridData.value = await props.getData();
            isGridLoading.value = false;
        };

        const setIsEditorVisible = (value: boolean) => {
            isEditorVisible.value = value;
        };

        const setIsDeleteVisible = (value: boolean) => {
            isDeleteVisible.value = value;
        };

        const confirmDeleteRow = async () => {
            isDeleteLoading.value = true;
            try {
                await props.deleteData(selectedRow.value.id);
                isDeleteVisible.value = false;
                refreshGrid();
            } catch (e) {
                window.console.log(e);
                // fail toast here
            }

            isDeleteLoading.value = false;
        };

        const saveRow = async () => {
            isEditorLoading.value = true;
            try {
                await editor?.value?.save();
                refreshGrid();
                setIsEditorVisible(false);
            } catch (e) {
                window.console.log(e);
                // fail toast here
            }

            isEditorLoading.value = false;
        };

        // lifecycle
        onMounted(async () => {
            refreshGrid();
        });

        const actionButtonConfig = [
            {
                icon: 'pi pi-pencil',
                callback: () => setIsEditorVisible(true),
            },
            {
                icon: 'pi pi-trash',
                callback: () => setIsDeleteVisible(true),
            },
        ];

        return {
            grid,
            editor,
            gridData,
            actionButtonConfig,
            isEditorVisible,
            isDeleteVisible,
            isGridLoading,
            isEditorLoading,
            isDeleteLoading,
            setIsDeleteVisible,
            setIsEditorVisible,
            selectedRow,
            confirmDeleteRow,
            saveRow,
        };
    },
});
</script>
