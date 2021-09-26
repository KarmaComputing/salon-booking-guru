<template>
    <div v-if="!isLoading">
        <div class="space-y-4 w-full">
            <div class="flex flex-col w-full">
                <label>Name</label>
                <InputText
                    class="w-full p-inputtext-sm"
                    type="text"
                    v-model="qualification.name"
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

import { Qualification } from '@/api/models';

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
        const { getQualification, updateQualification, createQualification } =
            useService();

        // reactive
        const qualification = ref({} as Qualification);
        const isChangePassword = ref(false);
        const isLoading = ref(true);

        // methods
        const save = async () => {
            if (props.id) {
                await updateQualification(qualification.value);
            } else {
                await createQualification(qualification.value);
            }
        };

        // lifecycle
        onMounted(async () => {
            if (props.id) {
                isLoading.value = true;
                qualification.value = await getQualification(props.id);
                isLoading.value = false;
            } else {
                isLoading.value = false;
            }
        });

        return {
            qualification,
            save,
            isLoading,
        };
    },
});
</script>
