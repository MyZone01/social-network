<script lang="ts" setup>
const colorMode = useColorMode()

// logic for authenticated user
const authenticated = false

const color = computed(() => colorMode.value === 'dark' ? '#111827' : 'white')

useHead({
  meta: [
    { charset: 'utf-8' },
    { name: 'viewport', content: 'width=device-width, initial-scale=1' },
    { key: 'theme-color', name: 'theme-color', content: color }
  ],
  link: [
    { rel: 'icon', href: '/favicon.ico' }
  ],
  htmlAttrs: {
    lang: 'en'
  }
})

useSeoMeta({
  titleTemplate: '%s - Social Network',
  ogSiteName: 'Social Network',
  twitterCard: 'summary_large_image'
})

onMounted(() => {
  // On page load or when changing themes, best to add inline in `head` to avoid FOUC
  if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
    document.documentElement.classList.add('dark')
  } else {
    document.documentElement.classList.remove('dark')
  }

  // Whenever the user explicitly chooses light mode
  localStorage.theme = 'light'

  // Whenever the user explicitly chooses dark mode
  localStorage.theme = 'dark'

  // Whenever the user explicitly chooses to respect the OS preference
  localStorage.removeItem('theme')
})
</script>
<template>
  <Html>
    <NuxtLoadingIndicator />
    <NuxtPage />
    <UNotifications />
  </Html>
</template><template>
  <Html>
    <NuxtLoadingIndicator />
    <NuxtPage />
    <UNotifications />
  </Html>
</template>
