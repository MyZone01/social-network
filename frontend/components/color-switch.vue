<script lang="ts" setup>
const colorMode = useColorMode()

function toggleMode() {
  colorMode.preference = colorMode.value === 'light' ? 'dark' : 'light'
}
</script>
   
<template>
  <UButton
    :icon="colorMode.value === 'dark' ? 'i-heroicons-moon' : 'i-heroicons-sun'" variant="ghost"
    @click="toggleMode"
  />
</template>