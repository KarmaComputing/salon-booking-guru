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
                        <i :class="link.icon" class="mr-3 w-5" />
                        <span class="">{{ link.label }}</span>
                    </Button>
                </RouterLink>
                <Divider />
                <Button class="p-button-text w-full" @click="logOutClose">
                    <i class="far fa-sign-out mr-3 w-5" />
                    <span>Log out</span>
                </Button>
            </div>
        </Sidebar>
    </div>
</template>

<script lang="ts">
// vue
import { ref } from 'vue';

// services
import { useService } from '@/api/services';

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
        // hooks
        const { logOut } = useService();

        // properties
        const links = navLinks;

        // reactive
        const isNavVisible = ref(false);

        // methods
        const logOutClose = () => {
            logOut();
            isNavVisible.value = false;
        };

        return {
            isNavVisible,
            links,
            logOutClose,
        };
    },
};
</script>
