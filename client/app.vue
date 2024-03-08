<script setup lang="ts">
useHead({
  meta: [
    { charset: "utf-8" },
    { name: "viewport", content: "width=device-width, initial-scale=1" },
    // { key: 'theme-color', name: 'theme-color', content: color }
  ],
  link: [
    { rel: "icon", href: "/favicon.ico" },
  ],
  htmlAttrs: {
    lang: "en",
  },
});

useSeoMeta({
  titleTemplate: "%s - Social Network",
  ogSiteName: "Social Network",
  twitterCard: "summary_large_image",
});


</script>

<template>
  <Html>
    <NuxtLoadingIndicator />
    <NuxtPage />
    <UNotifications />
  </Html>
</template>

<style lang="css"></style>
