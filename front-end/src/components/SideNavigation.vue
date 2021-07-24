<template>
    <div>
        <Button
            icon="pi pi-bars"
            class="p-button-link"
            style="color: white"
            @click="isNavVisible = true"
        />
        <Sidebar
            v-model:visible="isNavVisible"
            :baseZIndex="1000"
            position="right"
        >
            <div class="text-3xl">Menu</div>
            <Divider />
            <div class="flex flex-col">
                <RouterLink v-for="(link, i) in linkConfig" :key="i" to="/">
                    <Button
                        class="p-button-text w-full"
                        :icon="link.icon"
                        :label="link.label"
                        @click="isNavVisible = false"
                    />
                </RouterLink>
            </div>
        </Sidebar>
    </div>
</template>

<script lang="ts">
import { ref } from 'vue';
import Sidebar from 'primevue/sidebar';
import Button from 'primevue/button';
import Divider from 'primevue/divider';

export default {
    components: {
        Sidebar,
        Button,
        Divider,
    },
    setup() {
        const isNavVisible = ref(false);

        const linkConfig = [
            {
                icon: 'pi pi-user',
                label: 'Account',
            },
            {
                icon: 'pi pi-calendar',
                label: 'Availability',
            },
            {
                icon: 'pi pi-users',
                label: 'Accounts',
            },
            {
                icon: 'far fa-user-md',
                label: 'Treatments',
            },
            {
                icon: 'pi pi-calendar',
                label: 'Treatment category',
            },
            {
                icon: 'far fa-graduation-cap',
                label: 'Qualifications',
            },
            {
                icon: 'pi pi-cog',
                label: 'Settings',
            },
        ];

        return {
            isNavVisible,
            linkConfig,
        };
    },
};
</script>

<style>
.p-button-label {
    flex: none;
}
</style>
