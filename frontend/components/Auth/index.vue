<template>
  <div class="dark:bg-gray-900 bg-white overflow-visible">
    <div class="sm:flex overflow-visible">
      <div
        class="relative lg:w-[580px] md:w-96 w-full p-10 min-h-screen bg-white shadow-xl flex items-center pt-10 dark:bg-slate-900 z-10"
      >
        <div
          class="w-full lg:max-w-sm mx-auto space-y-10"
          uk-scrollspy="target: > *; cls: uk-animation-scale-up; delay: 100 ;repeat: true  overflow-visible"
        >
          <!-- logo image-->
          <a href="#"> <img src="assets/images/logo.png" class="w-28 absolute top-10 left-10 dark:hidden" alt=""></a>
          <a href="#"> <img
            src="assets/images/logo-light.png" class="w-28 absolute top-10 left-10 hidden dark:!block"
            alt=""
          ></a>

          <!-- logo icon optional -->
          <div class="hidden">
            <img
              class="w-12" src="https://tailwindui.com/img/logos/mark.svg?color=indigo&amp;shade=600"
              alt="Socialite html template"
            >
          </div>

          <AuthLoginForm v-if="data.loginProcess">
            <!-- title -->
            <div>
              <h2 class="text-2xl font-semibold mb-1.5">
                Sign in to your account
              </h2>
              <p class="text-sm text-gray-700 font-normal">
                If you haven’t signed up yet.
                <a class="text-blue-700" @click="handleProcess()">Register here!</a>
              </p>
            </div>
          </AuthLoginForm>
          <AuthRegisterForm v-else>
            <!-- title -->
            <div>
              <h2 class="text-2xl font-semibold mb-1.5">
                Sign up to get started
              </h2>
              <p class="text-sm text-gray-700 font-normal">
                If you already have an account, <a class="text-blue-700" @click="handleProcess()">Login
                  here!</a>
              </p>
            </div>
          </AuthRegisterForm>
        </div>
      </div>

      <UIAuthSlider />
    </div>
  </div>
</template>
<script setup>
useHead(() => ({
  title: (title) => 'Social - Authentication',
  meta: [
    { charset: 'utf-8' },
    { name: 'viewport', content: 'width=device-width, initial-scale=1' },
    { hid: 'description', name: 'description', content: 'Authentication for Access' }
  ],
  link: [
    { rel: 'icon', type: 'image/x-icon', href: '~/assets/images/favicon.png' },
    { rel: 'stylesheet', href: 'https://fonts.googleapis.com/css2?family=Inter:wght@200;300;400;500;600;700;800&display=swap' },
    { rel: 'stylesheet', href: '~/assets/css/style.css' },
    { rel: 'stylesheet', href: '~/assets/css/tailwind.css' },
  ],
  script: [
    { src: 'https://cdnjs.cloudflare.com/ajax/libs/jquery/3.5.1/jquery.min.js' },
    { type: 'module', src: 'https://unpkg.com/ionicons@5.5.2/dist/ionicons/ionicons.esm.js' },
    { src: 'https://unpkg.com/ionicons@5.5.2/dist/ionicons/ionicons.js' }

  ]
}))


const data = reactive({
  loginProcess: true
})

function handleProcess() {
  if (this.data.loginProcess) {
    this.data.loginProcess = false
  } else {
    this.data.loginProcess = true
  }
}

onMounted(() => {
  // On page load or when changing themes, best to add inline in `head` to avoid FOUC
  if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
    document.documentElement.classList.add('dark')
  } else {
    document.documentElement.classList.remove('dark')
  }

  // Whenever the user explicitly chooses light mode
  localStorage.theme = 'light'

  // Whenever the user explicitly chooses dark mode
  localStorage.theme = 'dark'

  // Whenever the user explicitly chooses to respect the OS preference
  localStorage.removeItem('theme')

  useHead(() => {
    return {
      script: [
        {
          hid: 'theme-script',
          innerHTML: `
            if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
              document.documentElement.classList.add('dark')
            } else {
              document.documentElement.classList.remove('dark')
            }
          `
        }
      ]
    }
  })
})

// useHead(ref="https://fonts.googleapis.com/css2?family=Inter:wght@200;300;400;500;600;700;800&display=swap")
</script>