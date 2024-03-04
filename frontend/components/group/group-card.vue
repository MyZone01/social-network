<script setup>
const props = defineProps(['group'])
</script>

<template>
  <UCard class="card w-60 bg-gray-100 dark:bg-gray-950  p-0"
    :ui="{ header: { padding: 'p-0 rounded-t-lg overflow-hidden' } }">
    <NuxtLink :href="'/groups/' + props.group.ID">
      <div class="card-media h-24 rounded-t-lg">
        <img src="assets/images/post/img-2.jpg" class="" alt="">
      </div>
    </NuxtLink>
    <div class="card-body z-10 relative w-full">
      <NuxtLink :href="'/groups/' + props.group.ID">
        <h4 class="card-title text-xl text-black font-bold text-ellipsis w-full">
          {{ props.group.Title }}
        </h4>
      </NuxtLink>
      <div class="flex mb-4 text-sm mt-2">
        <div class="md:block hidden" />
        <div> 16k members </div>
      </div>
      <div class="flex flex-row gap-2 justify-center w-full">
        <UButton class="items-center bg-blue-500 justify-center flex-1">
          Join
        </UButton>
        <UButton class="items-center justify-center" color="gray">
          View
        </UButton>
      </div>
    </div>
  </UCard>
</template>