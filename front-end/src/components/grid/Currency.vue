<template>
    <span>{{ currency }}</span>
</template>

<script lang="ts">
// vue
import { defineComponent, computed } from 'vue';

export default {
    props: {
        data: {
            type: Object,
            default: null,
        },
        field: {
            type: String,
            default: '',
        },
    },
    setup(props: any) {
        // computed
        const currency = computed(() => {
            return props.data[props.field].toLocaleString('en-GB', {
                style: 'currency',
                currency: 'GBP',
            });
        });

        return {
            currency,
        };
    },
};
</script>
