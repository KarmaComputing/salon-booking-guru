<template>
    <GridEditor
        title="Products"
        addButtonLabel="ADD PRODUCT"
        editorHeader="Product Editor"
        rowName="product"
        deleteRowValue="name"
        :gridConfig="gridConfig"
        :getData="getAllProduct"
        :deleteData="deleteProduct"
        :editorComponent="ProductEditor"
    />
</template>

<script lang="ts">
// vue
import { defineComponent } from 'vue';

// components
import GridEditor from '@/components/GridEditor.vue';
import ProductEditor from '@/components/ProductEditor.vue';

// services
import { useService } from '@/api/services';

// configs
import gridConfig from '@/config/grid/productGrid.ts';

export default defineComponent({
    components: {
        GridEditor,
    },
    setup() {
        // hooks
        const { getAllProduct, deleteProduct } = useService();

        return {
            getAllProduct,
            deleteProduct,
            gridConfig,
            ProductEditor,
        };
    },
});
</script>
