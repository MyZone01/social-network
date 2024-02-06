<template>
  <aside
    class="fixed bg-white/80 dark:bg-gray-950 top-0 left-0 z-[99] pt-[--m-top] overflow-hidden transition-transform xl:duration-500 max-xl:w-full max-xl:-translate-x-full">
    <!-- sidebar inner -->
    <div
      class="p-2 max-xl:bg-white shadow-sm 2xl:w-72 sm:w-48 w-[80%] h-[calc(100vh-64px)] relative z-30 max-lg:border-r dark:max-xl:!bg-slate-700 dark:border-slate-700">
      <div class="flex items-center pr-4">
        <nav id="side">
          <ul>
            <li>
              <a href="feed.html">
                <img src="assets/images/icons/home.png" alt="feeds" class="w-6">
                <span> Feed </span>
              </a>
            </li>
            <li>
              <a href="messages.html">
                <img src="assets/images/icons/message.png" alt="messages" class="w-5">
                <span> messages </span>
              </a>
            </li>
            <li>
              <a href="event.html">
                <img src="assets/images/icons/event.png" alt="messages" class="w-6">
                <span> event </span>
              </a>
            </li>
            <li>
              <a href="groups.html">
                <img src="assets/images/icons/group.png" alt="groups" class="w-6">
                <span> Groups </span>
              </a>
            </li>
          </ul>
        </nav>
      </div>
    </div>
  </aside>
</template>
<style scoped>
#side {
  @apply text-sm font-medium capitalize tracking-wide text-gray-900 dark:text-white;
  --tw-text-opacity: 1;
}

#side a {
  @apply flex items-center gap-4 px-4 p-3 rounded-xl;
}

#side a:hover {
  @apply bg-gray-900 dark:bg-white/10;
}
</style>