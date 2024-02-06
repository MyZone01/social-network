<template>
  <div class="lg:flex 2xl:gap-16 gap-12 max-w-[1065px]">
    <div class="max-w-[680px] mx-auto">
      <div class="md:max-w-[580px] mx-auto flex-1 xl:space-y-6 space-y-3">
        <PostInput />
        <PostCard />
        <PostCard />
        <PostCard />
        <PostCard />
        <PostCard />
      </div>
    </div>
    <div class="flex-1"></div>
  </div>
</template>
<script lang="ts" setup>
definePageMeta({
  alias: ["/"],
});

useHead({
  title: 'Home',
})

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