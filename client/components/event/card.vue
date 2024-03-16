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

const status = ref("not going")
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
    <div class="flex gap-1 flex-row">
      <UButton v-if="!(status === 'going')" class="bg-blue-500 ">going</UButton>
      <div v-else class="text-blue-500 ">going</div>
      <UButton v-if="(status === 'not going')" class="bg-blue-500 ">not going</UButton>
      <div v-else class="text-blue-500 ">not going</div>
    </div>


  </UCard>
</template>

<style></style>