<template>
    <div class="p-shadow-4 bg-blue-600 text-white">
        <div class="container mx-auto px-6 flex justify-between items-center">
            <RouterLink class="text-2xl py-4" to="/"
                >Salon Booking Guru</RouterLink
            >
            <SideNavigation />
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import SideNavigation from '../components/SideNavigation.vue';

export default defineComponent({
    components: {
        SideNavigation,
    },
});
</script>
