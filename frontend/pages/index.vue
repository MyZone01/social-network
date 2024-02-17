<template>
  <NuxtLayout>
    <div class="max-w-[680px] mx-auto">
      <div class="md:max-w-[580px] mx-auto flex-1 xl:space-y-6 space-y-3">

        <!-- add story -->
        <div class="bg-white rounded-xl shadow-sm md:p-4 p-2 space-y-4 text-sm font-medium border1 dark:bg-dark2">

          <div class="flex items-center md:gap-3 gap-1">
            <div class="flex-1 bg-slate-100 hover:bg-opacity-80 transition-all rounded-lg cursor-pointer dark:bg-dark3"
              uk-toggle="target: #create-status">
              <div class="py-2.5 text-center dark:text-white"> What do you have in mind? </div>
            </div>
            <div
              class="cursor-pointer hover:bg-opacity-80 p-1 px-1.5 rounded-xl transition-all bg-pink-100/60 hover:bg-pink-100 dark:bg-white/10 dark:hover:bg-white/20"
              uk-toggle="target: #create-status">
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
              class="cursor-pointer hover:bg-opacity-80 p-1 px-1.5 rounded-xl transition-all bg-sky-100/60 hover:bg-sky-100 dark:bg-white/10 dark:hover:bg-white/20"
              uk-toggle="target: #create-status">
              <svg xmlns="http://www.w3.org/2000/svg" class="w-8 h-8 stroke-sky-600 fill-sky-200/70 " viewBox="0 0 24 24"
                stroke-width="1.5" stroke="#2c3e50" fill="none" stroke-linecap="round" stroke-linejoin="round">
                <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                <path d="M15 10l4.553 -2.276a1 1 0 0 1 1.447 .894v6.764a1 1 0 0 1 -1.447 .894l-4.553 -2.276v-4z" />
                <path d="M3 6m0 2a2 2 0 0 1 2 -2h8a2 2 0 0 1 2 2v8a2 2 0 0 1 -2 2h-8a2 2 0 0 1 -2 -2z" />
              </svg>
            </div>
          </div>

        </div>

        <!--  post image-->
        <post-card-img />

        <!--  post image with slider-->


        <!-- post text-->
        <post-card-text />

        <!-- placeholder -->
        <post-placeholder />

      </div>

    </div>
    <div class="hidden lg:p-20 uk- open" id="create-status" uk-modal="">

      <div
        class="uk-modal-dialog tt relative overflow-hidden mx-auto bg-white shadow-xl rounded-lg md:w-[520px] w-full dark:bg-dark2">

        <div class="text-center py-4 border-b mb-0 dark:border-slate-700">
          <h2 class="text-sm font-medium text-black"> Create Status </h2>

          <!-- close button -->
          <button type="button" class="button-icon absolute top-0 right-0 m-2.5 uk-modal-close">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
              stroke="currentColor" class="w-6 h-6">
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>

        </div>

        <div class="space-y-5 mt-3 p-2">
          <textarea
            class="w-full !text-black placeholder:!text-black !bg-white !border-transparent focus:!border-transparent focus:!ring-transparent !font-normal !text-xl   dark:!text-white dark:placeholder:!text-white dark:!bg-slate-800"
            name="" id="" rows="6" placeholder="What do you have in mind?"></textarea>
        </div>

        <div class="flex items-center gap-2 text-sm py-2 px-4 font-medium flex-wrap">
          <button type="button"
            class="flex items-center gap-1.5 bg-sky-50 text-sky-600 rounded-full py-1 px-2 border-2 border-sky-100 dark:bg-sky-950 dark:border-sky-900">
            <ion-icon name="image" class="text-base"></ion-icon>
            Image
          </button>
          <button type="button"
            class="flex items-center gap-1.5 bg-teal-50 text-teal-600 rounded-full py-1 px-2 border-2 border-teal-100 dark:bg-teal-950 dark:border-teal-900">
            <ion-icon name="videocam" class="text-base"></ion-icon>
            Video
          </button>
          <button type="button"
            class="flex items-center gap-1.5 bg-orange-50 text-orange-600 rounded-full py-1 px-2 border-2 border-orange-100 dark:bg-yellow-950 dark:border-yellow-900">
            <ion-icon name="happy" class="text-base"></ion-icon>
            Feeling
          </button>
          <button type="button"
            class="flex items-center gap-1.5 bg-red-50 text-red-600 rounded-full py-1 px-2 border-2 border-rose-100 dark:bg-rose-950 dark:border-rose-900">
            <ion-icon name="location" class="text-base"></ion-icon>
            Check in
          </button>
          <button type="button" class="grid place-items-center w-8 h-8 text-xl rounded-full bg-secondery">
            <ion-icon name="ellipsis-horizontal"></ion-icon>
          </button>
        </div>

        <div class="p-5 flex justify-between items-center">
          <div>
            <button
              class="inline-flex items-center py-1 px-2.5 gap-1 font-medium text-sm rounded-full bg-slate-50 border-2 border-slate-100 group aria-expanded:bg-slate-100 aria-expanded: dark:text-white dark:bg-slate-700 dark:border-slate-600"
              type="button">
              Everyone
              <ion-icon name="chevron-down-outline"
                class="text-base duration-500 group-aria-expanded:rotate-180"></ion-icon>
            </button>

            <div
              class="p-2 bg-white rounded-lg shadow-lg text-black font-medium border border-slate-100 w-60 dark:bg-slate-700"
              uk-drop="offset:10;pos: bottom-left; reveal-left;animate-out: true; animation: uk-animation-scale-up uk-transform-origin-bottom-left ; mode:click">

              <form>
                <label>
                  <input type="radio" name="radio-status" id="monthly1" class="peer appearance-none hidden" checked />
                  <div
                    class=" relative flex items-center justify-between cursor-pointer rounded-md p-2 px-3 hover:bg-secondery peer-checked:[&_.active]:block dark:bg-dark3">
                    <div class="text-sm"> Everyone </div>
                    <ion-icon name="checkmark-circle"
                      class="hidden active absolute -translate-y-1/2 right-2 text-2xl text-blue-600 uk-animation-scale-up"></ion-icon>
                  </div>
                </label>
                <label>
                  <input type="radio" name="radio-status" id="monthly1" class="peer appearance-none hidden" />
                  <div
                    class=" relative flex items-center justify-between cursor-pointer rounded-md p-2 px-3 hover:bg-secondery peer-checked:[&_.active]:block dark:bg-dark3">
                    <div class="text-sm"> Friends </div>
                    <ion-icon name="checkmark-circle"
                      class="hidden active absolute -translate-y-1/2 right-2 text-2xl text-blue-600 uk-animation-scale-up"></ion-icon>
                  </div>
                </label>
                <label>
                  <input type="radio" name="radio-status" id="monthly" class="peer appearance-none hidden" />
                  <div
                    class=" relative flex items-center justify-between cursor-pointer rounded-md p-2 px-3 hover:bg-secondery peer-checked:[&_.active]:block dark:bg-dark3">
                    <div class="text-sm"> Only me </div>
                    <ion-icon name="checkmark-circle"
                      class="hidden active absolute -translate-y-1/2 right-2 text-2xl text-blue-600 uk-animation-scale-up"></ion-icon>
                  </div>
                </label>
              </form>

            </div>
          </div>
          <div class="flex items-center gap-2">
            <button type="button" class="button bg-blue-500 text-white py-2 px-12 text-[14px]"> Create</button>
          </div>
        </div>

      </div>

    </div>
    
    <post-preview />
  </NuxtLayout>
</template>
<script lang="ts" setup>


</script>