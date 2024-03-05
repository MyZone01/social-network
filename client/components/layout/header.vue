<template>
  <header>
    <h1>Header</h1>
    <color-switch />
    <UButton variant="outline" color="gray" class="flex items-center justify-center w-full" mb-2 rounded-lg
      @click="onLogoutClick">
      <span class="mx-2 font-medium">Logout</span>
      <i class="i-heroicons:arrow-right-on-rectangle-20-solid text-2xl"></i>
    </UButton>
  </header>
</template>

<script setup lang="ts">
const currentUser = useAuthUser();
const loading = ref(false);
const { logout, me } = useAuth();

onMounted(async () => {
  if (!currentUser.value?.firstName) {
    await me();
  }
});

async function onLogoutClick() {
  try {
    loading.value = true;

    await logout();

    navigateTo("/login");
  } catch (error) {
    console.error(error);
  } finally {
    loading.value = false;
  }
}
</script>