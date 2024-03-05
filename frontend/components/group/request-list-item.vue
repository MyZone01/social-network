<template>
    <div class="px-10  h-12 w-full flex flex-row justify-between">
        <div class="h-full flex-3 bg-blue-700">
            <div class="flex flex-row gap-3">
                <img src="http://localhost:8081/uploads/default-avatar.png" />
                <div>
                    <div>{{ `${props.member.User?.firstName} ${props.member.User?.lastName}` }}</div>

                </div>
            </div>
        </div>
        <div class="h-full flex-2">
            {{ props.member?.Role }}
        </div>
        <div class="h-full flex-1">
            <div v-if="status === 'requesting'">
                <UButton @click="handleAccept" class="text-green-500">Accept</UButton>
                <UButton @click="handleDecline" class="text-red-400">Decline</UButton>
            </div>
            <div v-else-if="status === 'accepted'">
                <div class="text-blue-500">Accepted</div>
            </div>
            <div v-else>
                <div class="text-red-400">Declined</div>
            </div>

        </div>

    </div>
</template>

<script lang="ts" setup>
import { acceptJoinRequest, declneJoinRequest } from '@/composables/group/requests';

type Member = {
    firstname: String,
    lastname: String,
    memberSince: String,
    role: String

}
const props = defineProps(['member'])
const status = ref<string>(props.member.Status)


async function handleAccept(){
   const response =  await acceptJoinRequest(props.member.GroupID,props.member.ID)
   if (response) {
    const data = JSON.parse(response.data)
    status.value = data.Status
   }
}

async function handleDecline(){
    const response =  await declneJoinRequest(props.member.GroupID,props.member.ID)
   if (response) {
    const data = JSON.parse(response.data)
    status.value = data.Status
   }
}

</script>

<style></style>