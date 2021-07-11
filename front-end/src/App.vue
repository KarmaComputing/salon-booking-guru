<template>
    <AppBar />
    <RouterView />
</template>

<script lang="ts">
// vue
import { defineComponent } from 'vue';

// components
import AppBar from '@/components/AppBar.vue';

export default defineComponent({
    components: {
        AppBar,
    },
});
</script>

<style scoped>
@import './styles/index.css';
</style>
