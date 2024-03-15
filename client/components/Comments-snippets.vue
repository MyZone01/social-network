<template>
  <div v-if="comments.length" class="sm:p-4 p-2.5 border-t border-gray-100 font-normal space-y-3 relative dark:border-slate-700/40">
    <Comment v-for="comment in comments.slice(-2) " :comment="comment" />
    <button v-if="comments.length >2" type="button" class="flex items-center gap-1.5 text-gray-500 hover:text-blue-500 mt-2">
      <ion-icon name="chevron-down-outline" class="ml-auto duration-200 group-aria-expanded:rotate-180"></ion-icon>
      More Comment
    </button>
  </div>
</template>

<script>
export default {
  props: {
    comments: {
      type: Array,
      required: false,
    }
  }
}
</script>

<style></style>