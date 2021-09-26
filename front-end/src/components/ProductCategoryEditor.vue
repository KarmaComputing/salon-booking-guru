<template>
    <div v-if="!isLoading">
        <div class="space-y-4 w-full">
            <div class="flex flex-col w-full">
                <label>Name</label>
                <InputText
                    class="w-full p-inputtext-sm"
                    type="text"
                    v-model="productCategory.name"
                />
            </div>
        </div>
    </div>
    <div class="flex justify-center">
        <ProgressSpinner v-if="isLoading" />
    </div>
</template>

<script lang="ts">
// vue
import { defineComponent, onMounted, ref } from 'vue';

import { ProductCategory } from '@/api/models';

// services
import { useService } from '@/api/services';

// primevue
import InputText from 'primevue/inputtext';
import ProgressSpinner from 'primevue/progressspinner';

export default defineComponent({
    components: {
        InputText,
        ProgressSpinner,
    },
    props: {
        productCategoryId: {
            type: Number,
            default: null,
        },
        roles: {
            type: Array,
            default: () => [],
        },
    },
    setup(props) {
        // hooks
        const {
            getProductCategory,
            createProductCategory,
            updateProductCategory,
        } = useService();

        // reactive
        const productCategory = ref({} as ProductCategory);
        const isLoading = ref(true);

        // methods
        const save = async () => {
            if (props.productCategoryId) {
                await updateProductCategory(productCategory.value);
            } else {
                await createProductCategory(productCategory.value);
            }
        };

        // lifecycle
        onMounted(async () => {
            if (props.productCategoryId) {
                isLoading.value = true;
                productCategory.value = await getProductCategory(
                    props.productCategoryId,
                );
                isLoading.value = false;
            } else {
                isLoading.value = false;
            }
        });

        return {
            productCategory,
            save,
            isLoading,
        };
    },
});
</script>
