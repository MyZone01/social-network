<template>
  <slot />

  <div
    class="space-y-7 text-sm text-black font-medium dark:text-white"
    uk-scrollspy="target: > *; cls: uk-animation-scale-up; delay: 100 ;repeat: true"
  >
    <h2 class="space-y-7 text-xl text-red-500 font-bold dark:text-red-500">
      {{ data.registerError }}
    </h2>
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


      <UIInput v-model="data.dateOfBirth" placeholder="01/01/2000" label="Date Of Birth" type="date" required />
      <!-- Date Of Birth -->
      <!-- <UIInput v-model="data.dateOfBirth" label="Date Of Birth" placeholder="" type="date" /> -->

      <!-- avatar image -->
      <UIImage label="Avatar Image" @image-selected="handleImageSelected" />
      <div v-if="imageUrl" class="col-span-1">
        <img :src="imageUrl" alt="Uploaded Image">
      </div>

      <!-- about me -->
      <div class="col-span-2">
        <UITextField
          v-model="data.aboutMe" label="About Me" placeholder="Short introduction about yourself ..."
          type="textarea" rows="4"
        />
      </div>

      <div class="col-span-2">
        <button class="button bg-primary text-white w-full cursor-pointer" @click="handleRegister()">
          Get
          Started
        </button>
      </div>
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
    password: '',
    confirmPassword: '',
    dateOfBirth: '',
    aboutMe: '',
    avatarImg: File,
    avatarLocalUrl: '',
    registerError: '',
    loading: false
})

const imageUrl = ref(null)

const handleImageSelected = (imageFile) => {
    // Handle the uploaded image file, like upload it to a cloud storage service

    imageUrl.value = URL.createObjectURL(imageFile)
    data.avatarLocalUrl = imageUrl.value
    data.avatarImg = imageFile
}

async function handleRegister() {
    data.registerError = ''
    data.loading = true
    try {
        const fomdata = JSON.stringify({
            firstName: data.firstName.trim(),
            lastName: data.lastName.trim(),
            email: data.email.trim(),
            nickname: data.nickname.trim(),
            password: data.password.trim(),
            repeatPassword: data.confirmPassword.trim(),
            dateOfBirth: data.dateOfBirth,
            aboutMe: data.aboutMe.trim(),
        })
        const idSession = await register({ avatarImage: data.avatarImg, data: fomdata })

        // if (idSession) {
        //     console.log('User registered successfully');
        //     await navigateTo('/')
        // }
    } catch (error) {
        data.registerError = error.statusMessage
        data.loading = false
    } finally {
        // data.firstName = ''
        // data.lastName = ''
        // data.email = ''
        // data.nickname = ''
        // data.confirmPassword = ''
        // data.password = ''
        // data.aboutMe = ''
        // data.avatarImg = ''
        data.loading = false
    }
}

</script>