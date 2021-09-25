<template>
    <DataTable :value="gridData" responsiveLayout="stack">
        <Column
            v-for="(config, i) in gridConfig"
            :field="config.field"
            :key="i"
            :header="config.title"
        />
        <Column v-if="actionButtonConfig" header="Actions">
            <template #body="slotProps">
                <div class="flex space-x-2">
                    <div v-for="(config, i) in actionButtonConfig" :key="i">
                        <Button
                            v-if="!config.route"
                            class="p-button-rounded"
                            :icon="config.icon"
                            :class="config.style"
                            @click="config.callback"
                        />
                        <RouterLink v-if="config.route" :to="config.route">
                            <Button
                                class="p-button-rounded"
                                :icon="config.icon"
                                :class="config.style"
                                @click="config.callback(slotProps.data)"
                            />
                        </RouterLink>
                    </div>
                </div>
            </template>
        </Column>
    </DataTable>
</template>

<script lang="ts">
// vue
import { defineComponent } from 'vue';

// primevue
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';

export default defineComponent({
    props: {
        // the data that will be displayed on the grid
        data: {
            type: Object,
        },
        // grid configuration data to determine column names, field names etc
        gridConfig: {
            type: Object,
            required: true,
        },
        // action button configuration data to dermine the buttons and callbacks
        actionButtonConfig: {
            type: Object,
            default: null,
        },
        gridData: {
            type: Array,
            default: () => [],
        },
    },
    components: {
        DataTable,
        Column,
    },
    setup() {
        return {};
    },
});
</script>
