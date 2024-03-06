<template>
  <h1 class="text-6xl">Welcome to Social Network</h1>
  <nuxt-img :src="'http://localhost:8081/' + currentUser?.avatarImage" />
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