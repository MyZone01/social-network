<template>
  <header
    class="z-[100] h-[--m-top] fixed top-0 left-0 w-full flex items-center bg-white/80 sky-50 backdrop-blur-xl border-b border-gray-200 dark:bg-gray-950 dark:border-gray-800">

    <div class="flex items-center w-full xl:px-6 px-2 max-lg:gap-10">

      <div class="2xl:w-[--w-side] lg:w-[--w-side-sm]">

        <!-- left -->
        <div class="flex items-center gap-1">

          <!-- icon menu -->
          <UButton class="xl:hidden" square size="lg" color="gray" :ui="{ rounded: 'rounded-full' }"
            icon="i-heroicons:bars-3-solid">
          </UButton>
          <div id="logo">
            <a href="feed.html">
              <img src="assets/images/logo.png" alt="" class="w-28 md:block hidden dark:!hidden">
              <img src="assets/images/logo-light.png" alt="" class="dark:md:block hidden">
              <img src="assets/images/logo-mobile.png" class="hidden max-md:block w-20 dark:!hidden" alt="">
              <img src="assets/images/logo-mobile-light.png" class="hidden dark:max-md:block w-20" alt="">
            </a>
          </div>

        </div>

      </div>
      <div class="flex-1 relative">

        <div class="max-w-[1220px] mx-auto flex items-center">
          <UInput size="xl" color="white" placeholder="Search Friends, videos..." class="dark:!bg-white/5 w-1/3">
            <template #leading>
              <span class="text-sm">🔍</span>
            </template>
          </UInput>

          <!-- header icons -->
          <div class="flex items-center sm:gap-4 gap-2 absolute right-5 top-1/2 -translate-y-1/2 text-black">
            <!-- create -->
            <UPopover>
              <UButton variant="solid" color="gray" :ui="{ rounded: 'rounded-full' }" square label="➕"></UButton>
              <template #panel>
                <div
                  class="bg-white p-4 rounded-lg overflow-hidden drop-shadow-xl dark:bg-gray-900 md:w-[324px] w-screen border2">

                  <h3 class="font-bold text-md text-gray-900 dark:text-white mb-3"> Create </h3>

                  <!-- slider -->
                  <UCarousel v-slot="{ item }" :items="items" :ui="{ item: 'basis-full md:basis-1/3 lg:basis-1/4' }">
                    <div
                      class="p-3 w-full flex items-center mr-4 rounded-lg bg-teal-100/60 text-teal-600 dark:text-white dark:bg-gray-700">
                      <span class="text-3xl -mt-3 drop-shadow-md">{{ item.icon }}</span>
                      <div class="text-sm font-medium">{{ item.label }}</div>
                    </div>
                  </UCarousel>

                  <!-- list -->
                  <ul class="-m-1 mt-4 pb-1 text-xs text-gray-500 dark:text-white">
                    <li
                      class="flex items-center gap-4 hover:bg-secondery rounded-md p-1.5 cursor-pointer dark:hover:bg-white/10">
                      <span class="text-3xl">👥</span>
                      <div class="flex-1">
                        <a href="timeline.html">
                          <h4 class="font-medium text-sm text-black dark:text-white"> Groups </h4>
                        </a>
                        <div class="mt-1 text-xs text-gray-500 dark:text-white"> Meet people with similar interests.
                        </div>
                      </div>
                    </li>
                    <li
                      class="flex items-center gap-4 hover:bg-secondery rounded-md p-1.5 cursor-pointer dark:hover:bg-white/10">
                      <span class="text-3xl">📆</span>
                      <div class="flex-1">
                        <a href="timeline.html">
                          <h4 class="font-medium text-sm text-black dark:text-white"> Event </h4>
                        </a>
                        <div class="mt-1">Discover fun activities near you .</div>
                      </div>
                    </li>
                    <li
                      class="flex items-center gap-4 hover:bg-secondery rounded-md p-1.5 cursor-pointer dark:hover:bg-white/10">
                      <span class="text-3xl">🎮</span>
                      <div class="flex-1">
                        <a href="timeline.html">
                          <h4 class="font-medium text-sm text-black dark:text-white"> Games </h4>
                        </a>
                        <div class="mt-1"> play game with friends have fun. </div>
                      </div>
                    </li>
                  </ul>
                </div>
              </template>
            </UPopover>

            <!-- notification -->
            <UPopover>
              <UButton variant="solid" color="gray" :ui="{ rounded: 'rounded-full' }" square label="🔔"></UButton>
              <template #panel>
                <div class="bg-white pr-1.5 rounded-lg drop-shadow-xl dark:bg-gray-900 md:w-[365px] w-screen border2">

                  <!-- heading -->
                  <div class="flex items-center justify-between gap-2 p-4 pb-2">
                    <h3 class="font-bold text-xl text-gray-900 dark:text-white"> Notifications </h3>
                  </div>

                  <div class="text-sm h-[400px] w-full overflow-y-auto pr-2">

                    <!-- contents list -->
                    <div class="pl-2 p-1 text-sm font-normal dark:text-white">

                      <a href="#"
                        class="relative flex items-center gap-3 p-2 duration-200 rounded-xl pr-10 hover:bg-secondery dark:hover:bg-white/10 bg-teal-500/5">
                        <div class="relative w-12 h-12 shrink-0"> <img src="assets/images/avatars/avatar-3.jpg" alt=""
                            class="object-cover w-full h-full rounded-full"></div>
                        <div class="flex-1 ">
                          <p> <b class="font-bold mr-1"> Alexa Gray</b> started following you. Welcome him to your
                            profile. 👋
                          </p>
                          <div class="text-xs text-gray-500 mt-1.5 dark:text-white/80"> 4 hours ago </div>
                          <div class="w-2.5 h-2.5 bg-teal-600 rounded-full absolute right-3 top-5"></div>
                        </div>
                      </a>
                      <a href="#"
                        class="relative flex items-center gap-3 p-2 duration-200 rounded-xl pr-10 hover:bg-secondery dark:hover:bg-white/10">
                        <div class="relative w-12 h-12 shrink-0"> <img src="assets/images/avatars/avatar-7.jpg" alt=""
                            class="object-cover w-full h-full rounded-full"></div>
                        <div class="flex-1 ">
                          <p> <b class="font-bold mr-1">Jesse Steeve</b> mentioned you in a story. Check it out and reply.
                            📣
                          </p>
                          <div class="text-xs text-gray-500 mt-1.5 dark:text-white/80"> 8 hours ago </div>
                        </div>
                      </a>
                      <a href="#"
                        class="relative flex items-center gap-3 p-2 duration-200 rounded-xl pr-10 hover:bg-secondery dark:hover:bg-white/10">
                        <div class="relative w-12 h-12 shrink-0"> <img src="assets/images/avatars/avatar-6.jpg" alt=""
                            class="object-cover w-full h-full rounded-full"></div>
                        <div class="flex-1 ">
                          <p> <b class="font-bold mr-1"> Alexa stella</b> commented on your photo “Wow, stunning shot!” 💬
                          </p>
                          <div class="text-xs text-gray-500 mt-1.5 dark:text-white/80"> 8 hours ago </div>
                        </div>
                      </a>
                      <a href="#"
                        class="relative flex items-center gap-3 p-2 duration-200 rounded-xl pr-10 hover:bg-secondery dark:hover:bg-white/10">
                        <div class="relative w-12 h-12 shrink-0"> <img src="assets/images/avatars/avatar-2.jpg" alt=""
                            class="object-cover w-full h-full rounded-full"></div>
                        <div class="flex-1 ">
                          <p> <b class="font-bold mr-1"> John Michael</b> who you might know, is on socialite.</p>
                          <div class="text-xs text-gray-500 mt-1.5 dark:text-white/80"> 2 hours ago </div>
                        </div>
                        <button type="button" class="button text-white bg-primary">fallow</button>
                      </a>
                      <a href="#"
                        class="relative flex items-center gap-3 p-2 duration-200 rounded-xl pr-10 hover:bg-secondery dark:hover:bg-white/10 bg-teal-500/5">
                        <div class="relative w-12 h-12 shrink-0"> <img src="assets/images/avatars/avatar-3.jpg" alt=""
                            class="object-cover w-full h-full rounded-full"></div>
                        <div class="flex-1 ">
                          <p> <b class="font-bold mr-1"> Sarah Gray</b> sent you a message. He wants to chat with you. 💖
                          </p>
                          <div class="text-xs text-gray-500 mt-1.5 dark:text-white/80"> 4 hours ago </div>
                          <div class="w-2.5 h-2.5 bg-teal-600 rounded-full absolute right-3 top-5"></div>
                        </div>
                      </a>
                      <a href="#"
                        class="relative flex items-center gap-3 p-2 duration-200 rounded-xl pr-10 hover:bg-secondery dark:hover:bg-white/10">
                        <div class="relative w-12 h-12 shrink-0"> <img src="assets/images/avatars/avatar-4.jpg" alt=""
                            class="object-cover w-full h-full rounded-full"></div>
                        <div class="flex-1 ">
                          <p> <b class="font-bold mr-1"> Jesse Steeve</b> sarah tagged you <br> in a photo of your
                            birthday
                            party. 📸 </p>
                          <div class="text-xs text-gray-500 mt-1.5 dark:text-white/80"> 8 hours ago </div>
                        </div>
                      </a>
                      <a href="#"
                        class="relative flex items-center gap-3 p-2 duration-200 rounded-xl pr-10 hover:bg-secondery dark:hover:bg-white/10">
                        <div class="relative w-12 h-12 shrink-0"> <img src="assets/images/avatars/avatar-2.jpg" alt=""
                            class="object-cover w-full h-full rounded-full"></div>
                        <div class="flex-1 ">
                          <p> <b class="font-bold mr-1"> Lewis Lewis</b> mentioned you in a story. Check it out and reply.
                            📣
                          </p>
                          <div class="text-xs text-gray-500 mt-1.5 dark:text-white/80"> 8 hours ago </div>
                        </div>
                      </a>
                      <a href="#"
                        class="relative flex items-center gap-3 p-2 duration-200 rounded-xl pr-10 hover:bg-secondery dark:hover:bg-white/10">
                        <div class="relative w-12 h-12 shrink-0"> <img src="assets/images/avatars/avatar-7.jpg" alt=""
                            class="object-cover w-full h-full rounded-full"></div>
                        <div class="flex-1 ">
                          <p> <b class="font-bold mr-1"> Martin Gray</b> liked your photo of the Eiffel Tower. 😍 </p>
                          <div class="text-xs text-gray-500 mt-1.5 dark:text-white/80"> 8 hours ago </div>
                        </div>
                      </a>

                    </div>

                  </div>


                  <!-- footer -->
                  <a href="#">
                    <div
                      class="text-center py-4 border-t border-gray-100 text-sm font-medium text-blue-600 dark:text-white dark:border-gray-600">
                      View Notifications </div>
                  </a>

                  <div
                    class="w-3 h-3 absolute -top-1.5 right-3 bg-white border-l border-t rotate-45 max-md:hidden dark:bg-dark3 dark:border-transparent">
                  </div>
                </div>
              </template>
            </UPopover>

            <!-- messages -->
            <UPopover>
              <UButton variant="solid" color="gray" :ui="{ rounded: 'rounded-full' }" square label="💬"></UButton>
              <template #panel>
                <div class="bg-white pr-1.5 rounded-lg drop-shadow-xl dark:bg-gray-900 md:w-[360px] w-screen border2">

                  <!-- heading -->
                  <div class="flex items-center justify-between gap-2 p-4 pb-1">
                    <h3 class="font-bold text-xl text-gray-900 dark:text-white"> Chats </h3>
                  </div>

                  <div class="h-80 overflow-y-auto pr-2">

                    <div class="p-2 pt-0 pr-1 dark:text-white/80">

                      <a href="#"
                        class="relative flex items-center gap-4 p-2 py-3 duration-200 rounded-lg hover:bg-secondery dark:hover:bg-white/10">
                        <div class="relative w-10 h-10 shrink-0">
                          <img src="assets/images/avatars/avatar-2.jpg" alt=""
                            class="object-cover w-full h-full rounded-full">
                        </div>
                        <div class="flex-1 min-w-0">
                          <div class="flex items-center gap-2 mb-1">
                            <div class="mr-auto text-sm text-black dark:text-white font-medium">Jesse Steeve</div>
                            <div class="text-xs text-gray-500 dark:text-white/80"> 09:40AM </div>
                            <div class="w-2.5 h-2.5 bg-blue-600 rounded-full dark:bg-gray-900"></div>
                          </div>
                          <div class="font-normal overflow-hidden text-ellipsis text-xs whitespace-nowrap">Love your
                            photos 😍
                          </div>
                        </div>
                      </a>
                      <a href="#"
                        class="relative flex items-center gap-4 p-2 py-3 duration-200 rounded-lg hover:bg-secondery dark:hover:bg-white/10">
                        <div class="relative w-10 h-10 shrink-0">
                          <img src="assets/images/avatars/avatar-4.jpg" alt=""
                            class="object-cover w-full h-full rounded-full">
                        </div>
                        <div class="flex-1 min-w-0">
                          <div class="flex items-center gap-2 mb-1">
                            <div class="mr-auto text-sm text-black dark:text-white font-medium">Martin Gray</div>
                            <div class="text-xs text-gray-500 dark:text-white/80"> 02:40AM </div>
                          </div>
                          <div class="font-normal overflow-hidden text-ellipsis text-xs whitespace-nowrap">Product
                            photographer wanted? 📷</div>
                        </div>
                      </a>
                      <a href="#"
                        class="relative flex items-center gap-4 p-2 py-3 duration-200 rounded-lg hover:bg-secondery dark:hover:bg-white/10">
                        <div class="relative w-10 h-10 shrink-0">
                          <img src="assets/images/avatars/avatar-5.jpg" alt=""
                            class="object-cover w-full h-full rounded-full">
                        </div>
                        <div class="flex-1 min-w-0">
                          <div class="flex items-center gap-2 mb-1">
                            <div class="mr-auto text-sm text-black dark:text-white font-medium">Jesse Steeve</div>
                            <div class="text-xs text-gray-500 dark:text-white/80"> 2 day </div>
                          </div>
                          <div class="font-normal overflow-hidden text-ellipsis text-xs whitespace-nowrap">Want to buy
                            landscape photo? 🌄</div>
                        </div>
                      </a>
                      <a href="#"
                        class="relative flex items-center gap-4 p-2 py-3 duration-200 rounded-lg hover:bg-secondery dark:hover:bg-white/10">
                        <div class="relative w-10 h-10 shrink-0">
                          <img src="assets/images/avatars/avatar-3.jpg" alt=""
                            class="object-cover w-full h-full rounded-full">
                        </div>
                        <div class="flex-1 min-w-0">
                          <div class="flex items-center gap-2 mb-1">
                            <div class="mr-auto text-sm text-black dark:text-white font-medium">Monroe Parker</div>
                            <div class="text-xs text-gray-500 dark:text-white/80"> 4 week </div>
                            <div class="w-2.5 h-2.5 bg-blue-600 rounded-full dark:bg-gray-900"></div>
                          </div>
                          <div class="font-normal overflow-hidden text-ellipsis text-xs whitespace-nowrap">I’m glad you
                            like
                            it.😊</div>
                        </div>
                      </a>
                      <a href="#"
                        class="relative flex items-center gap-4 p-2 py-3 duration-200 rounded-lg hover:bg-secondery dark:hover:bg-white/10">
                        <div class="relative w-10 h-10 shrink-0">
                          <img src="assets/images/avatars/avatar-7.jpg" alt=""
                            class="object-cover w-full h-full rounded-full">
                        </div>
                        <div class="flex-1 min-w-0">
                          <div class="flex items-center gap-2 mb-1">
                            <div class="mr-auto text-sm text-black dark:text-white font-medium">Alex Dolve</div>
                            <div class="text-xs text-gray-500 dark:text-white/80"> 2 month </div>
                          </div>
                          <div class="font-normal overflow-hidden text-ellipsis text-xs whitespace-nowrap">Photo editor
                            needed. Fix photos? 🛠️</div>
                        </div>
                      </a>
                      <a href="#"
                        class="relative flex items-center gap-4 p-2 py-3 duration-200 rounded-lg hover:bg-secondery dark:hover:bg-white/10">
                        <div class="relative w-10 h-10 shrink-0">
                          <img src="assets/images/avatars/avatar-4.jpg" alt=""
                            class="object-cover w-full h-full rounded-full">
                        </div>
                        <div class="flex-1 min-w-0">
                          <div class="flex items-center gap-2 mb-1">
                            <div class="mr-auto text-sm text-black dark:text-white font-medium">Jesse Steeve</div>
                            <div class="text-xs text-gray-500 dark:text-white/80"> 09:40AM </div>
                          </div>
                          <div class="font-normal overflow-hidden text-ellipsis text-xs whitespace-nowrap">Love your
                            photos 😍
                          </div>
                        </div>
                      </a>

                    </div>

                  </div>


                  <!-- footer -->
                  <a href="#">
                    <div
                      class="text-center py-4 border-t border-gray-100 text-sm font-medium text-blue-600 dark:text-white dark:border-gray-600">
                      See all Messages </div>
                  </a>

                  <div
                    class="w-3 h-3 absolute -top-1.5 right-3 bg-white border-l border-t rotate-45 max-md:hidden dark:bg-dark3 dark:border-transparent">
                  </div>
                </div>
              </template>
            </UPopover>

            <!-- profile -->
            <UPopover>
              <UAvatar size="sm" src="/assets/images/avatars/avatar-2.jpg" alt="Avatar" />

              <template #panel>
                <div class="bg-white rounded-lg drop-shadow-xl dark:bg-gray-900 w-64 border2">

                  <a href="timeline.html">
                    <div class="p-4 py-5 flex items-center gap-4">
                      <img src="assets/images/avatars/avatar-2.jpg" alt="" class="w-10 h-10 rounded-full shadow">
                      <div class="flex-1">
                        <h4 class="text-sm font-medium text-black dark:text-white">Stell johnson</h4>
                        <div class="text-sm mt-1 text-blue-600 font-light dark:text-white/70">@mohnson</div>
                      </div>
                    </div>
                  </a>

                  <hr class="dark:border-gray-600/60">

                  <nav class="p-2 text-sm text-black font-normal dark:text-white">
                    <a href="upgrade.html">
                      <div
                        class="flex items-center gap-2.5 hover:bg-secondery p-2 px-2.5 rounded-md dark:hover:bg-white/10 text-blue-600">
                        ⚡
                        Upgrade To Premium
                      </div>
                    </a>
                    <a href="setting.html">
                      <div
                        class="flex items-center gap-2.5 hover:bg-secondery p-2 px-2.5 rounded-md dark:hover:bg-white/10">
                        💳
                        My Billing
                      </div>
                    </a>
                    <a href="setting.html">
                      <div
                        class="flex items-center gap-2.5 hover:bg-secondery p-2 px-2.5 rounded-md dark:hover:bg-white/10">
                        📢
                        Advertising
                      </div>
                    </a>
                    <a href="setting.html">
                      <div
                        class="flex items-center gap-2.5 hover:bg-secondery p-2 px-2.5 rounded-md dark:hover:bg-white/10">
                        ⚙️
                        My Account
                      </div>
                    </a>
                    <ColorSwitch />
                    <hr class="-mx-2 my-2 dark:border-gray-600/60">
                    <a href="form-login.html">
                      <div
                        class="flex items-center gap-2.5 hover:bg-secondery p-2 px-2.5 rounded-md dark:hover:bg-white/10">
                        👋
                        Log Out
                      </div>
                    </a>

                  </nav>

                </div>
              </template>
            </UPopover>
          </div>
        </div>
      </div>
    </div>
  </header>
</template>
<script setup="ts">
const items = [
  { icon: '📖', label: 'Story' },
  { icon: '📷', label: 'Post' },
  { icon: '🎥', label: 'Reel' },
  { icon: '📌', label: 'location' },
  { icon: '😆', label: 'Status' },
]
</script>