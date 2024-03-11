<template>
  <div class="sm:p-4 p-2.5 border-t border-gray-100 font-normal space-y-3 relative dark:border-slate-700/40">
    <Comment v-for="comment in comments" :comment="comment" />
    <button type="button" class="flex items-center gap-1.5 text-gray-500 hover:text-blue-500 mt-2">
      <ion-icon name="chevron-down-outline" class="ml-auto duration-200 group-aria-expanded:rotate-180"></ion-icon>
      More Comment
    </button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      comments: [
        {
          userAvarUrl: "assets/images/avatars/avatar-2.jpg",
          name: "Steeve",
          content: "What a beautiful photo! I love it. 😍"
        },
        {
          userAvarUrl: "assets/images/avatars/avatar-3.jpg",
          name: "Monroe",
          content: " You captured the moment.😎"
        }
      ]
    }
  },
  // props: {
  //   comments: {
  //     type: Array,
  //     required: false,
  //   }
  // }
}
</script>

<style></style>