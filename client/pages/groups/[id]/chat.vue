<template>
  <h1>Group Chat: </h1>
  {{ $route.params.id }}
  {{ group }}
  <nuxt-link to="/">Back to home</nuxt-link>
</template>

<script lang="ts" setup>
import type { Group } from '~/types';

definePageMeta({
  alias: ["/groups/[id]/chat"],
  middleware: ["auth-only"],
});

const { getGroupByID } = useGroups();
const route = useRoute();
const id = route.params.id as string;
const group = ref<Group | null>(null);

onMounted(async () => {
  group.value = await getGroupByID(id);
});
</script>

<style></style>