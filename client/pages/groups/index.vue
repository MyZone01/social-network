<template>
  <h1>Groups: </h1>
  <input-group />
  <ul>
    <li v-for="group in groups" :key="group.ID">
      <nuxt-link :to="`/groups/${group.ID}`">{{ group.Title }}</nuxt-link>
    </li>
  </ul>
  <nuxt-link to="/">Back to home</nuxt-link>
</template>

<script setup lang="ts">
const { groups, getAllGroups } = useGroups();

onMounted(async () => {
  await getAllGroups();
});
</script>