<template>
    <GridEditor
        title="Products"
        addButtonLabel="ADD PRODUCT"
        editorHeader="Product Editor"
        rowName="product"
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
import gridConfig from '@/config/grid/productGrid';
import editorConfig from '@/config/editor/productEditor';

export default defineComponent({
    components: {
        GridEditor,
    },
    setup() {
        // hooks
        const { generateDataServices } = useService();

        // properties
        const dataServices = generateDataServices('Product', true);

        return {
            dataServices,
            gridConfig,
            editorConfig,
        };
    },
});
</script>
