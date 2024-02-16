<template>
  <Body class="bg-15-23-42 overflow-x-hidden">
    <Header />
    <Sidebar />
    <main id="site__main" class="ml-48 2xl:ml-0 py-5 px-6 h-[calc(100vh-var(--m-top))] mt-[--m-top]">
      <slot />
    </main>
    <Footer />
  </Body>
</template>