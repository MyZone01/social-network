<template>
  <div class="overflow-visible">
    <slot />

<<<<<<< HEAD
    <div
      class="space-y-7 text-sm text-black font-medium dark:text-white"
      uk-scrollspy="target: > *; cls: uk-animation-scale-up; delay: 100 ;repeat: true"
    >
      <h2 class="space-y-7 text-xl text-red-500 font-bold dark:text-red-500">
        {{ data.loginError }}
      </h2>

      <UIInput v-model="data.email" label="Email" placeholder="mail@social.net" />
=======
        <div class="space-y-7 text-sm text-black font-medium dark:text-white"
            uk-scrollspy="target: > *; cls: uk-animation-scale-up; delay: 100 ;repeat: true">
            <h2 class="space-y-7 text-xl text-red-500 font-bold dark:text-red-500">{{ data.loginError }}</h2>

            <UIInput v-model="data.email" label="Email" placeholder="mail@social.net" required />

            <UIInput v-model="data.password" label="Password" placeholder="********" type="password" required />
>>>>>>> origin/master

      <UIInput v-model="data.password" label="Password" placeholder="********" type="password" />

      <div class="flex items-center justify-between">
        <div class="flex items-center gap-2.5">
          <input id="rememberme" name="rememberme" type="checkbox" class="!rounded-md accent-red-800">
          <label for="rememberme" class="font-normal">Remember me</label>
        </div>
        <a href="#" class="text-blue-700">Forgot password </a>
      </div>

      <div>
        <button class="button bg-primary text-white w-full" @click="handleLogin()">
          Sign in
        </button>
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
    </div>
  </div>
</template>
<script setup>
const { login } = useAuth()

const data = reactive({
    email: '',
    password: '',
    loginError: '',
    loading: false
})
// const loginError = ''
async function handleLogin() {
    data.loading = true
    try {
        if (email && password) {
            await login({
                email: data.email.trim(),
                password: data.password.trim()
            })
        } else {

        }
    } catch (error) {
        data.loginError = error.statusMessage
        setTimeout(() => {
            data.loginError = ''
        }, 2000)
        data.loading = false
    } finally {
        data.email = ''
        data.password = ''
        data.loading = false
    }
}

</script>