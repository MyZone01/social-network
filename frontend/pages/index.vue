<template>
  <h1>Home</h1>
</template>
<script lang="ts" setup>
useSeoMeta({
  titleTemplate: '',
  title: 'Home',
  ogTitle: 'Home',
  description: 'Description',
  ogDescription: 'Description'
})

defineOgImageComponent('NuxtSeo', {
  title: 'Hello OG Image 👋',
  description: 'Look what at me in dark mode',
  theme: '#ff0000',
  colorMode: 'dark',
})
</script>