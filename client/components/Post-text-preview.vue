<template>
  <div class=" lg:p-20 max-lg:!items-start " id="text_preview_modal" uk-modal="">

      <div
        class="lg:w-[400px] w-full bg-white h-full relative  overflow-y-auto shadow-xl dark:bg-dark2 flex flex-col justify-between">

        <div class="p-5 pb-0">

          <!-- story heading -->
          <div class="flex gap-3 text-sm font-medium">
            <img src="assets/images/avatars/avatar-5.jpg" alt="" class="w-9 h-9 rounded-full">
            <div class="flex-1">
              <h4 class="text-black font-medium dark:text-white"> Steeve </h4>
              <div class="text-gray-500 text-xs dark:text-white/80"> 2 hours ago</div>
            </div>

            <!-- dropdown -->
            <div class="-m-1">
              <button type="button" class="button__ico w-8 h-8"> <i class='bx bx-dots-horizontal-rounded text-xl' ></i> </button>
              <div class="w-[253px]"
                uk-dropdown="pos: bottom-right; animation: uk-animation-scale-up uk-transform-origin-top-right; animate-out: true">
                <nav>
                  <a href="#"> <i class='bx bx-bookmark text-xl shrink-0' ></i> 
                    Add to favorites </a>
                  <a href="#"> <i class='bx bx-bell-off text-xl shrink-0' ></i> 
                    Mute Notification </a>
                  <a href="#"> <i class='bx bx-flag text-xl shrink-0' ></i> 
                    Report this post </a>
                  <a href="#"> <i class='bx bx-share text-xl shrink-0' ></i>
                    Share your profile </a>
                  <hr>
                  <a href="#" class="text-red-400 hover:!bg-red-50 dark:hover:!bg-red-500/50">
                    <i class='bx bx-stop-circle text-xl shrink-0' ></i>
                    Unfollow </a>
                </nav>
              </div>
            </div>
          </div>

          <p class="font-normal text-sm leading-6 mt-4"> Photography is the art of capturing light with a
            camera. it can be fun, challenging. It can also be a hobby, a passion. 📷 </p>

          <div class="shadow relative -mx-5 px-5 py-3 mt-3">
            <div class="flex items-center gap-4 text-xs font-semibold">
              <div class="flex items-center gap-2.5">
                <button type="button" class="button__ico text-red-500 bg-red-100 dark:bg-slate-700">
                  <i class='bx bx-heart text-lg' ></i>
                  </button>
                <a href="#">1,300</a>
              </div>
              <div class="flex items-center gap-3">
                <button type="button" class="button__ico bg-slate-100 dark:bg-slate-700"> <i class='bx bx-message-rounded-dots text-lg'>
                </i> </button>
                <span>260</span>
              </div>
              <button type="button" class="button__ico ml-auto"> <i class='bx bx-share text-xl' ></i> 
              </button>
              <button type="button" class="button__ico"> <i class='bx bx-bookmark text-xl' ></i>
              </button>
            </div>
          </div>

        </div>

        <div class="p-5 h-full overflow-y-auto flex-1">

          <!-- comment list -->
          <div class="relative text-sm font-medium space-y-5">

            <div class="flex items-start gap-3 relative">
              <img src="assets/images/avatars/avatar-2.jpg" alt="" class="w-6 h-6 mt-1 rounded-full">
              <div class="flex-1">
                <a href="#" class="text-black font-medium inline-block dark:text-white"> Steeve </a>
                <p class="mt-0.5">What a beautiful, I love it. 😍 </p>
              </div>
            </div>
            <div class="flex items-start gap-3 relative">
              <img src="assets/images/avatars/avatar-3.jpg" alt="" class="w-6 h-6 mt-1 rounded-full">
              <div class="flex-1">
                <a href="#" class="text-black font-medium inline-block dark:text-white"> Monroe </a>
                <p class="mt-0.5"> You captured the moment.😎 </p>
              </div>
            </div>
            <div class="flex items-start gap-3 relative">
              <img src="assets/images/avatars/avatar-7.jpg" alt="" class="w-6 h-6 mt-1 rounded-full">
              <div class="flex-1">
                <a href="#" class="text-black font-medium inline-block dark:text-white"> Alexa </a>
                <p class="mt-0.5"> This photo is amazing! </p>
              </div>
            </div>
            <div class="flex items-start gap-3 relative">
              <img src="assets/images/avatars/avatar-4.jpg" alt="" class="w-6 h-6 mt-1 rounded-full">
              <div class="flex-1">
                <a href="#" class="text-black font-medium inline-block dark:text-white"> John </a>
                <p class="mt-0.5"> Wow, You are so talented 😍 </p>
              </div>
            </div>
            <div class="flex items-start gap-3 relative">
              <img src="assets/images/avatars/avatar-5.jpg" alt="" class="w-6 h-6 mt-1 rounded-full">
              <div class="flex-1">
                <a href="#" class="text-black font-medium inline-block dark:text-white"> Michael </a>
                <p class="mt-0.5"> I love taking photos 🌳🐶</p>
              </div>
            </div>
            <div class="flex items-start gap-3 relative">
              <img src="assets/images/avatars/avatar-3.jpg" alt="" class="w-6 h-6 mt-1 rounded-full">
              <div class="flex-1">
                <a href="#" class="text-black font-medium inline-block dark:text-white"> Monroe </a>
                <p class="mt-0.5"> Awesome. 😊😢 </p>
              </div>
            </div>
            <div class="flex items-start gap-3 relative">
              <img src="assets/images/avatars/avatar-5.jpg" alt="" class="w-6 h-6 mt-1 rounded-full">
              <div class="flex-1">
                <a href="#" class="text-black font-medium inline-block dark:text-white"> Jesse </a>
                <p class="mt-0.5"> Well done 🎨📸 </p>
              </div>
            </div>
            <div class="flex items-start gap-3 relative">
              <img src="assets/images/avatars/avatar-2.jpg" alt="" class="w-6 h-6 mt-1 rounded-full">
              <div class="flex-1">
                <a href="#" class="text-black font-medium inline-block dark:text-white"> Steeve </a>
                <p class="mt-0.5">What a beautiful, I love it. 😍 </p>
              </div>
            </div>
            <div class="flex items-start gap-3 relative">
              <img src="assets/images/avatars/avatar-7.jpg" alt="" class="w-6 h-6 mt-1 rounded-full">
              <div class="flex-1">
                <a href="#" class="text-black font-medium inline-block dark:text-white"> Alexa </a>
                <p class="mt-0.5"> This photo is amazing! </p>
              </div>
            </div>
            <div class="flex items-start gap-3 relative">
              <img src="assets/images/avatars/avatar-4.jpg" alt="" class="w-6 h-6 mt-1 rounded-full">
              <div class="flex-1">
                <a href="#" class="text-black font-medium inline-block dark:text-white"> John </a>
                <p class="mt-0.5"> Wow, You are so talented 😍 </p>
              </div>
            </div>
            <div class="flex items-start gap-3 relative">
              <img src="assets/images/avatars/avatar-5.jpg" alt="" class="w-6 h-6 mt-1 rounded-full">
              <div class="flex-1">
                <a href="#" class="text-black font-medium inline-block dark:text-white"> Michael </a>
                <p class="mt-0.5"> I love taking photos 🌳🐶</p>
              </div>
            </div>
            <div class="flex items-start gap-3 relative">
              <img src="assets/images/avatars/avatar-3.jpg" alt="" class="w-6 h-6 mt-1 rounded-full">
              <div class="flex-1">
                <a href="#" class="text-black font-medium inline-block dark:text-white"> Monroe </a>
                <p class="mt-0.5"> Awesome. 😊😢 </p>
              </div>
            </div>

          </div>

        </div>

        <div class="bg-white p-3 text-sm font-medium flex items-center gap-2">

          <img src="assets/images/avatars/avatar-2.jpg" alt="" class="w-6 h-6 rounded-full">

          <div class="flex-1 relative overflow-hidden ">
            <textarea placeholder="Add Comment...." rows="1"
              class="w-full resize-  px-4 py-2 focus:!border-transparent focus:!ring-transparent resize-y"></textarea>

            <div class="flex items-center gap-2 absolute bottom-0.5 right-0 m-3">
              <i class='bx bx-bookmark text-xl flex text-blue-700' ></i>
              <i class='bx bx-happy-alt text-xl flex text-yellow-500' ></i>
            </div>

          </div>

          <button type="submit" class="hidden text-sm rounded-full py-1.5 px-4 font-semibold bg-secondery">
            Replay</button>

        </div>

      </div>

  </div>
</template>

<script>
</script> 

<style>
</style>