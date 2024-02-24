<script>
  const props = defineProps(['group'])
</script>

<template>
  <UCard class="card bg-gray-100 dark:bg-gray-950  p-0 flex-1" :ui="{ header: { padding: 'p-0 rounded-t-lg overflow-hidden' } }">
    <NuxtLink to="/groups/id">
      <div class="card-media h-24 rounded-t-lg">
        <img src="assets/images/post/img-2.jpg" class="" alt="">
      </div>
    </NuxtLink>
    <div class="card-body z-10 relative w-full">
      <UAvatar size="sm" src="https://avatars.githubusercontent.com/u/739984?v=4" alt="Avatar"
        class="w-10 rounded-full mb-2 shadow -mt-8 relative border-2 border-white dark:border-slate-800" />
      <NuxtLink to="/groups/id">
        <h4 class="card-title text-xl font-bold"> props.group.Title</h4>
      </NuxtLink>
      <div class="flex mb-4 text-sm mt-2">
        <div>Travel</div>
        <div class="md:block hidden">
        </div>
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