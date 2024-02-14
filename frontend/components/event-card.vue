<template>
  <UCard class="bg-gray-100 dark:bg-gray-950 w-72" :ui="{ header: { padding: 'p-0 rounded-t-lg overflow-hidden' } }">
    <template #header>
      <img src="../assets/images/post/img-2.jpg" class="" alt="">
    </template>
    <div>
      <p class="text-xs font-semibold text-teal-600 mb-1">
        Opening
      </p>
      <a href="timeline-event.html">
        <h4 class="card-title text-sm"> Perspective is everything </h4>
      </a>
      <a href="#">
        <p class="card-text text-black dark:text-white mt-2"> London </p>
      </a>
      <div class="flex mb-3 text-xs mt-1">
        <div> 20 Intersted</div>
        <div class="md:block hidden">
          ·
        </div>
        <div> 16 Going</div>
      </div>
      <div class="flex gap-2">
        <UButton>Intersted</UButton>
        <UButton icon="i-heroicons:share" color="gray" />
      </div>
    </div>
  </UCard>
</template>