<template>
  <h1>Group: </h1>
  {{ id }}
  {{ group }}
  <nuxt-link :to="`/groups/${id}/chat`">Go to chat</nuxt-link>
  <nuxt-link to="/">Back to home</nuxt-link>
  <nuxt-link to="/groups">Back to group</nuxt-link>
</template>

<script lang="ts" setup>
import type { Group } from '~/types';

const { getGroupByID } = useGroups();
const route = useRoute();
const id = route.params.id as string;
const group = ref<Group | null>(null);

onMounted(async () => {
  group.value = await getGroupByID(id);
});
</script>

<style></style>