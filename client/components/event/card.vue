<script lang="ts" setup>
import { useDateFormat } from "@vueuse/core";
import type { PropType } from "vue";
import type { Event } from "~/types";

const props = defineProps({
  event: {
    type: Object as PropType<Event>,
    required: true,
  },
});

const joined = ref(false)
const date = useDateFormat(new Date(props.event.date_time),'DD,MMMM HH:MM')


</script>

<template>
  <UCard class="w-full flex flex-row justify-between">
    <div>
      <div class="text-lg text-blue-700">
        {{ props.event.title }}
      </div>
      <div>
        {{ props.event.description }}
      </div>
    </div>
    <div>
      <div>{{ date }}</div>
    </div>
    <UButton v-if="!joined" class="bg-blue-500 ">Join</UButton>
    <div v-else class="text-blue-500 ">joined</div>

  </UCard>
</template>

<style></style>