<script lang="ts" setup>
import { Event } as '~/types'
const props = defineProps({
  event: {
    type: Event,
    required: true
  }
})
</script>

<template>
  <div>
    <div>
      {{ props.event.ID }}
      </div>
  </div>
</template>

<style></style>