<template>
  <UCard
    class="bg-gray-100 dark:bg-gray-950 w-48"
    :ui="{ body: { padding: 'p-2' }, header: { padding: 'p-0 h-38 rounded-t-lg overflow-hidden' } }"
  >
    <template #header>
      <img src="assets/images/avatars/avatar-2.jpg" class="w-full" alt="">
    </template>
    <div>
      <a href="timeline-event.html">
        <h4 class="card-title text-xl font-bold"> Delicious </h4>
      </a>
      <div class="text-sm my-1">
        16k following
      </div>
      <UButton block>
        Follow
      </UButton>
    </div>
  </UCard>
</template>