<template>
    <div id="site__sidebar"
        class="fixed top-0 left-0 z-[99] pt-[--m-top] overflow-hidden transition-transform xl:duration-500 max-xl:w-full max-xl:-translate-x-full">

        <!-- sidebar inner -->
        <div
            class="p-2 max-xl:bg-white shadow-sm 2xl:w-72 sm:w-64 w-[80%] h-[calc(100vh-64px)] relative z-30 max-lg:border-r dark:max-xl:!bg-slate-700 dark:border-slate-700">

            <div class="pr-4" data-simplebar>

                <nav id="side">

                    <ul>
                        <li :class="{ active: $route.path === '/' }">
                            <a href="/">
                                <img src="assets/images/icons/home.png" alt="feeds" class="w-6">
                                <span> Feed </span>
                            </a>
                        </li>
                        <li :class="{ active: $route.path === '/messages' }">
                            <a href="/messages">
                                <img src="assets/images/icons/message.png" alt="messages" class="w-5">
                                <span> messages </span>
                            </a>
                        </li>
                        <li :class="{ active: $route.path === '/events' }">
                            <a href="/events">
                                <img src="assets/images/icons/event.png" alt="messages" class="w-6">
                                <span> events </span>
                            </a>
                        </li>
                        <li :class="{ active: $route.path === '/groups' }">
                            <a href="/groups">
                                <img src="assets/images/icons/group.png" alt="groups" class="w-6">
                                <span> Groups </span>
                            </a>
                        </li>
                    </ul>

                    <!-- <button type="button"
                        class="flex items-center gap-4 py-2 px-4 w-full font-medium text-sm text-black dark:text-white"
                        uk-toggle="target: #show__more; cls: !hidden uk-animation-fade">
                        <svg class="bg-gray-200 rounded-full w-6 h-6 dark:bg-slate-700" fill="currentColor"
                            viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                            <path fill-rule="evenodd"
                                d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"
                                clip-rule="evenodd"></path>
                        </svg>
                        <span id="show__more"> See More </span>
                        <span class="!hidden" id="show__more"> See Less </span>
                    </button> -->

                </nav>
                <nav>

                    
                    <div class="font-medium text-sm text-black border-t pt-3 mt-2 dark:text-white dark:border-slate-800">
                    <div class="px-3 pb-2 text-sm font-medium">
                        <div class="text-black dark:text-white">Shortcut</div>
                    </div>
                    <a href="#">
                        <div class="flex items-center gap-2 p-3 px-4 rounded-xl hover:bg-secondery">
                            <img src="assets/images/avatars/avatar-2.jpg" alt="" class="w-6 rounded-full object-cover">
                            <div> Marin Gray</div>
                        </div>
                    </a>
                    <a href="#">
                        <div class="flex items-center gap-2 p-3 px-4 rounded-xl hover:bg-secondery">
                            <img src="assets/images/avatars/avatar-7.jpg" alt="" class="w-6 rounded-full object-cover">
                            <div> Alexa Stella</div>
                        </div>
                    </a>
                    <a href="#">
                        <div class="flex items-center gap-2 p-3 px-4 rounded-xl hover:bg-secondery">
                            <img src="assets/images/avatars/avatar-3.jpg" alt="" class="w-6 rounded-full object-cover">
                            <div> Sarah Ali</div>
                        </div>
                    </a>
                </div>

                <br>
                <br>
                <div v-if="avatarImage" id="toast-warning"
                    class="flex items-center w-full max-w-xs p-4 text-gray-500 bg-white rounded-lg shadow dark:text-gray-400 dark:bg-gray-800 space-x-3"
                    role="alert">
                    <div
                        class="inline-flex items-center justify-center flex-shrink-0 w-8 h-8 text-orange-500 bg-orange-100 rounded-lg dark:bg-orange-700 dark:text-orange-200">
                        <svg class="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor"
                            viewBox="0 0 20 20">
                            <path
                                d="M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5ZM10 15a1 1 0 1 1 0-2 1 1 0 0 1 0 2Zm1-4a1 1 0 0 1-2 0V6a1 1 0 0 1 2 0v5Z" />
                        </svg>
                        <!-- <ion-icon class="sr-only" name="alert-outline"></ion-icon> -->
                    </div>
                    <div class="ms-3 text-sm text-red-600 font-bold"> Please Upload an Avatar </div>
                    <button type="button"
                        class="ms-auto -mx-1.5 -my-1.5 bg-white text-gray-400 hover:text-gray-900 rounded-lg focus:ring-2 focus:ring-gray-300 p-1.5 hover:bg-gray-100 inline-flex items-center justify-center h-8 w-8 dark:text-gray-500 dark:hover:text-white dark:bg-gray-800 dark:hover:bg-gray-700"
                        data-dismiss-target="#toast-warning" aria-label="Close">
                        <!-- <ion-icon class="sr-only" name="close-circle-outline"></ion-icon> -->
                        <!-- <svg class="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none"
                            viewBox="0 0 14 14">
                            <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6" />
                        </svg> -->
                    </button>
                </div>

            </nav>
                <nav id="side"
                    class="font-medium text-sm text-black border-t pt-3 mt-2 dark:text-white dark:border-slate-800 absolute inset-x-0 bottom-0">
                    <!-- <div class="px-3 pb-2 text-sm font-medium">
                            <div class="text-black dark:text-white">Pages</div>
                        </div> -->

                    <ul class="mt-2 -space-y-2" uk-nav="multiple: true">

                        <li>
                            <a href="/settings">
                                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                                    stroke="currentColor" class="w-4 h-4">
                                    <path stroke-linecap="round" stroke-linejoin="round"
                                        d="M9.594 3.94c.09-.542.56-.94 1.11-.94h2.593c.55 0 1.02.398 1.11.94l.213 1.281c.063.374.313.686.645.87.074.04.147.083.22.127.324.196.72.257 1.075.124l1.217-.456a1.125 1.125 0 011.37.49l1.296 2.247a1.125 1.125 0 01-.26 1.431l-1.003.827c-.293.24-.438.613-.431.992a6.759 6.759 0 010 .255c-.007.378.138.75.43.99l1.005.828c.424.35.534.954.26 1.43l-1.298 2.247a1.125 1.125 0 01-1.369.491l-1.217-.456c-.355-.133-.75-.072-1.076.124a6.57 6.57 0 01-.22.128c-.331.183-.581.495-.644.869l-.213 1.28c-.09.543-.56.941-1.11.941h-2.594c-.55 0-1.02-.398-1.11-.94l-.213-1.281c-.062-.374-.312-.686-.644-.87a6.52 6.52 0 01-.22-.127c-.325-.196-.72-.257-1.076-.124l-1.217.456a1.125 1.125 0 01-1.369-.49l-1.297-2.247a1.125 1.125 0 01.26-1.431l1.004-.827c.292-.24.437-.613.43-.992a6.932 6.932 0 010-.255c.007-.378-.138-.75-.43-.99l-1.004-.828a1.125 1.125 0 01-.26-1.43l1.297-2.247a1.125 1.125 0 011.37-.491l1.216.456c.356.133.751.072 1.076-.124.072-.044.146-.087.22-.128.332-.183.582-.495.644-.869l.214-1.281z" />
                                    <path stroke-linecap="round" stroke-linejoin="round"
                                        d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                                </svg>
                                <span> Settings </span>
                            </a>
                        </li>
                        <li>
                            <a @click="handleLogout()" class="cursor-pointer">
                                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                                    stroke="currentColor" class="w-4 h-4">
                                    <path stroke-linecap="round" stroke-linejoin="round"
                                        d="M15.75 9V5.25A2.25 2.25 0 0013.5 3h-6a2.25 2.25 0 00-2.25 2.25v13.5A2.25 2.25 0 007.5 21h6a2.25 2.25 0 002.25-2.25V15m3 0l3-3m0 0l-3-3m3 3H9" />
                                </svg>
                                <span> Logout </span>
                            </a>
                        </li>
                        <!-- <li class="uk-parent">
                                <a href="#" class="group">
                                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                                        stroke-width="1.5" stroke="currentColor" class="w-4 h-4">
                                        <path stroke-linecap="round" stroke-linejoin="round"
                                            d="M2.25 8.25h19.5M2.25 9h19.5m-16.5 5.25h6m-6 2.25h3m-3.75 3h15a2.25 2.25 0 002.25-2.25V6.75A2.25 2.25 0 0019.5 4.5h-15a2.25 2.25 0 00-2.25 2.25v10.5A2.25 2.25 0 004.5 19.5z" />
                                    </svg>
                                    <span> Development </span>
                                    <ion-icon name="ChevronDown"
                                        class="text-base ml-auto duration-200 group-aria-expanded:rotate-180"></ion-icon>
                                </a>
                                <ul class="pl-10 my-1 space-y-0 text-sm">
                                    <li><a href="components.html" class="!py-2 !rounded -md">Elements</a></li>
                                    <li><a href="components.html" class="!py-2 !rounded -md">Components</a></li>
                                    <li><a href="components.html" class="!py-2 !rounded -md">Icons</a></li>
                                </ul>
                            </li> -->

                    </ul>

                </nav>


                <!-- <div class="text-xs font-medium flex flex-wrap gap-2 gap-y-0.5 p-2 mt-2">
                        <a href="#" class="hover:underline">About</a>
                        <a href="#" class="hover:underline">Blog </a>
                        <a href="#" class="hover:underline">Careers</a>
                        <a href="#" class="hover:underline">Support</a>
                        <a href="#" class="hover:underline">Contact Us </a>
                        <a href="#" class="hover:underline">Developer</a>
                    </div> -->

            </div>

        </div>

        <!-- sidebar overly -->
        <div id="site__sidebar__overly" class="absolute top-0 left-0 z-20 w-screen h-screen xl:hidden backdrop-blur-sm"
            uk-toggle="target: #site__sidebar ; cls :!-translate-x-0">
        </div>
    </div>
</template>
<script setup>
import { useGlobalAuthStore } from '@/stores/useGlobalStateAuthStore';

const store = useGlobalAuthStore();
const props = defineProps({
    modelValue: {
        type: String,
        required: true
    },
})


const avatarImage = store.user.avatarImage === "uploads/default-avatar.png" ? true : false

function handleLogout() {
    store.logout()

    navigateTo('/auth')
}
</script>
@/stores/useGlobalStateAuthStore