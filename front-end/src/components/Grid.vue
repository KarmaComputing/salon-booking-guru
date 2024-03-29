<template>
    <DataTable
        class="rounded overflow-hidden"
        :value="gridData"
        responsiveLayout="stack"
        :loading="isLoading"
    >
        <Column
            v-for="(config, i) in gridConfig"
            :field="config.field"
            :key="i"
            :header="config.title"
        >
            <template #body="slotProps">
                <!-- plain -->
                <span v-if="!config.cellRenderer">
                    {{ slotProps.data[config.field] }}
                </span>

                <!-- cell renderer -->
                <span v-if="config.cellRenderer">
                    <component
                        :is="config.cellRenderer"
                        :data="slotProps.data"
                        :field="config.field"
                    >
                        {{ slotProps.data[config.field] }}
                    </component>
                </span>
            </template>
        </Column>

        <Column v-if="actionButtonConfig" header="Actions">
            <template #body="slotProps">
                <div class="flex space-x-2">
                    <div v-for="(config, i) in actionButtonConfig" :key="i">
                        <Button
                            v-if="!config.route"
                            class="p-button-rounded p-button-text"
                            :icon="config.icon"
                            :class="config.style"
                            @click="triggerAction(config.callback, slotProps)"
                        />
                        <RouterLink v-if="config.route" :to="config.route">
                            <Button
                                class="p-button-rounded p-button-text"
                                :icon="config.icon"
                                :class="config.style"
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
import { defineComponent, ref } from 'vue';

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
        isLoading: {
            type: Boolean,
            default: false,
        },
    },
    components: {
        DataTable,
        Column,
    },
    setup() {
        // reactive
        const selectedRow = ref<any>(null);

        // methods
        const triggerAction = (callback: any, slotProps: any) => {
            selectedRow.value = slotProps.data;
            callback();
        };

        return {
            triggerAction,
            selectedRow,
        };
    },
});
</script>
