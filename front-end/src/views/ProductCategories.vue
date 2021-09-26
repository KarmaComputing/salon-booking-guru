<template>
    <div>
        <div class="text-2xl border-b pb-2 mb-4">Product Categories</div>
        <div class="pb-4 space-y-2">
            <Button
                class="p-shadow-2"
                label="ADD PRODUCT CATEGORY"
                icon="pi pi-plus"
                @click="
                    () => {
                        selectedProductCategory = {};
                        setIsEditorVisible(true);
                    }
                "
            />
        </div>
        <Grid
            class="p-shadow-2 mb-4"
            :actionButtonConfig="actionButtonConfig"
            :gridConfig="productCategoryGridConfig"
            :gridData="productCategories"
            :isLoading="isGridLoading"
            ref="grid"
        />
    </div>

    <!-- editor modal -->
    <BinaryDialog
        header="Product Category Editor"
        v-model:isVisible="isEditorVisible"
        :confirmCallback="saveProductCategory"
        :declineCallback="() => setIsEditorVisible(false)"
        confirmLabel="SAVE"
        confirmClass="p-button-success"
        :isLoading="isEditorLoading"
    >
        <ProductCategoryEditor
            ref="productCategoryEditor"
            :productCategoryId="selectedProductCategory.id"
            :roles="roles"
        />
    </BinaryDialog>

    <!-- delete modal -->
    <BinaryDialog
        header="Delete Confirmation"
        v-model:isVisible="isDeleteVisible"
        :confirmCallback="confirmDeleteProductCategory"
        :declineCallback="() => setIsDeleteVisible(false)"
        :isLoading="isDeleteLoading"
    >
        <div>
            Are you sure you want to delete product category
            <span class="font-semibold">
                {{ selectedProductCategory.name }} </span
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
import ProductCategoryEditor from '@/components/ProductCategoryEditor.vue';
import BinaryDialog from '@/components/BinaryDialog.vue';

// config
import productCategoryGridConfig from '@/config/grid/productCategoryGrid';

// services
import { useService } from '@/api/services';

// models
import { ProductCategory } from '@/api/models';

export default defineComponent({
    components: {
        Grid,
        Button,
        ProductCategoryEditor,
        BinaryDialog,
    },
    setup() {
        // hooks
        const { getAllProductCategory, deleteProductCategory, getAllRole } =
            useService();

        // refs
        const grid = ref<InstanceType<typeof Grid>>();
        const productCategoryEditor =
            ref<InstanceType<typeof ProductCategoryEditor>>();

        // reactive
        const productCategories = ref();
        const roles = ref();

        const isEditorVisible = ref(false);
        const isDeleteVisible = ref(false);

        const isGridLoading = ref(true);
        const isEditorLoading = ref(false);
        const isDeleteLoading = ref(false);

        // computed
        const selectedProductCategory = computed({
            get: (): ProductCategory => {
                return grid?.value?.selectedRow as ProductCategory;
            },
            set: (value: ProductCategory) => {
                if (grid.value) {
                    grid.value.selectedRow = value;
                }
            },
        });

        // methods
        const refreshGrid = async () => {
            isGridLoading.value = true;
            productCategories.value = await getAllProductCategory();
            isGridLoading.value = false;
        };

        const setIsEditorVisible = (value: boolean) => {
            isEditorVisible.value = value;
        };

        const setIsDeleteVisible = (value: boolean) => {
            isDeleteVisible.value = value;
        };

        const confirmDeleteProductCategory = async () => {
            isDeleteLoading.value = true;
            try {
                await deleteProductCategory(selectedProductCategory.value.id);
                isDeleteVisible.value = false;
                refreshGrid();
            } catch (e) {
                window.console.log(e);
                // fail toast here
            }

            isDeleteLoading.value = false;
        };

        const saveProductCategory = async () => {
            isEditorLoading.value = true;
            await productCategoryEditor?.value?.save();
            refreshGrid();
            setIsEditorVisible(false);
            isEditorLoading.value = false;
        };

        // lifecycle
        onMounted(async () => {
            refreshGrid();
            roles.value = await getAllRole();
        });

        const actionButtonConfig = [
            {
                icon: 'pi pi-clock',
                route: '/productCategory/availability',
            },
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
            productCategoryEditor,
            productCategories,
            productCategoryGridConfig,
            actionButtonConfig,
            isEditorVisible,
            isDeleteVisible,
            isGridLoading,
            isEditorLoading,
            isDeleteLoading,
            setIsDeleteVisible,
            setIsEditorVisible,
            selectedProductCategory,
            confirmDeleteProductCategory,
            saveProductCategory,
            roles,
        };
    },
});
</script>
