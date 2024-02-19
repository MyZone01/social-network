<template>
    <aside
      class="fixed bg-gray-100 dark:bg-gray-950 top-0 left-0 z-[99] pt-[--m-top] overflow-hidden transition-transform xl:duration-500 max-xl:w-full max-xl:-translate-x-full"
    >
      <!-- sidebar inner -->
      <div
        class="p-2 max-xl:bg-white shadow-sm 2xl:w-72 sm:w-48 w-[80%] h-[calc(100vh-64px)] relative z-30 max-lg:border-r dark:max-xl:!bg-slate-700 dark:border-slate-700"
      >
        <div class="flex items-center pr-4">
          <nav id="side">
            <ul>
              <li>
                <router-link to="/feed">
                  <img src="../../assets/images/icons/home.png" alt="feeds" class="w-6">
                  <span> Feed </span>
                </router-link>
              </li>
              <li>
                <router-link to="/message">
                  <img src="../../assets/images/icons/message.png" alt="messages" class="w-5">
                  <span> messages </span>
                </router-link>
              </li>
              <li>
                <router-link to="/event">
                  <img src="../../assets/images/icons/event.png" alt="messages" class="w-6">
                  <span> event </span>
                </router-link>
              </li>
              <li>
                <router-link to="/groups">
                  <img src="../../assets/images/icons/group.png" alt="groups" class="w-6">
                  <span> Groups </span>
                </router-link>
              </li>
            </ul>
          </nav>
        </div>
      </div>
    </aside>
  </template>
  <script setup>
  </script>