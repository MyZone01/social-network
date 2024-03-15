<template>
  <main id="site__main"
    class="2xl:ml-[--w-side]  xl:ml-[--w-side-sm] py-10 p-2.5 h-[calc(100vh-var(--m-top))] mt-[--m-top]">
    <div class="h-screen flex flex-col justify-between">
      <div class="page-heading">
        <h1>🗨️Group Chat: </h1>
        <div class="flex flex-col" v-if="group">
          {{ groupId }}
          {{ group.Title }}
          {{ group.Description }}
          <div v-if="messagesList">
            <div v-for="message in messagesList" :key="message.ID">
              <div>
                <p class="bg-red-400">{{ message.Sender.nickname }}</p>
                <p>{{ message.Content }}</p>
                <p>{{ message.CreatedAt }}</p>
              </div>
            </div>
          </div>
        </div>
        <u-input v-model="message" placeholder="Type your message..." @keydown.enter="send" />
        <u-button @click="send">Send</u-button>
        <nuxt-link :to="`/groups/${groupId}`">Back Group</nuxt-link>
      </div>
    </div>
  </main>
  <!-- <main>
      <div class="flex flex-col text-white" style="position: fixed; right: 0;">
        <a href="https://github.com/pi0/nuxt-websocket">
          view source on GitHub
        </a>
      </div>

      <div id="messages" class="flex-grow flex flex-col justify-end px-4 pt-8 pb-21 sm:pb-12 bg-slate-900 min-h-screen">
        <div class="flex items-center mb-4 overflow-x-scroll" v-for="message in messages" :key="message.id">
          <div class="flex flex-col">
            <p class="text-gray-500 mb-1 text-xs ml-10">{{ message.user }}</p>
            <div class="flex items-center">
              <img :src="'https://www.gravatar.com/avatar/' + encodeURIComponent(message.user) + '?s=512&d=monsterid'"
                alt="Avatar" class="w-8 h-8 rounded-full" />
              <div class="ml-2 bg-gray-800 rounded-lg p-2">
                <p class="text-white">{{ message.message }}</p>
              </div>
            </div>
            <p class="text-gray-500 mt-1 text-xs ml-10">{{ message.created_at }}</p>
          </div>
        </div>
      </div>

      <div class="bg-gray-800 px-4 py-2 flex items-center justify-between fixed bottom-0 w-full flex-col sm:flex-row">
        <div class="w-full min-w-6">
          <input type="text" placeholder="Type your message..."
            class="w-full rounded-none px-4 py-2 bg-gray-700 text-white focus:outline-none focus:ring focus:border-blue-300 sm:rounded-l-lg"
            @keydown.enter="send" v-model="message" />
        </div>
        <div class="flex w-full">
          <button class="bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 w-1/4" @click="send">
            Send
          </button>
          <button class="bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 w-1/4" @click="ping">
            Ping
          </button>
          <button class="bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 w-1/4" @click="connect">
            Reconnect
          </button>
          <button class="bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 sm:rounded-r-lg w-1/4" @click="clear">
            Clear
          </button>
        </div>
      </div>
    </main> -->
</template>

<script lang="ts" setup>
import type { Group, GroupMessage } from '~/types';

<<<<<<< HEAD
const { getGroupByID, getAllMessagesByGroup } = useGroups();
=======
definePageMeta({
  alias: ["/groups/[id]/chat"],
  middleware: ["auth-only"],
});

const { getGroupByID } = useGroups();
>>>>>>> group-timeline
const route = useRoute();
const groupId = route.params.id as string;
const group = ref<Group | null>(null);
const message = ref<string>("");
const messagesList = ref<GroupMessage[]>([]);
let ws: WebSocket | undefined;

useHead({
  title: "Chat " + group.value ? group.value?.Title : '',
})

definePageMeta({
  alias: ["/groups/:id/chat"],
  middleware: ["auth-only"],
});

const scroll = () => {
  nextTick(() => {
    console.log('scrooling')
    window.scrollTo(0, document.body.scrollHeight + 100);
  })
}

const log = (user: string, message: GroupMessage | string) => {
  if (typeof message === "string") {
    console.log("[ws]", user, message);
  } else {
    console.log("[ws]", user, message.Content);
    if (!message.CreatedAt) {
      message.CreatedAt = new Date().toISOString();
    }
    messagesList.value = [...messagesList.value, message];
    scroll();
  }
};

const send = () => {
  console.log("sending message...");
  if (message.value) {
    ws!.send(message.value);
  }
  message.value = "";
};

const connect = async () => {
  const url = "ws://" + location.host + "/api/groups/chat-ws?group_id=" + groupId;
  if (ws) {
    log("ws", "Closing previous connection before reconnecting...");
    ws.close();
  }

  log("ws", "Connecting to" + url + "...");
  ws = new WebSocket(url);

  ws.addEventListener("message", (event) => {
    const message = (event.data.startsWith("{")
      ? JSON.parse(event.data)
      : event.data) as GroupMessage;
    log(
      message.Sender.nickname,
      message,
    );
  });

  await new Promise((resolve) => ws!.addEventListener("open", resolve));
  log("ws", "Connected!");
};

onMounted(async () => {
  group.value = await getGroupByID(groupId)
  await connect();
  const _messages = await getAllMessagesByGroup(groupId)
  if (_messages) {
    messagesList.value = [...messagesList.value, ..._messages]
  }
  scroll();
});
</script>