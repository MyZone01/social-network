<template>
  <div class="flex items-start gap-3 relative">
    <a> <nuxt-img  :src="'http://localhost:8081/'+  comment.userAvatarImageUrl" alt="" class="w-6 h-6 mt-1 rounded-full"/>
    </a>
    <div class="flex-1">
      <a class="text-black font-medium inline-block dark:text-white"> {{ comment.userCompleteName }} </a>
      <p class="mt-0.5"> {{ comment.content }} </p>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    comment: {
      type: Object,
      required: true
    }
  }
}
</script>

<style></style>
