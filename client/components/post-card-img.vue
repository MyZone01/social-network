<template>
  <!--  post image-->
  <div class="bg-white rounded-xl shadow-sm text-sm font-medium border1 dark:bg-dark2 ">

    <!-- post heading -->
    <div class="flex gap-3 sm:p-4 p-2.5 text-sm font-medium">
      <a href="timeline.html">
        <nuxt-img :src="'http://localhost:8081/' + post.userAvatarImageUrl" alt="" class="w-9 h-9 rounded-full" />
      </a>
      <div class="flex-1">
        <a href="timeline.html">
          <h4 class="text-black dark:text-white"> {{ post.userCompletName }}</h4>
        </a>
        <div class="text-xs text-gray-500 dark:text-white/80"> {{ post.createdAt }}</div>
      </div>

      <div class="-mr-1">
        <button type="button" class="button-icon w-8 h-8"> <ion-icon class="text-xl"
            name="ellipsis-horizontal"></ion-icon> </button>
        <div class="w-[245px]"
          uk-dropdown="pos: bottom-right; animation: uk-animation-scale-up uk-transform-origin-top-right; animate-out: true; mode: click">
          <nav>
            <a href="#"> <ion-icon class="text-xl shrink-0" name="bookmark-outline"></ion-icon> Add to favorites
            </a>
            <a href="#"> <ion-icon class="text-xl shrink-0" name="notifications-off-outline"></ion-icon> Mute
              Notification </a>
            <a href="#"> <ion-icon class="text-xl shrink-0" name="flag-outline"></ion-icon> Report this post </a>
            <a href="#"> <ion-icon class="text-xl shrink-0" name="share-outline"></ion-icon> Share your profile </a>
            <hr>
            <a href="#" class="text-red-400 hover:!bg-red-50 dark:hover:!bg-red-500/50">
              <ion-icon class="text-xl shrink-0" name="stop-circle-outline"></ion-icon> Unfollow </a>
          </nav>
        </div>
      </div>
    </div>
    <div class="sm:px-4 p-2.5 pt-0">
      <p class="font-normal"> {{ post.content }} </p>
    </div>
    <!-- post image -->

    <a href="#preview_modal" uk-toggle v-if="post.imageUrl">
      <div class="relative w-full lg:h-96 h-full sm:px-4">
        <nuxt-img :src="'http://localhost:8081/' + post.imageUrl" class="sm:rounded-lg w-full h-full object-cover" />
      </div>
    </a>

    <!-- post icons -->
    <div class="sm:p-4 p-2.5 flex items-center gap-4 text-xs font-semibold">

      <div class="flex items-center gap-3">
        <button type="button" class="button-icon bg-slate-200/70 dark:bg-slate-700">
          <ion-icon class="text-lg" name="chatbubble-ellipses"></ion-icon> </button>
        <span>260</span>
      </div>
      <button type="button" class="button-icon ml-auto"> <ion-icon class="text-xl"
          name="paper-plane-outline"></ion-icon>
      </button>
      <button type="button" class="button-icon"> <ion-icon class="text-xl" name="share-outline"></ion-icon>
      </button>
    </div>

    <!-- comments -->
    <!-- <div class="sm:p-4 p-2.5 border-t border-gray-100 font-normal space-y-3 relative dark:border-slate-700/40">

      <div class="flex items-start gap-3 relative">
        <a > <img src="assets/images/avatars/avatar-2.jpg" alt="" class="w-6 h-6 mt-1 rounded-full">
        </a>
        <div class="flex-1">
          <a href="timeline.html" class="text-black font-medium inline-block dark:text-white"> Steeve </a>
          <p class="mt-0.5">What a beautiful photo! I love it. 😍 </p>
        </div>
      </div>
      <div class="flex items-start gap-3 relative">
        <a href="timeline.html"> <img src="assets/images/avatars/avatar-3.jpg" alt="" class="w-6 h-6 mt-1 rounded-full">
        </a>
        <div class="flex-1">
          <a href="timeline.html" class="text-black font-medium inline-block dark:text-white"> Monroe </a>
          <p class="mt-0.5"> You captured the moment.😎 </p>
        </div>
      </div>

      <button type="button" class="flex items-center gap-1.5 text-gray-500 hover:text-blue-500 mt-2">
        <ion-icon name="chevron-down-outline" class="ml-auto duration-200 group-aria-expanded:rotate-180"></ion-icon>
        More Comment
      </button>

    </div> -->
    <Comments-snippets />

    <!-- add comment -->

    <Comment-input :postId="post.id" />

  </div>
</template>
<script>
export default {
  mounted() {
    console.log(this.post);
  },
  props: {
    post: {
      type: Object,
      required: true
    }
  }
}
</script>