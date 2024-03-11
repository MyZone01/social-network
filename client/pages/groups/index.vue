<template>
    <NuxtLayout>
        <main
          id="site__main"
          class="2xl:ml-[--w-side] xl:ml-[--w-side-sm] py-10 p-2.5 h-[calc(100vh-var(--m-top))] mt-[--m-top]"
        >
          <div class="2xl:max-w-[1220px] max-w-[1065px] mx-auto h-full">
            <div
              class="page-heading flex flex-row justify-between align-baseline flex-wrap"
            >
              <h1 class="page-title">Groups</h1>
              <div
                uk-toggle="target: #create-group-overlay"
                class="h-fit self-center"
              >
                <UButton class="bg-blue-500 h-fit">
                  <UIcon name="i-heroicons-plus" class="" /><span>New</span>
                </UButton>
              </div>
            </div>
            <!-- group list tabs -->
            <div id="group-tabs">
              <nav>
                <ul
                  class="uk-subnav uk-subnav-pill flex gap-0.5 rounded-xl overflow-hidden -mb-px text-gray-500 font-medium text-sm overflow-x-auto dark:text-white"
                  uk-switcher="connect: #group-tab ; animation: uk-animation-slide-right-medium, uk-animation-slide-left-medium"
                >
                  <li>
                    <a href="#" class="inline-block py-3 leading-8 px-3.5"
                      >My Groups</a
                    >
                  </li>
                  <li>
                    <a href="#" class="inline-block py-3 leading-8 px-3.5"
                      >Browse</a
                    >
                  </li>
                </ul>
              </nav>
            </div>
            <!-- card layout 1 -->
            <div id="group-tab" class="uk-switcher">
              <div>
                <div
                  v-if="groups.length"
                  class="flex flex-row flex-wrap overflow-scroll gap-2"
                >
                  <GroupCard
                    v-bind:key="group.ID"
                    v-for="group in groups"
                    :group="group"
                    :joined="true"
                  />
                </div>
              </div>
              <div
                v-if="groups.length"
                class="flex flex-row flex-wrap overflow-scroll"
              >
                <GroupCard
                  v-bind:key="group.ID"
                  v-for="group in groups"
                  :group="group"
                />
              </div>
              <div v-else class="flex w-full flex-row justify-center align-middle">
                <div class="flex flex-col justify-center">
                  <p class="text-base text-lg">there is no group at the moment</p>
                  <div
                    uk-toggle="target: #create-group-overlay"
                    class="h-fit self-center"
                  >
                    <UButton class="bg-blue-500 h-fit">
                      <UIcon name="i-heroicons-plus" class="" /><span
                        >Create new</span
                      >
                    </UButton>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <GroupCreateModal />
        </main>
    </NuxtLayout>
</template>
<style scoped>
li.uk-active {
  border-bottom: 2px solid rgb(37 99 235 / var(--tw-text-opacity));
  --tw-text-opacity: 1;
  color: rgb(37 99 235 / var(--tw-text-opacity));
}
</style>
<script setup lang="ts">
const { groups, getAllGroups } = useGroups();
onMounted(async () => {
  console.log(groups);
  
  await getAllGroups();
});
</script>