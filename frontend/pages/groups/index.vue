<script setup>

const data = ref([])
onMounted(async () => {
    const store = useGlobalAuthStore()
    const response = await $fetch('/api/group', {
        headers: {
            Authorization: `Bearer ${store.token}`,
        }
    })
    data.value = response.groups
    console.log(response)
})
</script>

<template>
    <NuxtLayout>
        <main id="site__main"
            class="2xl:ml-[--w-side]  xl:ml-[--w-side-sm] py-10 p-2.5 h-[calc(100vh-var(--m-top))] mt-[--m-top]">
            <div class="2xl:max-w-[1220px] max-w-[1065px] mx-auto h-full">
                <div class="page-heading flex flex-row justify-between align-baseline flex-wrap">
                    <h1 class="page-title">
                        Groups
                    </h1>
                    <div uk-toggle="target: #create-group-overlay" class="h-fit self-center">
                        <UButton class="bg-blue-500 h-fit">
                            <UIcon name="i-heroicons-plus" class="" /><span>New</span>
                        </UButton>
                    </div>
                </div>
                <!-- group list tabs -->
                <!-- <div id="group-tabs" class="uk-switcher"> -->
                <!-- card layout 1 -->
                <div v-if="data.length" class="flex flex-row flex-wrap overflow-scroll">
                    <GroupCard v-for="group in data" :group="group" />
                </div>
                <div class="flex w-full flex-row justify-center align-middle" v-else>
                    <div class="flex flex-col justify-center">
                        <p class="text-base text-lg">there is no group at the moment</p>
                        <div uk-toggle="target: #create-group-overlay" class="h-fit self-center">

                            <UButton class="bg-blue-500 h-fit">
                                <UIcon name="i-heroicons-plus" class="" /><span>Create new</span>
                            </UButton>
                        </div>
                    </div>

                    <!-- </div> -->
                </div>
            </div>
            <GroupCreateGroupModal></GroupCreateGroupModal>
        </main>
    </NuxtLayout>
</template>
