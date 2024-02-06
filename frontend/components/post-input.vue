<template>
  <div class="bg-gray-100 dark:bg-gray-950 rounded-xl shadow-sm md:p-4 p-2 space-y-4 text-sm font-medium border1">

    <div class="flex items-center md:gap-3 gap-1">
      <div class="flex-1 hover:bg-opacity-80 transition-all rounded-lg cursor-pointer bg-gray-200 dark:bg-gray-900">
        <div @click="isOpen = true" class="py-2.5 text-center dark:text-white">What do you have in mind?</div>
      </div>
      <div
        class="cursor-pointer hover:bg-opacity-80 p-1 px-1.5 rounded-xl transition-all bg-pink-100/60 hover:bg-pink-100 dark:bg-white/10 dark:hover:bg-white/20">
        <svg xmlns="http://www.w3.org/2000/svg" class="w-8 h-8 stroke-pink-600 fill-pink-200/70" viewBox="0 0 24 24"
          stroke-width="1.5" stroke="#2c3e50" fill="none" stroke-linecap="round" stroke-linejoin="round">
          <path stroke="none" d="M0 0h24v24H0z" fill="none" />
          <path d="M15 8h.01" />
          <path d="M12 3c7.2 0 9 1.8 9 9s-1.8 9 -9 9s-9 -1.8 -9 -9s1.8 -9 9 -9z" />
          <path d="M3.5 15.5l4.5 -4.5c.928 -.893 2.072 -.893 3 0l5 5" />
          <path d="M14 14l1 -1c.928 -.893 2.072 -.893 3 0l2.5 2.5" />
        </svg>
      </div>
      <div
        class="cursor-pointer hover:bg-opacity-80 p-1 px-1.5 rounded-xl transition-all bg-sky-100/60 hover:bg-sky-100 dark:bg-white/10 dark:hover:bg-white/20">
        <svg xmlns="http://www.w3.org/2000/svg" class="w-8 h-8 stroke-sky-600 fill-sky-200/70 " viewBox="0 0 24 24"
          stroke-width="1.5" stroke="#2c3e50" fill="none" stroke-linecap="round" stroke-linejoin="round">
          <path stroke="none" d="M0 0h24v24H0z" fill="none" />
          <path d="M15 10l4.553 -2.276a1 1 0 0 1 1.447 .894v6.764a1 1 0 0 1 -1.447 .894l-4.553 -2.276v-4z" />
          <path d="M3 6m0 2a2 2 0 0 1 2 -2h8a2 2 0 0 1 2 2v8a2 2 0 0 1 -2 2h-8a2 2 0 0 1 -2 -2z" />
        </svg>
      </div>
    </div>
    <UModal v-model="isOpen">
      <div class="relative overflow-hidden flex justify-between items-center flex-wrap shadow-xl rounded-lg w-full bg-white dark:bg-gray-900">

        <div class="text-center p-4 border-b mb-0 dark:border-slate-700 flex w-full items-center justify-between">
          <h2 class="text-lg font-medium text-black  dark:text-white">Create Post</h2>
          <UButton variant="solid" color="gray" :ui="{ rounded: 'rounded-full' }" square @click="isOpen = false"
            icon="i-heroicons:x-mark"></UButton>
        </div>

        <textarea
          class="w-full m-4 p-2 !text-black placeholder:!text-black !bg-white !border-transparent focus:!border-transparent focus:!ring-transparent !font-normal !text-xl dark:!text-white dark:placeholder:!text-white dark:!bg-slate-800" rows="6" placeholder="What do you have in mind?"></textarea>

        <div class="p-5 flex w-full justify-between items-center">
          <USelectMenu color="primary" variant="outline" v-model="selected" :options="privacyLevel" />
          <div class="flex items-center gap-2">
            <UButton>Create</UButton>
          </div>
        </div>

      </div>


    </UModal>
  </div>
</template>

<script setup lang="ts">
const isOpen = ref(false)

const privacyLevel = [
  'public',
  'private',
  'almost private',
  'unlisted',
]

const selected = ref(privacyLevel[0])
</script>