<template>
  <UCard class="bg-gray-100 dark:bg-gray-950 w-72" :ui="{ header: { padding: 'p-0 rounded-t-lg overflow-hidden' } }">
    <template #header>
      <img src="assets/images/post/img-2.jpg" class="" alt="">
    </template>
    <div class="relative">
      <UAvatar size="sm" src="https://avatars.githubusercontent.com/u/739984?v=4" alt="Avatar" class="absolute -top-9" />
      <a href="timeline-event.html">
        <h4 class="card-title text-2xl font-bold"> Delicious Foods </h4>
      </a>
      <div class="flex mb-4 text-sm mt-2">
        <div>Travel</div>
        <div class="md:block hidden">
          ·
        </div>
        <div> 16k members</div>
      </div>
      <div class="flex gap-2">
        <UButton class="items-center justify-center w-10/12">
          Join
        </UButton>
        <UButton class="items-center justify-center w-2/12" color="gray">
          View
        </UButton>
      </div>
    </div>
  </UCard>
</template>