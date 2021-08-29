<template>
    <div>
        <Button
            class="p-button-link"
            icon="pi pi-bars"
            style="color: white"
            @click="isNavVisible = true"
        />
        <Sidebar v-model:visible="isNavVisible" position="right">
            <div class="text-3xl">Menu</div>
            <Divider />
            <div class="flex flex-col">
                <RouterLink
                    v-for="(link, i) in links"
                    :key="i"
                    :to="{ path: link.route }"
                >
                    <Button
                        class="p-button-text w-full"
                        @click="isNavVisible = false"
                    >
                        <i :class="link.icon" class="mr-3 w-5 h-5" />
                        <span class="">{{ link.label }}</span>
                    </Button>
                </RouterLink>
            </div>
        </Sidebar>
    </div>
</template>

<script lang="ts">
// vue
import { ref } from 'vue';

// primevue
import Sidebar from 'primevue/sidebar';
import Button from 'primevue/button';
import Divider from 'primevue/divider';

// config
import { navLinks } from '../config/navLinks';

export default {
    components: {
        Sidebar,
        Button,
        Divider,
    },
    setup() {
        const isNavVisible = ref(false);
        const links = navLinks;

        return {
            isNavVisible,
            links,
        };
    },
};
</script>
