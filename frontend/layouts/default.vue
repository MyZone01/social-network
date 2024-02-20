<template>
  <Body class="bg-15-23-42 overflow-x-hidden">
    <Header />
    <Sidebar />
    <section style="max-height: 100%; overflow-y: auto;">
      <slot />
    </section>
    <Footer />
  </Body>
</template>

<script setup>
</script>