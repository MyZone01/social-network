<template>
  <main id="site__main" class="2xl:ml-[--w-side] xl:ml-[--w-side-sm] p-2.5 h-[calc(100vh-var(--m-top))] mt-[--m-top]">
    <div class="max-w-[1065] mx-auto bg-slate-500">
      <div class="bg-white shadow lg:rounded-b-2xl lg:-mt-10 dark:bg-dark2 pt-2">
        <div class="relative overflow-hidden w-full lg:h-72 h-32">
          <img src="assets/images/post/img-2.jpg" alt="" class="h-full w-full object-cover inset-0" />
          <div class="w-full bottom-0 absolute left-0 bg-gradient-to-t from-black/60 pt-10 z-10"></div>
          <div class="absolute bottom-0 right-0 m-a z-20">
            <div class="flex items-center gap-3 px-2">
              <UButton class="button bg-white/20 text-white flex items-center gap-2 back-drop-blur-small">Crop</UButton>
              <UButton class="button bg-white/10 text-white flex items-center gap-2 back-drop-blur-small">Edit</UButton>
            </div>
          </div>
        </div>
        <div class="lg:px-10 md:p-5 p-3">
          <div class="flex flex-col justify-center">
            <div class="flex lg:items-center justify-between max-md:flex-col">
              <div class="flex-1">
                <h3 class="md:text-2xl font-bold">Group Name</h3>
                <p class="font-normal text-gray-500 mt-2 flex gap-2 flex-wrap dark:text-white/80">
                  <span class="max-lg:hidden"> Public group </span>
                  <span class="max-lg:hidden"> • </span>
                  <span>
                    <b class="font-medium text-black dark:text-white">1.2K</b>
                    members
                  </span>
                </p>
              </div>

              <div>
                <div class="flex items-center gap-2 mt-1">
                  <div class="flex -space-x-4 mr-3">
                    <img src="assets/images/avatars/avatar-2.jpg" alt=""
                      class="w-10 rounded-full border-4 border-white dark:border-slate-800" />
                    <img src="assets/images/avatars/avatar-3.jpg" alt=""
                      class="w-10 rounded-full border-4 border-white dark:border-slate-800" />
                    <img src="assets/images/avatars/avatar-7.jpg" alt=""
                      class="w-10 rounded-full border-4 border-white dark:border-slate-800" />
                    <img src="assets/images/avatars/avatar-4.jpg" alt=""
                      class="w-10 rounded-full border-4 border-white dark:border-slate-800" />
                    <img src="assets/images/avatars/avatar-5.jpg" alt=""
                      class="w-10 rounded-full border-4 border-white dark:border-slate-800" />
                  </div>
                  <div>
                    <button type="button" class="rounded-lg bg-secondery flex px-2.5 py-2 dark:bg-dark2">
                      <i class='bx bx-dots-horizontal-rounded text-xl' ></i>
                    </button>
                    <div class="w-[100px]"
                      uk-dropdown="pos: bottom-right; animation: uk-animation-scale-up uk-transform-origin-top-right; animate-out: true; mode: click;offset:10">
                      <nav>
                        <a href="#">
                          <i class='bx bx-link text-xl' ></i>
                          Copy link
                        </a>
                        <a href="#" class="text-red-400 hover:!bg-red-50 dark:hover:!bg-red-500/50">
                          <i class='bx bx-stop-circle text-xl' ></i>
                          Block
                        </a>
                      </nav>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="flex items-center justify-between border-t border-gray-100 px-2 dark:border-slate-700">
          <nav
            class="flex gap-0.5 rounded-xl overflow-hidden -mb-px text-gray-500 font-medium text-sm overflow-x-auto dark:text-white">
            <a href="#" class="inline-block py-3 leading-8 px-3.5 border-b-2 border-blue-600 text-blue-600">Posts</a>
            <a href="#" class="inline-block py-3 leading-8 px-3.5">Events</a>
            <a href="#" class="inline-block py-3 leading-8 px-3.5">Members</a>
            <a href="#" class="inline-block py-3 leading-8 px-3.5">Media</a>
          </nav>

          <div class="flex items-center  gap-1 text-sm p-3 bg-blue py-2 mr-2 rounded-xl max-md:hidden dark:bg-white/5">
            <i class='bx bx-search text-lg' ></i>
            <input placeholder="Search .."
              class="!bg-transparent outline-none focus:outline-none focus:border-b-blue-500" />
          </div>
        </div>
      </div>
      <div class="flex 2xl:gap-12 gap-10 mt-8 max-lg:flex-col" id="js-oversized">
        <!-- add story -->
        <div class="bg-white rounded-xl shadow-sm p-4 space-y-4 text-sm font-medium border1 dark:bg-dark2">
          <div class="flex items-center gap-3">
            <div class="flex-1 bg-slate-100 hover:bg-opacity-80 transition-all rounded-lg cursor-pointer dark:bg-dark3"
              uk-toggle="target: #create-status">
              <div class="py-2.5 text-center dark:text-white">
                What do you have in mind?
              </div>
            </div>
            <div
              class="cursor-pointer hover:bg-opacity-80 p-1 px-1.5 rounded-lg transition-all bg-pink-100/60 hover:bg-pink-100"
              uk-toggle="target: #create-status">
              <svg xmlns="http://www.w3.org/2000/svg" class="w-8 h-8 stroke-pink-600 fill-pink-200/70"
                viewBox="0 0 24 24" stroke-width="1.5" stroke="#2c3e50" fill="none" stroke-linecap="round"
                stroke-linejoin="round">
                <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                <path d="M15 8h.01" />
                <path d="M12 3c7.2 0 9 1.8 9 9s-1.8 9 -9 9s-9 -1.8 -9 -9s1.8 -9 9 -9z" />
                <path d="M3.5 15.5l4.5 -4.5c.928 -.893 2.072 -.893 3 0l5 5" />
                <path d="M14 14l1 -1c.928 -.893 2.072 -.893 3 0l2.5 2.5" />
              </svg>
            </div>
            <div
              class="cursor-pointer hover:bg-opacity-80 p-1 px-1.5 rounded-lg transition-all bg-sky-100/60 hover:bg-sky-100"
              uk-toggle="target: #create-status">
              <svg xmlns="http://www.w3.org/2000/svg" class="w-8 h-8 stroke-sky-600 fill-sky-200/70" viewBox="0 0 24 24"
                stroke-width="1.5" stroke="#2c3e50" fill="none" stroke-linecap="round" stroke-linejoin="round">
                <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                <path d="M15 10l4.553 -2.276a1 1 0 0 1 1.447 .894v6.764a1 1 0 0 1 -1.447 .894l-4.553 -2.276v-4z" />
                <path d="M3 6m0 2a2 2 0 0 1 2 -2h8a2 2 0 0 1 2 2v8a2 2 0 0 1 -2 2h-8a2 2 0 0 1 -2 -2z" />
              </svg>
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>
