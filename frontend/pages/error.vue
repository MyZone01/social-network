<script setup lang="ts">
import type { NuxtError } from '#app'
import type { Html } from '#build/components'

useSeoMeta({
  title: 'Page not found',
  description: 'We are sorry but this page could not be found.'
})

defineProps({
  error: {
    type: Object as PropType<NuxtError>,
    required: true
  }
})

useHead({
  htmlAttrs: {
    lang: 'en'
  }
})
</script>

<template>
  <Html>
    <Body>
      {{ error.statusCode }}
      {{ error.message }}
    </Body>
    <UNotifications />
  </Html>
</template>