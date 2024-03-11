<template>
</template>

<script setup lang="ts">
definePageMeta({
  alias: ["/"],
  middleware: ["auth-only"],
});
</script>