<template>
  <h1 class="text-6xl">Welcome to Social Network</h1>
  <p>
    {{ currentUser }}
  </p>
</template>

<script setup lang="ts">
const currentUser = useAuthUser();

useHead({
  title: "Home",
})

definePageMeta({
  alias: ["/"],
  middleware: ["auth-only"],
});
</script>