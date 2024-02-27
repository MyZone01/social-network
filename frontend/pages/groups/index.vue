<script setup>

const groups = ref([])

onMounted(async () => {
  const store = useGlobalAuthStore()
  const response = await $fetch('/api/group', {
    headers: {
      Authorization: `Bearer ${store.token}`,
    }
  })
  data.value = response.groups
  console.log(response)
})
</script>

<template>
  <NuxtLayout>
    <main id="site__main"
      class="2xl:ml-[--w-side]  xl:ml-[--w-side-sm] py-10 p-2.5 h-[calc(100vh-var(--m-top))] mt-[--m-top]">
      <div class="2xl:max-w-[1220px] max-w-[1065px] mx-auto">
        <div class="page-heading">
          <h1 class="page-title">
            Groups
          </h1>
          <nav class="nav__underline">
            <ul class="group"
              uk-switcher="connect: #group-tabs ; animation: uk-animation-slide-right-medium, uk-animation-slide-left-medium">
              <li> <a href="#"> Suggestions </a> </li>
              <li> <a href="#"> Popular </a> </li>
              <li> <a href="#"> My groups </a> </li>
            </ul>
          </nav>
        </div>
        <!-- group list tabs -->
        <div id="group-tabs" class="uk-switcher">
          <!-- card layout 1 -->
          <div v-for="group in groups" class="flex flex-row gap-5 overflow-hidden">
            <GroupCard :group="group" />
          </div>
        </div>
      </div>
    </main>
  </NuxtLayout>
</template>
