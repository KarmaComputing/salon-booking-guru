<template>
    <div v-if="!isLoading">
        <div class="space-y-4 w-full">
            <div class="flex flex-col w-full">
                <label>Name</label>
                <InputText
                    class="w-full p-inputtext-sm"
                    type="text"
                    v-model="product.name"
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

import { Product } from '@/api/models';

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
        id: {
            type: Number,
            default: null,
        },
    },
    setup(props) {
        // hooks
        const { getProduct, updateProduct, createProduct } = useService();

        // reactive
        const product = ref({} as Product);
        const isChangePassword = ref(false);
        const isLoading = ref(true);

        // methods
        const save = async () => {
            if (props.id) {
                await updateProduct(product.value);
            } else {
                await createProduct(product.value);
            }
        };

        // lifecycle
        onMounted(async () => {
            if (props.id) {
                isLoading.value = true;
                product.value = await getProduct(props.id);
                isLoading.value = false;
            } else {
                isLoading.value = false;
            }
        });

        return {
            product,
            save,
            isLoading,
        };
    },
});
</script>
