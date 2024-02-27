<template>
    <div class="overflow-visible">
        <slot></slot>

        <div class="space-y-7 text-sm text-black font-medium dark:text-white"
            uk-scrollspy="target: > *; cls: uk-animation-scale-up; delay: 100 ;repeat: true">
            <h2 class="space-y-7 text-xl text-red-500 font-bold dark:text-red-500">{{ data.loginError }}</h2>

            <UIInput v-model="data.email" label="Email" placeholder="mail@social.net" required />

            <UIInput v-model="data.password" label="Password" placeholder="********" type="password" required />

            <div>
                <button @click="handleLogin()" class="button bg-primary text-white w-full cursor-pointer">Sign in</button>
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
    loading: false,
})
// const loginError = ''
async function handleLogin() {
    data.loginError = ''
    data.loading = true
    try {
        const idSession = await login({
            email: data.email.trim(),
            password: data.password.trim()
        })
        
        // await navigateTo('/')
    } catch (error) {
        data.loginError = error.statusMessage
        data.loading = false
    } finally {
        data.loading = false
        // data.email = ''
        // data.password = ''
    }
}
</script>