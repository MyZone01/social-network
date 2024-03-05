<template>
    <div class="px-10  h-12 w-full flex flex-row justify-between">
        <div class="h-full flex-2 bg-blue-700">
            <div class="flex flex-row gap-3">
                <img src="http://localhost:8081/uploads/default-avatar.png" />
                <div>
                    <div>{{ `${props.member.User?.firstName} ${props.member.User?.lastName}` }}</div>

                </div>
            </div>
        </div>
        <div class="h-full ">
            {{ props.member?.Role }}
        </div>
        <div class="h-full">
            <UButton  v-if="props.member?.Role !== 'admin'"  class="text-red-400">Delete</UButton>
        </div>

    </div>
</template>

<script lang="ts" setup>
type Member = {
    firstname: String,
    lastname: String,
    memberSince: String,
    role: String

}
const props = defineProps(['member'])

onMounted(() => {
    console.log(props.member);

})

</script>

<style></style>