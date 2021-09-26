<template>
    <GridEditor
        title="Product Categories"
        addButtonLabel="ADD PRODUCT CATEGORY"
        editorHeader="Product Category Editor"
        rowName="product category"
        deleteRowValue="name"
        :gridConfig="gridConfig"
        :getData="getAllProductCategory"
        :deleteData="deleteProductCategory"
        :editorComponent="ProductCategoryEditor"
    />
</template>

<script lang="ts">
// vue
import { defineComponent } from 'vue';

// components
import GridEditor from '@/components/GridEditor.vue';
import ProductCategoryEditor from '@/components/ProductCategoryEditor.vue';

// services
import { useService } from '@/api/services';

// configs
import gridConfig from '@/config/grid/productCategoryGrid.ts';

export default defineComponent({
    components: {
        GridEditor,
    },
    setup() {
        // hooks
        const { getAllProductCategory, deleteProductCategory } = useService();

        return {
            getAllProductCategory,
            deleteProductCategory,
            gridConfig,
            ProductCategoryEditor,
        };
    },
});
</script>
