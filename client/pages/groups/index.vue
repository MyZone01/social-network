<template>
  <h1>Groups: </h1>
  <ul>
    <li v-for="group in groups" :key="group.ID">
      <nuxt-link :to="`/groups/${group.ID}`">{{ group.Title }}</nuxt-link>
    </li>
  </ul>
  <nuxt-link to="/">Back to home</nuxt-link>
</template>

<script setup lang="ts">
import type { Group } from '~/types';

const { getAllGroups } = useGroups();
const groups = ref<Group[]>([]);

onMounted(async () => {
  groups.value = await getAllGroups();
});
</script>