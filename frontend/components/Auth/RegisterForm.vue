<template>
    <slot></slot>

    <div class="space-y-7 text-sm text-black font-medium dark:text-white"
        uk-scrollspy="target: > *; cls: uk-animation-scale-up; delay: 100 ;repeat: true">

        <h2 class="space-y-7 text-xl text-red-500 font-bold dark:text-red-500">{{ data.registerError }}</h2>
        <div class="grid grid-cols-2 gap-4 gap-y-7">

            <!-- last name -->
            <UIInput v-model="data.firstName" label="Last Name" placeholder="Khoulé" />

            <!-- first name -->
            <UIInput v-model="data.lastName" label="First Name" placeholder="Serigne Mamadou" />

            <!-- email -->
            <UIInput v-model="data.email" label="Email" placeholder="mail@social.net" />

            <!-- nickname -->
            <UIInput v-model="data.nickname" label="Nickname" placeholder="Modou" />

            <!-- password -->
            <UIInput v-model="data.password" label="Password" placeholder="********" type="password" />

            <!-- confirm password -->
            <UIInput v-model="data.confirmPassword" label="Confirm Password" placeholder="********" type="password" />

            <!-- avatar image -->
            <UIImage @image-selected="handleImageSelected" />
            <div v-if="imageUrl" class="col-span-2">
                <img :src="imageUrl" alt="Uploaded Image">
            </div>

            <!-- about me -->
            <UITextField v-model="data.aboutMe" label="About Me" placeholder="Short introduction about yourself ..."
                type="textarea" rows="4" />

            <div class="col-span-2">
                <label class="inline-flex items-center" id="rememberme">
                    <input type="checkbox" id="accept-terms" class="!rounded-md accent-red-800" />
                    <span class="ml-2">you agree to our <a href="#" class="text-blue-700 hover:underline">terms of use </a>
                    </span>
                </label>
            </div>

            <div class="col-span-2">
                <button @click="handleRegister()" class="button bg-primary text-white w-full">Get Started</button>
            </div>
        </div>

        <div class="text-center flex items-center gap-6">
            <hr class="flex-1 border-slate-200 dark:border-slate-800">
            Or continue with
            <hr class="flex-1 border-slate-200 dark:border-slate-800">
        </div>
    </div>
</template>

<script setup>
const { register } = useAuth()

const data = reactive({
    firstName: '',
    lastName: '',
    email: '',
    nickname: '',
    confirmPassword: '',
    password: '',
    aboutMe: '',
    avatarImg: '',
    registerError: '',
    loading: false
})

const imageUrl = ref(null)

const handleImageSelected = (imageFile) => {
    // Handle the uploaded image file, like upload it to a cloud storage service

    imageUrl.value = URL.createObjectURL(imageFile)
    data.avatarImg = imageFile
}

async function handleRegister() {
    data.loading = true
    try {
        const idSession = await register({
            firstName: data.firstName.trim(),
            lastName: data.lastName.trim(),
            email: data.email.trim(),
            nickname: data.nickname.trim(),
            password: data.password.trim(),
            repeatPassword: data.confirmPassword.trim(),
            aboutMe: data.aboutMe.trim(),
            avatarImg: data.avatarImg,
        })
        if (idSession) {
            await navigateTo('/')
        }
    } catch (error) {
        data.registerError = error.statusMessage
        setTimeout(() => {
            data.registerError = ''
        }, 2000)
        data.loading = false
    } finally {
        data.firstName = ''
        data.lastName = ''
        data.email = ''
        data.nickname = ''
        data.confirmPassword = ''
        data.password = ''
        data.aboutMe = ''
        data.avatarImg = ''
        data.loading = false
    }
}

</script>