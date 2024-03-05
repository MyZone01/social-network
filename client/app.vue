<script setup lang="ts">
useHead({
  title: 'Home',
  titleTemplate: `%s | ${'Social Network'}`,
  meta: [
    { charset: 'utf-8' },
    { name: 'viewport', content: 'width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=0' },
    { hid: 'description', name: 'description', content: 'An interactive programming education platform with embedded coding environments and immersive, hands-on videos.' },
  ],
  link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }],
})
</script>

<template>
  <Html scroll-smooth snap-mandatory snap-y>
    <Body w-screen h-screen flex overflow-hidden items-start>
      <NuxtLoadingIndicator />
      <NuxtLayout>
        <NuxtPage />
        <UNotifications max-w-md />
      </NuxtLayout>
    </Body>
  </Html>
</template>

<style lang="css">
body>#__nuxt {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  justify-content: flex-start;
}
</style>