<template>
    <GridEditor
        title="Product Categories"
        addButtonLabel="ADD PRODUCT CATEGORY"
        editorHeader="Product Category Editor"
        rowName="product category"
        deleteRowValue="name"
        :gridConfig="gridConfig"
        :editorConfig="editorConfig"
        :dataServices="dataServices"
    />
</template>

<script lang="ts">
// vue
import { defineComponent } from 'vue';

// components
import GridEditor from '@/components/GridEditor.vue';

// services
import { useService } from '@/api/services';

// configs
import gridConfig from '@/config/grid/productCategoryGrid';
import editorConfig from '@/config/editor/productCategoryEditor';

export default defineComponent({
    components: {
        GridEditor,
    },
    setup() {
        // hooks
        const { generateDataServices } = useService();

        // properties
        const dataServices = generateDataServices('ProductCategory');

        return {
            dataServices,
            gridConfig,
            editorConfig,
        };
    },
});
</script>
