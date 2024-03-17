<template>
  <div class="px-10  h-12 w-full flex flex-row justify-between">
    <div class="h-full  bg-blue-700">
      <div class="flex flex-row gap-3">
        <img src="http://localhost:8081/uploads/default-avatar.png" class="w-10 h-8"/>
        <div>
          <div>{{ `${props.member.User?.firstName} ${props.member.User?.lastName}` }}</div>
        </div>
      </div>
    </div>
    <div class="h-full  text-left">
      {{ props.member?.Role }}
    </div>
    <div class="h-full  w-max">
      <UButton v-if="props.member?.Role !== 'admin' && props.isAdmin" class="w-10 rounded-full text-red-500 border-4 border-white dark:border-slate-800">
        Delete
      </UButton>
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
const props = defineProps(['member','isAdmin'])



</script>

<style></style>