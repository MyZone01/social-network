<script lang="ts" setup>
const colorMode = useColorMode()

function toggleMode() {
  colorMode.preference = colorMode.value === 'light' ? 'dark' : 'light'
}
</script>

<template>
  <UButton
    variant="solid" color="gray" :ui="{ rounded: 'rounded-full' }" square size="lg"
    :icon="colorMode.value === 'dark' ? 'i-heroicons-moon' : 'i-heroicons-sun'" @click="toggleMode"
  />
</template>