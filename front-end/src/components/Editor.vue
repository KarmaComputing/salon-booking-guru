<template>
    <!-- editor -->
    <div v-if="!isLoading">
        <div class="space-y-4 w-full">
            <div
                class="flex flex-col w-full"
                v-for="(config, i) in editorConfig"
                :key="i"
            >
                <label v-if="config.title">{{ config.title }}</label>
                <component
                    v-model="reactiveModel[config.field]"
                    :is="typeComponentMap[config.type]"
                    :dropdown="config.dropdown"
                    :rowId="selectedRow.id"
                />
            </div>
        </div>
    </div>

    <!-- loading spinner -->
    <div class="flex justify-center">
        <ProgressSpinner v-if="isLoading" />
    </div>
</template>

<script lang="ts">
// vue
import { defineComponent, onMounted, ref } from 'vue';

// primevue
import ProgressSpinner from 'primevue/progressspinner';

// components
import Text from '@/components/editor/Text.vue';
import Number from '@/components/editor/Number.vue';
import Currency from '@/components/editor/Currency.vue';
import Dropdown from '@/components/editor/Dropdown.vue';
import AccountPassword from '@/components/editor/AccountPassword.vue';

export default defineComponent({
    components: {
        ProgressSpinner,
    },
    props: {
        selectedRow: {
            type: Object,
            default: null,
        },
        editorConfig: {
            type: Object,
            default: () => ({}),
        },
        getData: {
            type: Function,
            default: () => ({}),
        },
        createData: {
            type: Function,
            default: () => ({}),
        },
        updateData: {
            type: Function,
            default: () => ({}),
        },
    },
    setup(props: any) {
        // properties
        const typeComponentMap = {
            text: Text,
            number: Number,
            dropdown: Dropdown,
            currency: Currency,
            'account-password': AccountPassword,
        };

        // reactive
        const reactiveModel = ref<any>({});
        const isLoading = ref(true);

        // methods
        const save = async () => {
            if (props.selectedRow.id) {
                await props.updateData(reactiveModel.value);
            } else {
                await props.createData(reactiveModel.value);
            }
        };

        // lifecycle
        onMounted(async () => {
            if (props.selectedRow.id) {
                isLoading.value = true;
                reactiveModel.value = await props.getData(props.selectedRow.id);
                isLoading.value = false;
            } else {
                isLoading.value = false;
            }
        });

        return {
            reactiveModel,
            save,
            isLoading,
            typeComponentMap,
        };
    },
});
</script>
