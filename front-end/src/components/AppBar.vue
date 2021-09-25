<template>
    <div
        class="p-shadow-4 text-white z-10 mb-6 sticky top-0 left-0 px-6 w-full"
        style="background: #3f51b5"
    >
        <div class="container mx-auto flex justify-between items-center">
            <RouterLink class="text-2xl py-4" to="/">
                Salon Booking Guru
            </RouterLink>
            <SideNavigation v-if="store.getters.authToken" />
        </div>
    </div>
</template>

<script lang="ts">
// vue
import { defineComponent } from 'vue';
import { useStore } from 'vuex';

// components
import SideNavigation from '@/components/SideNavigation.vue';

export default defineComponent({
    components: {
        SideNavigation,
    },
    setup() {
        const store = useStore();

        return {
            store,
        };
    },
});
</script>
