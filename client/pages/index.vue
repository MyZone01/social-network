<template>
  <h1 class="text-6xl">Welcome to Social Network</h1>
  <img w-64 :src="'http://localhost:8081/' + currentUser?.avatarImage" />
  <p>
    {{ currentUser }}
  </p>
  <nuxt-link to="/home">Home</nuxt-link>
  <nuxt-link to="/groups">Groups</nuxt-link>
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