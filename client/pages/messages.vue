<template>
  <main id="site__main" class="2xl:ml-[--w-side]  xl:ml-[--w-side-sm] p-2.5 h-[calc(100vh-var(--m-top))] mt-[--m-top]">
    <nuxt-child />
    <div class="relative overflow-hidden border -m-2.5 dark:border-slate-700">
      <div class="flex bg-white dark:bg-dark2">
        <!-- sidebar -->
        <div class="md:w-[360px] relative border-r dark:border-slate-700">
          <div id="side-chat"
            class="top-0 left-0 max-md:fixed max-md:w-5/6 max-md:h-screen bg-white z-50 max-md:shadow max-md:-translate-x-full dark:bg-dark2">
            <!-- heading title -->
            <div class="p-4 border-b dark:border-slate-700">
              <div class="flex mt-2 items-center justify-between">
                <h2 class="text-2xl font-bold text-black ml-1 dark:text-white">
                  Chats
                </h2>

                <!-- right action buttons -->
                <div class="flex items-center gap-2.5">
                  <button class="group">
                    <ion-icon :icon="ioniconsSettingsOutline" class="text-2xl flex group-aria-expanded:rotate-180" />
                  </button>
                  <div class="md:w-[270px] w-full"
                    uk-dropdown="pos: bottom-left; offset:10; animation: uk-animation-slide-bottom-small">
                    <nav>
                      <a href="#"> <ion-icon class="text-2xl shrink-0 -ml-1" :icon="ioniconsCheckmarkOutline" />
                        Mark all as read </a>
                      <a href="#"> <ion-icon class="text-2xl shrink-0 -ml-1" :icon="ioniconsNotificationsOutline" />
                        notifications setting </a>
                      <a href="#"> <ion-icon class="text-xl shrink-0 -ml-1" :icon="ioniconsVolumeMuteOutline" />
                        Mute notifications </a>
                    </nav>
                  </div>

                  <button class="">
                    <ion-icon :icon="ioniconsCheckmarkCircleOutline" class="text-2xl flex" />
                  </button>

                  <!-- mobile toggle menu -->
                  <button type="button" class="md:hidden"
                    uk-toggle="target: #side-chat ; cls: max-md:-translate-x-full">
                    <ion-icon :icon="ioniconsChevronDownOutline" />
                  </button>
                </div>
              </div>

              <!-- search -->
              <div class="relative mt-4">
                <div class="absolute left-3 bottom-1/2 translate-y-1/2 flex">
                  <ion-icon :icon="ioniconsSearch" class="text-xl" />
                </div>
                <input type="text" placeholder="Search" class="w-full !pl-10 !py-2 !rounded-lg">
              </div>
            </div>


            <!-- users list -->
            <ul
              class="uk-subnav uk-subnav-pill space-y-2 p-2 overflow-y-auto md:h-[calc(100vh-204px)] h-[calc(100vh-130px)]"
              uk-switcher="animation: uk-animation-slide-left-medium, uk-animation-slide-right-medium">
              <li v-for="user in users" :key="user.id" @click="selectUser(user)"
                class="relative flex items-center gap-4 p-2 duration-200 rounded-xl hover:bg-secondery"
                id="{{ user.id }}">
                <a href="#" class="relative w-14 h-14 shrink-0">
                  <nuxt-img v-if="user && user.avatarImage" :src="'http://localhost:8081/' + user.avatarImage" alt=""
                    class="object-cover w-full h-full rounded-full" />

                  <div
                    class="w-4 h-4 absolute bottom-0 right-0  bg-green-500 rounded-full border border-white dark:border-slate-800" />
                </a>
                <a href="#" class="flex-1 min-w-0">
                  <div class="flex items-center gap-2 mb-1.5">
                    <div class="mr-auto text-sm text-black dark:text-white font-medium">{{ user.firstName }} {{
                      user.lastName }}</div>
                    <div class="text-xs font-light text-gray-500 dark:text-white/70"></div>
                  </div>
                  <div class="font-medium overflow-hidden text-ellipsis text-sm whitespace-nowrap">
                  </div>
                </a>
              </li>
            </ul>
            <!-- <ul  class="uk-switcher uk-margin flex-1"> -->
            <!-- </ul> -->
          </div>
          <!-- overly -->
          <div id="side-chat"
            class="bg-slate-100/40 backdrop-blur w-full h-full dark:bg-slate-800/40 z-40 fixed inset-0 max-md:-translate-x-full md:hidden"
            uk-toggle="target: #side-chat ; cls: max-md:-translate-x-full" />


        </div>

        <!-- message center -->
        <div v-for="user in users" :key="user.id">
          <div v-if="receiverid === user.id" class="flex-1">
            <div
              class="flex items-center justify-between gap-2 w- px-6 py-3.5 z-10 border-b dark:border-slate-700 uk-animation-sdivde-top-medium">
              <!-- Insérez ici les informations de l'utilisateur -->

              <div class="flex items-center sm:gap-4 gap-2">
                <!-- toggle for mobile -->
                <button type="button" class="md:hidden" uk-toggle="target: #side-chat ; cls: max-md:-translate-x-full">
                  <ion-icon :icon="ioniconsChevronBackOutline" class="text-2xl -ml-4" />
                </button>

                <div class="relative cursor-pointer max-md:hidden" uk-toggle="target: .rightt ; cls: hidden">
                  <nuxt-img v-if="user && user.avatarImage" :src="'http://localhost:8081/' + user.avatarImage"
                    class="w-8 h-8 rounded-full shadow" alt="" />
                  <div class="w-2 h-2 bg-teal-500 rounded-full absolute right-0 bottom-0 m-px" />
                </div>
                <div class="cursor-pointer" uk-toggle="target: .rightt ; cls: hidden">
                  <div class="text-base font-bold">
                    {{ user.firstName }} {{ user.lastName }}
                  </div>
                  <div class="text-xs text-green-500 font-semibold">
                    Online
                  </div>
                </div>
              </div>
              <div class="flex items-center gap-2">
                <button type="button" class="button__ico">
                  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="w-6 h-6">
                    <path fill-rule="evenodd"
                      d="M2 3.5A1.5 1.5 0 013.5 2h1.148a1.5 1.5 0 011.465 1.175l.716 3.223a1.5 1.5 0 01-1.052 1.767l-.933.267c-.41.117-.643.555-.48.95a11.542 11.542 0 006.254 6.254c.395.163.833-.07.95-.48l.267-.933a1.5 1.5 0 011.767-1.052l3.223.716A1.5 1.5 0 0118 15.352V16.5a1.5 1.5 0 01-1.5 1.5H15c-1.149 0-2.263-.15-3.326-.43A13.022 13.022 0 012.43 8.326 13.019 13.019 0 012 5V3.5z"
                      clip-rule="evenodd" />
                  </svg>
                </button>
                <button type="button" class="hover:bg-slate-100 p-1.5 rounded-full">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                    stroke="currentColor" class="w-6 h-6">
                    <path stroke-linecap="round"
                      d="M15.75 10.5l4.72-4.72a.75.75 0 011.28.53v11.38a.75.75 0 01-1.28.53l-4.72-4.72M4.5 18.75h9a2.25 2.25 0 002.25-2.25v-9a2.25 2.25 0 00-2.25-2.25h-9A2.25 2.25 0 002.25 7.5v9a2.25 2.25 0 002.25 2.25z" />
                  </svg>
                </button>
                <button type="button" class="hover:bg-slate-100 p-1.5 rounded-full"
                  uk-toggle="target: .rightt ; cls: hidden">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                    stroke="currentColor" class="w-6 h-6">
                    <path stroke-linecap="round" stroke-linejoin="round"
                      d="M11.25 11.25l.041-.02a.75.75 0 011.063.852l-.708 2.836a.75.75 0 001.063.853l.041-.021M21 12a9 9 0 11-18 0 9 9 0 0118 0zm-9-3.75h.008v.008H12V8.25z" />
                  </svg>
                </button>
              </div>
            </div>
            <div v-if="receiverid === user.id"
              class="w-full p-5 py-10 overflow-y-auto md:h-[calc(100vh-204px)] h-[calc(100vh-195px)]">
              <div class="py-10 text-center text-sm lg:pt-8">
                <nuxt-img v-if="user && user.avatarImage" :src="'http://localhost:8081/' + user.avatarImage"
                  class="w-24 h-24 rounded-full mx-auto mb-3" alt="" />
                <div class="mt-8">
                  <div class="md:text-xl text-base font-medium text-black dark:text-white">
                    {{ user.firstName }} {{ user.lastName }}
                  </div>
                  <div class="text-gray-500 text-sm   dark:text-white/80">
                    @{{ user.nickname }}
                  </div>
                </div>
                <div class="mt-3.5">
                  <a href="#" class="inline-block rounded-lg px-4 py-1.5 text-sm font-semibold bg-secondery">View
                    profile</a>
                </div>
              </div>
              <!-- <div class="px-4 py-2 rounded-[20px] max-w-sm bg-secondery">
                {{ chats.value }}
              </div> -->
              <div id="messages" class="text-sm font-medium space-y-6">
                <!-- received -->
                <div v-if="chats">
                  <div class="inline-block rounded-full px-3.5 py-0.5 text-sm font-semibold bg-secondery" style="display: flex; justify-content: center;"> {{ chats[0].Day }}</div>
                  <div v-for="chat in chats" :key="chat.ID">
                    <div v-if="chat.SenderID === user.id" class="flex gap-3">
                      <nuxt-img v-if="user && user.avatarImage" :src="'http://localhost:8081/' + user.avatarImage"
                        class="w-9 h-9 rounded-full shadow" />
                      <div class="px-4 py-2 rounded-[20px] max-w-sm bg-secondery">
                        {{ chat.Content }}
                      </div>
                    </div>
                    <div v-if="chat.SenderID === user.id" class="inline-block rounded-full px-3.5 py-0.5 text-sm font-semibold bg-secondery" >{{ chat.Hour }} </div>
                    <div v-if="chat.SenderID === currentUser.id" class="flex gap-2 flex-row-reverse items-end">
                      <nuxt-img v-if="user && user.avatarImage"
                        :src="'http://localhost:8081/' + currentUser.avatarImage" class="w-9 h-9 rounded-full shadow" />
                      <div
                        class="px-4 py-2 rounded-[20px] max-w-sm bg-gradient-to-tr from-sky-500 to-blue-500 text-white shadow">
                        {{ chat.Content }}
                      </div>

                    </div>
                    <div v-if="chat.SenderID === currentUser.id" class="inline-block rounded-full px-3.5 py-0.5 text-sm font-semibold bg-secondery " style="margin-left: 75%;">{{ chat.Hour }} </div>

                  </div>
                </div>
              </div>
            </div>
            <div class="flex items-center md:gap-4 gap-2 md:p-3 p-2 overflow-hidden">
              <div id="message__wrap" class="flex items-center gap-2 h-full dark:text-white -mt-1.5">
                <button type="button" class="shrink-0">
                  <ion-icon class="text-3xl flex" :icon="ioniconsAddCircleOutline" />
                </button>
                <div
                  class="dropbar pt-36 h-60 bg-gradient-to-t via-white from-white via-30% from-30% dark:from-slate-900 dark:via-900"
                  uk-drop="stretch: x; target: #message__wrap ;animation:  slide-bottom ;animate-out: true; pos: top-left; offset:10 ; mode: click ; duration: 200">
                  <div class="sm:w-full p-3 flex justify-center gap-5"
                    uk-scrollspy="target: > button; cls: uk-animation-slide-bottom-small; delay: 100;repeat:true">
                    <button type="button"
                      class="bg-sky-50 text-sky-600 border border-sky-100 shadow-sm p-2.5 rounded-full shrink-0 duration-100 hover:scale-[1.15] dark:bg-dark3 dark:border-0">
                      <ion-icon class="text-3xl flex" :icon="ioniconsImage" />
                    </button>
                    <button type="button"
                      class="bg-green-50 text-green-600 border border-green-100 shadow-sm p-2.5 rounded-full shrink-0 duration-100 hover:scale-[1.15] dark:bg-dark3 dark:border-0">
                      <ion-icon class="text-3xl flex" :icon="ioniconsImages" />
                    </button>
                    <button type="button"
                      class="bg-pink-50 text-pink-600 border border-pink-100 shadow-sm p-2.5 rounded-full shrink-0 duration-100 hover:scale-[1.15] dark:bg-dark3 dark:border-0">
                      <ion-icon class="text-3xl flex" :icon="ioniconsDocumentText" />
                    </button>
                    <button type="button"
                      class="bg-orange-50 text-orange-600 border border-orange-100 shadow-sm p-2.5 rounded-full shrink-0 duration-100 hover:scale-[1.15] dark:bg-dark3 dark:border-0">
                      <ion-icon class="text-3xl flex" :icon="ioniconsGift" />
                    </button>
                  </div>
                </div>

                <button type="button" class="shrink-0">
                  <ion-icon class="text-3xl flex" :icon="ioniconsHappyOutline" />
                </button>
                <!-- <div class="dropbar p-2"
                    uk-drop="stretch: x; target: #message__wrap ;animation: uk-animation-scale-up uk-transform-origin-bottom-left ;animate-out: true; pos: top-left ; offset:2; mode: click ; duration: 200 ">
                    <div class="sm:w-60 bg-white shadow-lg border rounded-xl  pr-0 dark:border-slate-700 dark:bg-dark3">
                      <h4 class="text-sm font-semibold p-3 pb-0">
                        Send Imogi
                      </h4>

                      <div class="grid grid-cols-5 overflow-y-auto max-h-44 p-3 text-center text-xl">
                        <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200">
                          😊
                        </div>
                        <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200">
                          🤩
                        </div>
                        <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200">
                          😎
                        </div>
                        <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200">
                          🥳
                        </div>
                        <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200">
                          😂
                        </div>
                        <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200">
                          🥰
                        </div>
                        <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200">
                          😡
                        </div>
                        <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200">
                          😊
                        </div>
                        <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200">
                          🤩
                        </div>
                        <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200">
                          😎
                        </div>
                        <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200">
                          🥳
                        </div>
                        <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200">
                          😂
                        </div>
                        <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200">
                          🥰
                        </div>
                        <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200">
                          😡
                        </div>
                        <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200">
                          🤔
                        </div>
                        <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200">
                          😊
                        </div>
                        <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200">
                          🤩
                        </div>
                        <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200">
                          😎
                        </div>
                        <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200">
                          🥳
                        </div>
                        <div class="hover:bg-secondery p-1.5 rounded-md hover:scale-125 cursor-pointer duration-200">
                          😂
                        </div>
                      </div>
                    </div>
                  </div> -->
              </div>

              <div class="relative flex ">
                <textarea placeholder="Write your message" rows="1"
                  class="mr-2 w-full resize-none bg-secondery rounded-full px-10 p-2" />
                <button class="bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 w-1/4 right-0.5" @click="send">
                  Send
                </button>

                <!-- <button type="button" class="text-white shrink-0 p-2 absolute right-0.5 top-0">
                    <ion-icon class="text-xl flex" :icon="ioniconsSendOutline" />
                  </button> -->
              </div>

              <button type="button" class="flex h-full dark:text-white">
                <ion-icon class="text-3xl flex -mt-3" :icon="ioniconsHeartOutline" />
              </button>
            </div>
          </div>
          <div class="rightt w-full h-full absolute top-0 right-0 z-10 hidden transition-transform">
            <div
              class="w-[360px] border-l shadow-lg h-screen bg-white absolute right-0 top-0 uk-animation-slide-right-medium delay-200 z-50 dark:bg-dark2 dark:border-slate-700">
              <div class="w-full h-1.5 bg-gradient-to-r to-purple-500 via-red-500 from-pink-500 -mt-px" />

              <div class="py-10 text-center text-sm pt-20">
                <nuxt-img :src="'http://localhost:8081/' + user.avatarImage" class="w-24 h-24 rounded-full mx-auto mb-3"
                  alt=""></nuxt-img>
                <div class="mt-8">
                  <div class="md:text-xl text-base font-medium text-black dark:text-white">
                    {{ user.firstName }} {{ user.lastName }}
                  </div>
                  <div class="text-gray-500 text-sm mt-1 dark:text-white/80">

                  </div>
                </div>
                <div class="mt-5">
                  <nuxt-link to="/profile/{{ user.nickname }}"
                    class="inline-block rounded-full px-4 py-1.5 text-sm font-semibold bg-secondery">View profile {{ user.nickname }}</nuxt-link>
                </div>
              </div>

              <hr class="opacity-80 dark:border-slate-700">

              <ul class="text-base font-medium p-3">
                <li>
                  <div class="flex items-center gap-5 rounded-md p-3 w-full hover:bg-secondery">
                    <ion-icon :icon="ioniconsNotificationsOffOutline" class="text-2xl" /> Mute Notification
                    <label class="switch cursor-pointer ml-auto"> <input type="checkbox" checked><span
                        class="switch-button !relative" /></label>
                  </div>
                </li>
                <li>
                  <button type="button" class="flex items-center gap-5 rounded-md p-3 w-full hover:bg-secondery">
                    <ion-icon :icon="ioniconsFlagOutline" class="text-2xl" /> Report
                  </button>
                </li>
                <li>
                  <button type="button" class="flex items-center gap-5 rounded-md p-3 w-full hover:bg-secondery">
                    <ion-icon :icon="ioniconsSettingsOutline" class="text-2xl" /> Ignore messages
                  </button>
                </li>
                <li>
                  <button type="button" class="flex items-center gap-5 rounded-md p-3 w-full hover:bg-secondery">
                    <ion-icon :icon="ioniconsStopCircleOutline" class="text-2xl" /> Block
                  </button>
                </li>
                <li>
                  <button type="button"
                    class="flex items-center gap-5 rounded-md p-3 w-full hover:bg-red-50 text-red-500">
                    <ion-icon :icon="ioniconsTrashOutline" class="text-2xl" /> Delete Chat
                  </button>
                </li>
              </ul>

              <!-- close button -->
              <button type="button" class="absolute top-0 right-0 m-4 p-2 bg-secondery rounded-full"
                uk-toggle="target: .rightt ; cls: hidden">
                <ion-icon :icon="ioniconsClose" class="text-2xl flex" />
              </button>
            </div>

            <!-- overly -->
            <div class="bg-slate-100/40 backdrop-blur absolute w-full h-full dark:bg-slate-800/40"
              uk-toggle="target: .rightt ; cls: hidden" />
          </div>
        </div>


        <!-- user profile right info -->
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
const currentUser = useAuthUser();

const users = ref([])
let ws: WebSocket | undefined;

const message = ref<string>("");
const messages = useState<{ id: number, user: string, message: string, created_at: string }[]>(() => []);

useHead({
  title: "Message",
})

const userId = currentUser.value.id
const user = `${currentUser.value.firstName} ${currentUser.value.lastName}`

const log = (user: string, ...args: string[]) => {
  console.log("[ws]", user, ...args);
  messages.value.push({
    id: 0,
    message: args.join(" "),
    user: user,
    created_at: new Date().toLocaleString(),
  });
  scroll();
};

const connect = async () => {
  const isSecure = location.protocol === "https:";
  const url = (isSecure ? "wss://" : "ws://") + location.host + "/api/chat-ws?userId=" + userId.value;
  if (ws) {
    log("ws", "Closing previous connection before reconnecting...");
    ws.close();
    clear();
  }

  log("ws", "Connecting to", url, "...");
  ws = new WebSocket(url);

  ws.addEventListener("message", (event) => {
    const { user = "system", message = "" } = event.data.startsWith("{")
      ? JSON.parse(event.data)
      : { message: event.data };
    log(
      user,
      typeof message === "string" ? message : JSON.stringify(message),
    );
  });

  await new Promise((resolve) => ws!.addEventListener("open", resolve));
  log("ws", "Connected!");
};

const clear = () => {
  messages.value.splice(0, messages.value.length);
  log("system", "previous messages cleared");
};

const scroll = () => {
  nextTick(() => {
    console.log('scrooling')
    window.scrollTo(0, document.body.scrollHeight + 100);
  })
}

const send = () => {
  console.log("sending message...");
  console.log(message.value)
  if (message.value) {
    ws!.send(message.value);
  }
  message.value = "";
};

// const ping = () => {
//   log("ws", "Sending ping");
//   ws!.send("ping");
// };

onMounted(async () => {
  try {
    const response = await fetch('http://localhost:8081/users', {
      headers: {
        'Accept': 'application/json'
      }
    });

    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }


    const data = await response.json();

    users.value = data.list;
    const sortUsersAlphabetically = (users) => {
      users.value.sort((a, b) => {
        const usernameA = a.lastName.toLowerCase() + a.firstName.toLowerCase();;
        const usernameB = b.lastName.toLowerCase() + b.firstName.toLowerCase();
        return usernameA.localeCompare(usernameB);
      });
    };

    // Appelez la fonction pour trier les utilisateurs au moment approprié
    sortUsersAlphabetically(users);
  } catch (error) {
    console.error('Error fetching users:', error.message);
  }
  connect();
  scroll();

});
let receiverid
const chats = ref([])
const selectUser = async (receiver) => {
  receiverid = receiver.id
  try {
    const data = await getMessage(receiverid)
    chats.value = data.body
    chats.value.forEach(message => {
      const updatedAt = new Date(message.UpdatedAt);
      const day = updatedAt.toLocaleDateString(); // Date
      const hour = updatedAt.toLocaleTimeString(); // Heure

      // Ajout des valeurs à l'objet message
      message.Day = day;
      message.Hour = hour;
    });
    console.log("messageeeeeee: ",)
  } catch (error) {
    console.log(error)
  }
}


</script>