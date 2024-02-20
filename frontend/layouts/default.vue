<template>
  <Body class="bg-15-23-42 overflow-x-hidden">
    <Header />
    <Sidebar />
    <main
      id="site__main"
      class="2xl:ml-[--w-side]  xl:ml-[--w-side-sm] p-2.5 h-[calc(100vh-var(--m-top))] mt-[--m-top]"
    >
      <nuxt-child />
      <slot />
    </main>
    <Footer />
  </Body>
</template>

<script setup>
</script>