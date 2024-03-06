<template>
  <div class="flex-1 relative bg-primary max-md:hidden">
    <div class="relative w-full h-full" tabindex="-1" uk-slideshow="animation: slide; autoplay: true">
      <ul class="uk-slideshow-items w-full h-full">
        <li class="w-full">
          <img
            src="~/assets/images/post/img-3.jpg" alt=""
            class="w-full h-full object-cover uk-animation-kenburns uk-animation-reverse uk-transform-origin-center-left"
          >
          <div class="absolute bottom-0 w-full uk-transition-slide-bottom-small z-10">
            <div
              class="max-w-xl w-full mx-auto pb-32 px-5 z-30 relative"
              uk-scrollspy="target: > *; cls: uk-animation-scale-up; delay: 100; repeat: true"
            >
              <img class="w-12" src="~/assets/images/logo-icon.png" alt="Socialite html template">
              <h4 class="text-white text-2xl font-semibold mt-7" uk-slideshow-parallax="y: 600,0,0">
                Connect
                With Friends
              </h4>
              <p class="text-white text-lg mt-7 leading-8" uk-slideshow-parallax="y: 800,0,0;">
                And Stay Connected by sharing wonderful memories.
              </p>
            </div>
          </div>
          <div class="w-full h-96 bg-gradient-to-t from-black absolute bottom-0 left-0" />
        </li>
        <li class="w-full">
          <img
            src="~/assets/images/post/img-2.jpg" alt=""
            class="w-full h-full object-cover uk-animation-kenburns uk-animation-reverse uk-transform-origin-center-left"
          >
          <div class="absolute bottom-0 w-full uk-transition-slide-bottom-small z-10">
            <div
              class="max-w-xl w-full mx-auto pb-32 px-5 z-30 relative"
              uk-scrollspy="target: > *; cls: uk-animation-scale-up; delay: 100; repeat: true"
            >
              <img class="w-12" src="~/assets/images/logo-icon.png" alt="Socialite html template">
              <h4 class="text-white text-2xl font-semibold mt-7" uk-slideshow-parallax="y: 600,0,0">
                Connect
                With Friends
              </h4>
              <p class="text-white text-lg mt-7 leading-8" uk-slideshow-parallax="y: 800,0,0;">
                And Stay Connected by sharing wonderful memories.
              </p>
            </div>
          </div>
          <div class="w-full h-96 bg-gradient-to-t from-black absolute bottom-0 left-0" />
        </li>
        <!-- Add more slides as needed -->
      </ul>
      <!-- slide nav -->
      <div class="flex justify-center">
        <ul class="inline-flex flex-wrap justify-center absolute bottom-8 gap-1.5 uk-dotnav uk-slideshow-nav" />
      </div>
    </div>
  </div>
</template>

<script>

</script>

<style scoped>/* Add custom styling here */</style>
