<template>
  <NuxtLayout>  

  </NuxtLayout>
</template>

<script setup lang="ts">
const currentUser = useAuthUser();



definePageMeta({
  alias: ["/"],
  middleware: ["auth-only"],
});
</script>