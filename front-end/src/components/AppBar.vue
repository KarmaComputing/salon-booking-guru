<template>
    <div
        class="
            p-shadow-4
            bg-blue-600
            text-white
            z-10
            mb-8
            sticky
            top-0
            left-0
            px-6
            w-full
        "
    >
        <div class="container mx-auto flex justify-between items-center">
            <RouterLink class="text-2xl py-4" to="/">
                Salon Booking Guru
            </RouterLink>
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
