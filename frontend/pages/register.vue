<template>
  <div class="sm:flex">
    <div
      class="relative lg:w-[580px] md:w-96 w-full p-10 min-h-screen bg-white shadow-xl flex items-center pt-10 dark:bg-slate-900 z-10"
    >
      <div
        class="w-full lg:max-w-sm mx-auto space-y-10"
        uk-scrollspy="target: > *; cls: uk-animation-scale-up; delay: 100 ;repeat: true"
      >
        <!-- logo image-->
        <a href="#"> <img
          src="../template/assets/images/logo.png" class="w-28 absolute top-10 left-10 dark:hidden"
          alt=""
        ></a>
        <a href="#"> <img
          src="../template/assets/images/logo-light.png"
          class="w-28 absolute top-10 left-10 hidden dark:!block" alt=""
        ></a>

        <!-- logo icon optional -->
        <div class="hidden">
          <img class="w-12" src="../template/assets/images/logo-icon.png" alt="Socialite html template">
        </div>

        <!-- title -->
        <div>
          <h2 class="text-2xl font-semibold mb-1.5">
            Sign up to get started
          </h2>
          <p class="text-sm text-gray-700 font-normal">
            If you already have an account, <a
              href="form-login.html"
              class="text-blue-700"
            >Login here!</a>
          </p>
        </div>


        <!-- form -->
        <UForm
          :validate="validate" :state="state" method="#" action="#"
          class="space-y-7 text-sm text-black font-medium dark:text-white"
          uk-scrollspy="target: > *; cls: uk-animation-scale-up; delay: 100 ;repeat: true" @submit="onSubmit"
        >
          <div class="grid grid-cols-2 gap-4 gap-y-7">
            <!-- Last name -->
            <UFormGroup label="First name" name="firstname">
              <div class="mt-2.5">
                <UInput
                  id="text" v-model="state.firstName" name="text" type="text" placeholder="First name"
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5"
                />
              </div>
            </UFormGroup>
            <!-- first name -->
            <UFormGroup label="Last name" name="lastname">
              <div class="mt-2.5">
                <UInput
                  id="text" v-model="state.lastName" name="text" type="text" placeholder="Last name"
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5"
                />
              </div>
            </UFormGroup>
            <!-- email -->
            <UFormGroup label="Email adresse" name="email" class="col-span-2">
              <div class="mt-2.5">
                <UInput
                  id="text" v-model="state.email" name="pass" type="text" placeholder="Email"
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5"
                />
              </div>
            </UFormGroup>
            <UFormGroup label="Date of birth" name="dateOfBirth" class="col-span-2">
              <div class="mt-2.5">
                <UInput
                  id="dateOfBirth"
                  v-model="state.dateOfBirth"
                  type="date"
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5"
                />
              </div>
            </UFormGroup>
            <!-- password -->
            <UFormGroup label="Password" name="password">
              <div class="mt-2.5">
                <UInput
                  id="text" v-model="state.password" name="password" type="password" placeholder="***"
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5"
                />
              </div>
            </UFormGroup>
            <UFormGroup label="Confirm Password" name="password">
              <div class="mt-2.5">
                <UInput
                  id="text" v-model="state.password" name="password" type="password" placeholder="***"
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5"
                />
              </div>
            </UFormGroup>
            <UFormGroup label="Nickname" name="nickname">
              <div class="mt-2.5">
                <UInput
                  id="text" v-model="state.password" name="nickname" type="text" placeholder="Nickname"
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5"
                />
              </div>
            </UFormGroup>

            <UFormGroup label="About you" name="aboutMe" class="col-span-2">
              <div class="mt-2.5">
                <UTextarea
                  v-model="state.aboutMe" type="text" placeholder=""
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5"
                />
              </div>
            </UFormGroup>
            <UFormGroup label="Profile Picture" name="avatar" class="col-span-2">
              <div class="mt-2.5">
                <UInput
                  id="avatar"
                  type="file"
                  accept="image/*"
                  class="!w-full !rounded-lg !bg-transparent !shadow-sm !border-slate-200 dark:!border-slate-800 dark:!bg-white/5"
                  @change="handleAvatarUpload"
                />
              </div>
            </UFormGroup>
            <!-- submit button -->
            <div class="col-span-0">
              <UButton type="submit" class=" button bg-primary text-white w-full">
                Get Started
              </UButton>
            </div>
          </div>

          <div class="text-center flex items-center gap-6">
            <hr class="flex-1 border-slate-200 dark:border-slate-800">
            Or continue with
            <hr class="flex-1 border-slate-200 dark:border-slate-800">
          </div>

          <!-- social login -->
          <div class="flex gap-2" uk-scrollspy="target: > *; cls: uk-animation-scale-up; delay: 400 ;repeat: true">
            <a href="#" class="button flex-1 flex items-center gap-2 bg-primary text-white text-sm"> <ion-icon
              name="logo-facebook" class="text-lg"
            /> facebook </a>
            <a href="#" class="button flex-1 flex items-center gap-2 bg-sky-600 text-white text-sm"> <ion-icon
              name="logo-twitter"
            /> twitter </a>
            <a href="#" class="button flex-1 flex items-center gap-2 bg-black text-white text-sm"> <ion-icon
              name="logo-github"
            /> github </a>
          </div>
        </UForm>
      </div>
    </div>

    <!-- image slider -->
    <div class="flex-1 relative bg-primary max-md:hidden">
      <div class="relative w-full h-full" tabindex="-1" uk-slideshow="animation: slide; autoplay: true">
        <ul class="uk-slideshow-items w-full h-full">
          <li class="w-full">
            <img
              src="../template/assets/images/post/img-3.jpg" alt=""
              class="w-full h-full object-cover uk-animation-kenburns uk-animation-reverse uk-transform-origin-center-left"
            >
            <div class="absolute bottom-0 w-full uk-tr ansition-slide-bottom-small z-10">
              <div
                class="max-w-xl w-full mx-auto pb-32 px-5 z-30 relative"
                uk-scrollspy="target: > *; cls: uk-animation-scale-up; delay: 100 ;repeat: true"
              >
                <img class="w-12" src="../template/assets/images/logo-icon.png" alt="Socialite html template">
                <h4 class="!text-white text-2xl font-semibold mt-7" uk-slideshow-parallax="y: 600,0,0">
                  Connect With
                  Friends
                </h4>
                <p class="!text-white text-lg mt-7 leading-8" uk-slideshow-parallax="y: 800,0,0;">
                  This phrase is more
                  casual and playful. It suggests that you are keeping your friends updated on what’s happening in your
                  life.
                </p>
              </div>
            </div>
            <div class="w-full h-96 bg-gradient-to-t from-black absolute bottom-0 left-0" />
          </li>
          <li class="w-full">
            <img
              src="../template/assets/images/post/img-2.jpg" alt=""
              class="w-full h-full object-cover uk-animation-kenburns uk-animation-reverse uk-transform-origin-center-left"
            >
            <div class="absolute bottom-0 w-full uk-tr ansition-slide-bottom-small z-10">
              <div
                class="max-w-xl w-full mx-auto pb-32 px-5 z-30 relative"
                uk-scrollspy="target: > *; cls: uk-animation-scale-up; delay: 100 ;repeat: true"
              >
                <img class="w-12" src="../template/assets/images/logo-icon.png" alt="Socialite html template">
                <h4 class="!text-white text-2xl font-semibold mt-7" uk-slideshow-parallax="y: 800,0,0">
                  Connect With
                  Friends
                </h4>
                <p class="!text-white text-lg mt-7 leading-8" uk-slideshow-parallax="y: 800,0,0;">
                  This phrase is more
                  casual and playful. It suggests that you are keeping your friends updated on what’s happening in your
                  life.
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
  password: undefined,
  firstName: undefined,
  lastName:undefined,
  dateOfBirth:undefined,
  avatar: null as File | null,
  nickname: undefined,
  aboutMe: undefined,
  avatarPreview: '' as any,
  avatarError: '' ,
})

const validate = (state: any): FormError[] => {
  const errors = []
  if (!state.email) errors.push({ path: 'email', message: 'Email is required' })
  if (!state.password) errors.push({ path: 'password', message: 'Password is required' })
  if (!state.firstName) errors.push({ path: 'firstname', message: 'First Name is required' })
  if (!state.lastName) errors.push({ path: 'lastname', message: 'Last Name is required' })
  if (!state.dateOfBirth) errors.push({ path: 'dateOfBirth', message: 'Date of Birth is required' })
  return errors
}

async function onSubmit (event: FormSubmitEvent<any>) {
  // Do something with data
  console.log(event.data)
}
// const value = ref('')
function handleAvatarUpload(event: Event) {
  const file = (event.target as HTMLInputElement).files?.[0]

  // Check if a file was selected
  if (file) {
    // Check the file extension (e.g., allow only images)
    const allowedExtensions = ['jpg', 'jpeg', 'png','img']
    const fileExtension = file.name.split('.').pop()?.toLowerCase() || ''

    if (allowedExtensions.includes(fileExtension)) {
      // Preview the image
      const reader = new FileReader()
      reader.onload = () => {
        state.avatarPreview = reader.result
      }
      reader.readAsDataURL(file)

      // Store the file in state.avatar
      state.avatar = file

      // Reset the error
      state.avatarError = ''
    } else {
      // Display an error if the extension is not allowed
      state.avatarError = 'Please select a valid image (jpg, jpeg, png, gif).'
      state.avatar = null
      state.avatarPreview = ''
    }
  }
}
</script>