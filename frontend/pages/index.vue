<template>
  <div class="lg:flex 2xl:gap-16 gap-12 max-w-[1065px]">
    <div class="max-w-[680px] mx-auto">
      <div class="md:max-w-[580px] mx-auto flex-1 xl:space-y-6 space-y-3">
        <PostInput />
        <UCarousel :items="items" :ui="{ item: 'basis-full md:basis-1/2 lg:basis-1/3' }" arrows>
          <UserCard class="mr-5" />
        </UCarousel>
        <UCarousel :items="items" :ui="{ item: 'basis-full md:basis-1/2 lg:basis-1/3' }" arrows>
          <GroupCard class="mr-5" />
        </UCarousel>
        <UCarousel :items="items" :ui="{ item: 'basis-full md:basis-1/2 lg:basis-1/3' }" arrows>
          <EventCard class="mr-5" />
        </UCarousel>
        <PostCard />
        <PostCard />
      </div>
    </div>
    <div class="flex-1" />
  </div>
</template>
<script lang="ts" setup>
const items = [
  'https://picsum.photos/600/600?random=1',
  'https://picsum.photos/600/600?random=2',
  'https://picsum.photos/600/600?random=3',
  'https://picsum.photos/600/600?random=4',
  'https://picsum.photos/600/600?random=5',
  'https://picsum.photos/600/600?random=6'
]

definePageMeta({
  alias: ['/'],
})

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