<template>
    <GridEditor
        title="Accounts"
        addButtonLabel="ADD ACCOUNT"
        editorHeader="Account Editor"
        rowName="account"
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
import gridConfig from '@/config/grid/accountGrid.ts';
import editorConfig from '@/config/editor/accountEditor.ts';

export default defineComponent({
    components: {
        GridEditor,
    },
    setup() {
        // hooks
        const { generateDataServices } = useService();

        // properties
        const dataServices = generateDataServices('Account', true);

        return {
            dataServices,
            gridConfig,
            editorConfig,
        };
    },
});
</script>
