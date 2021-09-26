<template>
    <span>{{ data[field] }} hours</span>
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
        return {};
    },
};
</script>
