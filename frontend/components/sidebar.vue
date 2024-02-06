<template>
  <div
    id="site__sidebar"
    class="fixed  top-0 left-0 z-[99] pt-[--m-top] overflow-hidden transition-transform xl:duration-500 max-xl:w-full max-xl:-translate-x-full"
  >
    <!-- sidebar inner -->
    <div
      class="p-2 max-xl:bg-white shadow-sm 2xl:w-72 sm:w-64 w-[80%] h-[calc(100vh-64px)] relative z-30 max-lg:border-r dark:max-xl:!bg-slate-700 dark:border-slate-700"
    >
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
  </div>
</template>
<style>

#side > :not([hidden]) ~ :not([hidden]) {
  --tw-space-y-reverse: 0;
  margin-top: calc(0.125rem * calc(1 - var(--tw-space-y-reverse)));
  margin-bottom: calc(0.125rem * var(--tw-space-y-reverse));
}

#side {
  font-size: 14px;
  font-weight: 500;
  text-transform: capitalize;
  letter-spacing: 0.025em;
  --tw-text-opacity: 1;
  color: rgb(0 0 0 / var(--tw-text-opacity));
}

@media (prefers-color-scheme: dark) {
  #side {
    --tw-text-opacity: 1;
    color: rgb(255 255 255 / var(--tw-text-opacity));
  }
}

#side a {
  display: flex;
  align-items: center;
  gap: 1rem;
  border-radius: 0.75rem;
  padding: 0.75rem;
  padding-left: 1rem;
  padding-right: 1rem;
}

@media (prefers-color-scheme: dark) {
  #side a:hover {
    --tw-bg-opacity: 1;
    background-color: rgb(30 41 59 / var(--tw-bg-opacity));
  }
}

#side a:hover {
  --tw-bg-opacity: 1;
  background-color: rgb(var(--color-secondery) / var(--tw-bg-opacity));
}

#side a svg {
  height: 1.5rem !important;
  width: 1.5rem !important;
}

#side a.active svg ,
    #side a.uk-active svg ,
    #side .active  a ,
    #side a.uk-active {
  --tw-bg-opacity: 1;
  background-color: rgb(var(--color-secondery) / var(--tw-bg-opacity));
  --tw-text-opacity: 1;
  color: rgb(var(--color-primary) / var(--tw-text-opacity));
}

@media (prefers-color-scheme: dark) {
  #side a.active svg ,
    #side a.uk-active svg ,
    #side .active  a ,
    #side a.uk-active {
    --tw-text-opacity: 1;
    color: rgb(255 255 255 / var(--tw-text-opacity));
  }
}

#site__sidebar .simplebar-scrollbar {
  width: 7px;
}

</style>