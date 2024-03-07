<template>
  <h1>Homess</h1>
  <nuxt-link to="/">Back to home</nuxt-link>
</template>

<script setup lang="ts">
definePageMeta({
  alias: ["/home"],
  middleware: ["auth-only"],
});
</script>