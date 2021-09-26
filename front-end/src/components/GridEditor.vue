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
        <Editor
            ref="editor"
            :editorConfig="hydratedEditorConfig"
            :selectedRow="selectedRow"
            :getData="dataServices.get"
            :createData="dataServices.create"
            :updateData="dataServices.update"
        />
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
import Editor from '@/components/Editor.vue';

// services
import { useService } from '@/api/services';

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
        editorConfig: {
            type: Object,
            default: () => ({}),
        },
        dataServices: {
            type: Object,
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
        Editor,
    },
    setup(props) {
        // refs
        const grid = ref<InstanceType<typeof Grid>>();
        const editor = ref<InstanceType<any>>();

        // reactive
        const gridData = ref([]);

        const isEditorVisible = ref(false);
        const isDeleteVisible = ref(false);

        const isGridLoading = ref(true);
        const isEditorLoading = ref(false);
        const isDeleteLoading = ref(false);

        const hydratedEditorConfig = ref<any>([]);

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
            gridData.value = await props.dataServices.getAll();
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
                await props.dataServices.delete(selectedRow.value.id);
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

            // hydrate dropdown property in editorConfig
            hydratedEditorConfig.value = props.editorConfig;
            props.editorConfig.forEach(
                async (config: typeof props.editorConfig, i: number) => {
                    if (config.type === 'dropdown') {
                        hydratedEditorConfig.value[i].dropdown.options =
                            await useService()[
                                config.dropdown.getDataService
                            ]();
                    }
                },
            );
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
            hydratedEditorConfig,
        };
    },
});
</script>
