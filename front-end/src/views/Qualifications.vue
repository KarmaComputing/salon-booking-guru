<template>
    <GridEditor
        title="Qualifications"
        addButtonLabel="ADD QUALIFICATION"
        editorHeader="Qualification Editor"
        rowName="qualification"
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
import gridConfig from '@/config/grid/qualificationGrid.ts';
import editorConfig from '@/config/editor/qualificationEditor.ts';

export default defineComponent({
    components: {
        GridEditor,
    },
    setup() {
        // hooks
        const { generateDataServices } = useService();

        // properties
        const dataServices = generateDataServices('Qualification');

        return {
            dataServices,
            gridConfig,
            editorConfig,
        };
    },
});
</script>
