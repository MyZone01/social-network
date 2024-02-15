<template>
  <div>
    <slot />

    <form
      class="space-y-7 text-sm text-black font-medium dark:text-white"
      uk-scrollspy="target: > *; cls: uk-animation-scale-up; delay: 100 ;repeat: true"
    >
      <div>
        <!-- email -->
        <h2>{{ data.email }}</h2>
        <UIInput v-model="data.email" label="Email" placeholder="@email" />

        <!-- password -->
        <UIInput v-model="data.password" label="Password" placeholder="********" type="password" />

        <div class="flex items-center justify-between">
          <div class="flex items-center gap-2.5">
            <input id="rememberme" name="rememberme" type="checkbox">
            <label for="rememberme" class="font-normal">Remember me</label>
          </div>
          <a href="#" class="text-blue-700">Forgot password </a>
        </div>

        <div>
          <button class="button bg-primary text-white w-full" @click="handleLogin">
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
          <a href="#" class="button flex-1 flex items-center gap-2 bg-primary text-white text-sm"> <i class="icon ion-social-facebook" /> facebook </a>
          <a href="#" class="button flex-1 flex items-center gap-2 bg-sky-600 text-white text-sm"> <ion-icon :icon="logo-twitter" /> twitter </a>
          <a href="#" class="button flex-1 flex items-center gap-2 bg-black text-white text-sm"> <ion-icon
            name="logo-github"
          /> github </a>
        </div>
      </div>
    </form>
  </div>
</template>
<script setup>
const { login } = useAuth()

const data = reactive({
    email: '',
    password: '',
    loading: false
})

async function handleLogin() {
    data.loading = true
    // console.log(data)
    try {
        await login({
            email: data.email,
            password: data.password
        })
    } catch (error) {
        console.log(error)
    } finally {
        alert(data)
        console.log(data)
        data.loading = false
    }
}

// useHead

</script>