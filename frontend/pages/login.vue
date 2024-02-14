<template>
  <div class="sm:flex">
      <div
        class="relative lg:w-[580px] md:w-96 w-full p-10 min-h-screen bg-white shadow-xl flex items-center pt-10 dark:bg-slate-900 z-10">
        <div class="w-full lg:max-w-sm mx-auto space-y-10"
          uk-scrollspy="target: > *; cls: uk-animation-scale-up; delay: 100 ;repeat: true">
          <!-- logo image-->
          <a href="#"> <img src="/_nuxt/assets/images/logo.png" class="w-28 absolute top-10 left-10 dark:hidden"
              alt=""></a>
          <a href="#"> <img src="/_nuxt/assets/images/logo-light.png"
              class="w-28 absolute top-10 left-10 hidden dark:!block" alt=""></a>

          <!-- logo icon optional -->
          <div class="hidden">
            <img class="w-12" src="https://tailwindui.com/img/logos/mark.svg?color=indigo&amp;shade=600"
              alt="Socialite html template">
          </div>

          <!-- title -->
          <div>
            <h2 class="text-2xl font-semibold mb-1.5">
              Sign in to your account
            </h2>
            <p class="text-sm text-gray-700 font-normal">
              If you haven’t signed up yet. <a href="register" class="text-blue-700">Register here!</a>
            </p>
          </div>


          <!-- form -->
          <UForm :validate="validate" :state="state" method="#" action="#"
            class="space-y-7 text-sm text-black font-medium dark:text-white"
            uk-scrollspy="target: > *; cls: uk-animation-scale-up; delay: 100 ;repeat: true" @submit="onSubmit">
            <!-- email -->
            <UFormGroup label="Email" name="email">
              <div class="mt-2.5">
                <UInput id="email" v-model="state.email" name="email" type="email" placeholder="Email"
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5"
                  icon="i-heroicons-envelope" />
              </div>
            </UFormGroup>
            <!-- password -->
            <UFormGroup label="Password" name="password">
              <div class="mt-2.5">
                <UInput id="password" v-model="state.password" name="password" type="password" placeholder="***"
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5" />
              </div>
            </UFormGroup>
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-2.5">
                <input id="rememberme" name="rememberme" type="checkbox">
                <label for="rememberme" class="font-normal">Remember me</label>
              </div>
              <a href="#" class="text-blue-700">Forgot password </a>
            </div>

            <!-- submit button -->
            <div>
              <UButton type="submit" class="button bg-primary text-white w-full">
                Sign in
              </UButton>
            </div>

            <div class="text-center flex items-center gap-6">
              <hr class="flex-1 border-slate-200 dark:border-slate-800">
              Or continue with
              <hr class="flex-1 border-slate-200 dark:border-slate-800">
            </div>

            <!-- social login -->
            <div class="flex gap-2" uk-scrollspy="target: > *; cls: uk-animation-scale-up; delay: 400 ;repeat: true">
              <a href="#" class="button flex-1 flex items-center gap-2 bg-primary text-white text-sm"> <ion-icon
                  name="logo-facebook" class="text-lg" /> facebook </a>
              <a href="#" class="button flex-1 flex items-center gap-2 bg-sky-600 text-white text-sm"> <ion-icon
                  name="logo-twitter" /> twitter </a>
              <a href="#" class="button flex-1 flex items-center gap-2 bg-black text-white text-sm"> <ion-icon
                  name="logo-github" /> github </a>
            </div>
          </UForm>
        </div>
      </div>

      <!-- image slider -->
      <div class="flex-1 relative bg-primary max-md:hidden">
        <div class="relative w-full h-full" tabindex="-1" uk-slideshow="animation: slide; autoplay: true">
          <ul class="uk-slideshow-items w-full h-full">
            <li class="w-full">
              <img src="/_nuxt/assets/images/post/img-3.jpg" alt=""
                class="w-full h-full object-cover uk-animation-kenburns uk-animation-reverse uk-transform-origin-center-left">
              <div class="absolute bottom-0 w-full uk-tr ansition-slide-bottom-small z-10">
                <div class="max-w-xl w-full mx-auto pb-32 px-5 z-30 relative"
                  uk-scrollspy="target: > *; cls: uk-animation-scale-up; delay: 100 ;repeat: true">
                  <img class="w-12" src="/_nuxt/assets/images/logo-icon.png" alt="Socialite html template">
                  <h4 class="!text-white text-2xl font-semibold mt-7" uk-slideshow-parallax="y: 600,0,0">
                    Connect With Friends
                  </h4>
                  <p class="!text-white text-lg mt-7 leading-8" uk-slideshow-parallax="y: 800,0,0;">
                    This phrase is more casual and playful. It suggests that you are keeping your friends updated on
                    what’s happening in your life.
                  </p>
                </div>
              </div>
              <div class="w-full h-96 bg-gradient-to-t from-black absolute bottom-0 left-0" />
            </li>
            <li class="w-full">
              <img src="/_nuxt/assets/images/post/img-2.jpg" alt=""
                class="w-full h-full object-cover uk-animation-kenburns uk-animation-reverse uk-transform-origin-center-left">
              <div class="absolute bottom-0 w-full uk-tr ansition-slide-bottom-small z-10">
                <div class="max-w-xl w-full mx-auto pb-32 px-5 z-30 relative"
                  uk-scrollspy="target: > *; cls: uk-animation-scale-up; delay: 100 ;repeat: true">
                  <img class="w-12" src="/_nuxt/assets/images/logo-icon.png" alt="Socialite html template">
                  <h4 class="!text-white text-2xl font-semibold mt-7" uk-slideshow-parallax="y: 800,0,0">
                    Connect With Friends
                  </h4>
                  <p class="!text-white text-lg mt-7 leading-8" uk-slideshow-parallax="y: 800,0,0;">
                    This phrase is more casual and playful. It suggests that you are keeping your friends updated on
                    what’s happening in your life.
                  </p>
                </div>
              </div>
              <div class="w-full h-96 bg-gradient-to-t from-black absolute bottom-0 left-0" />
            </li>
          </ul>

          <!-- slide nav -->
          <div class="flex justify-center">
            <ul class="inline-flex flex-wrap justify-center  absolute bottom-8 gap-1.5 uk-dotnav uk-slideshow-nav" />
          </div>
        </div>
      </div>
    </div>
</template>
<script setup lang="ts">



import type { FormError, FormSubmitEvent } from '#ui/types'

const state = reactive({
  email: undefined,
  password: undefined
})

const validate = (state: any): FormError[] => {
  const errors = []
  if (!state.email) errors.push({ path: 'email', message: 'Required' })
  if (!state.password) errors.push({ path: 'password', message: 'Required' })
  return errors
}

async function onSubmit(event: FormSubmitEvent<any>) {
  // Do something with data
  console.log(event.data)
}
</script>