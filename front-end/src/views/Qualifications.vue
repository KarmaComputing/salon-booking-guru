<template>
    <GridEditor
        title="Qualifications"
        addButtonLabel="ADD QUALIFICATION"
        editorHeader="Qualification Editor"
        rowName="qualification"
        deleteRowValue="name"
        :gridConfig="gridConfig"
        :getData="getAllQualification"
        :deleteData="deleteQualification"
        :editorComponent="QualificationEditor"
    />
</template>

<script lang="ts">
// vue
import { defineComponent } from 'vue';

// components
import GridEditor from '@/components/GridEditor.vue';
import QualificationEditor from '@/components/QualificationEditor.vue';

// services
import { useService } from '@/api/services';

// configs
import gridConfig from '@/config/grid/qualificationGrid.ts';

export default defineComponent({
    components: {
        GridEditor,
    },
    setup() {
        // hooks
        const { getAllQualification, deleteQualification } = useService();

        return {
            getAllQualification,
            deleteQualification,
            gridConfig,
            QualificationEditor,
        };
    },
});
</script>
