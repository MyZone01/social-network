<template>
  <main id="site__main"
    class="2xl:ml-[--w-side] xl:ml-[--w-side-sm] p-2.5 h-[calc(100vh-var(--m-top))] mt-[--m-top] overflow-y-auto">
    <div class="h-screen flex flex-col justify-between">
      <div class="page-heading">
        <h1>🗨️Group Chat: </h1>
        <div class="flex flex-col" v-if="group">
          {{ group.Title }}
          {{ group.Description }}
          <div v-if="messagesList">
            <div v-for="message in messagesList" :key="message.ID">
              <div>
                <div v-if="message.SenderID === currentUser?.id" class="flex gap-2 flex-row-reverse items-end">
                  <img v-if="currentUser && currentUser.avatarImage"
                    :src="'http://localhost:8081/' + currentUser.avatarImage" class="w-9 h-9 rounded-full shadow" />
                  <div
                    class="px-4 py-2 rounded-[20px] max-w-sm bg-gradient-to-tr from-sky-500 to-blue-500 text-white shadow">
                    {{ message.Content }}
                  </div>
                </div>
                <div v-else class="flex gap-3">
                      <img 
                       :src="'http://localhost:8081/' + message.Sender.avatarImage"
                        class="w-9 h-9 rounded-full shadow" />
                      <div class="px-4 py-2 rounded-[20px] max-w-sm bg-secondery">
                        {{ message.Content }}
                      </div>
                    </div>
                <div class="bg-red-400">{{ message.Sender.nickname }}</div>
                <div>{{ message.CreatedAt }}</div>
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
</template>

<script lang="ts" setup>
const currentUser = useAuthUser();
const userId = currentUser.value?.id
import type { Group, GroupMessage } from '~/types';

const { getGroupByID, getAllMessagesByGroup } = useGroups();
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