<template>
    <AppBar />
    <div class="px-6">
        <div class="container mx-auto">
            <RouterView />
        </div>
    </div>
</template>

<script lang="ts">
// vue
import { defineComponent } from 'vue';

// components
import AppBar from './components/AppBar.vue';

export default defineComponent({
    components: {
        AppBar,
    },
});
</script>

<style scoped>
@import './styles/index.css';
</style>
