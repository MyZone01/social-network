<template>
  <main id="site__main"
    class="2xl:ml-[--w-side]  xl:ml-[--w-side-sm] py-10 p-2.5 h-[calc(100vh-var(--m-top))] mt-[--m-top]">
    <div class="h-screen flex flex-col justify-between">
      <div class="page-heading">
        <div class="flex">
          <UButton to="/"> <- </UButton>
          <h1>Group: </h1>
        </div>
        {{ id }}
        {{ group }}
        <div class="flex flex-col">
          <UButton :to="`/groups/${id}/chat`">Go to chat</UButton>
          <UButton to="/groups">Back to group</UButton>
        </div>
      </div>
    </div>
  </main>
</template>

<script lang="ts" setup>
import type { Group } from '~/types';

const { getGroupByID } = useGroups();
const route = useRoute();
const id = route.params.id as string;
const group = ref<Group | null>(null);

onMounted(async () => {
  group.value = await getGroupByID(id) || null;
});

useHead({
  title: "Group",
})

definePageMeta({
  alias: ["/groups/:id"],
  middleware: ["auth-only"],
});
</script>

<style></style>