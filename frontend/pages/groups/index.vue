<script setup>

const data = ref([])
onMounted(async () => {
    const store = useGlobalAuthStore()
    const response = await $fetch('/api/group', {
        headers: {
            Authorization: `Bearer ${store.token}`,
        }
    })
    data.value = response.groups.slice(0, 4)
    console.log(response)
})
</script>

<template>
    <NuxtLayout>
        <main id="site__main"
            class="2xl:ml-[--w-side]  xl:ml-[--w-side-sm] py-10 p-2.5 h-[calc(100vh-var(--m-top))] mt-[--m-top]">
            <div class="2xl:max-w-[1220px] max-w-[1065px] mx-auto">
                <div class="page-heading">
                    <h1 class="page-title">
                        Groups
                    </h1>
                    <nav class="nav__underline flex flex-row justify-between align-baseline flex-wrap">
                        <ul class="group"
                            uk-switcher="connect: #group-tabs ; animation: uk-animation-slide-right-medium, uk-animation-slide-left-medium">
                            <li> <a href="#"> Suggestions </a> </li>
                            <li> <a href="#"> Popular </a> </li>
                            <li> <a href="#"> My groups </a> </li>
                        </ul>
                        <div uk-toggle="target: #create-group-overlay" class="h-fit self-center">
                            <UButton class="bg-blue-500 h-fit">
                                <UIcon name="i-heroicons-plus" class="" /><span>New</span>
                            </UButton>
                        </div>
                    </nav>
                </div>
                <!-- group list tabs -->
                <div id="group-tabs" class="uk-switcher">
                    <!-- card layout 1 -->
                    <div v-if="data.length" class="flex flex-row gap-2 overflow-hidden">
                        <GroupCard v-for="group in data" :group="group" />
                    </div>
                    <div class="flex w-full flex-row justify-center align-middle" v-else >
                        <div class="flex flex-col justify-center">
                           <p class="text-base text-lg">there is no group at the moment</p>
                           <div uk-toggle="target: #create-group-overlay" class="h-fit self-center">

                        <UButton class="bg-blue-500 h-fit">
                            <UIcon name="i-heroicons-plus" class="" /><span>Create new</span>
                        </UButton> 
                        </div>
                    </div>

                    </div>
                </div>
            </div>
            <CreateGroupModal></CreateGroupModal>
        </main>
    </NuxtLayout>
</template>
