<template>
  <Body class="dark:bg-gray-900 bg-white overflow-x-hidden">
    <Header />
    <Sidebar />
    <main class="ml-48 2xl:ml-0 py-5 px-6 h-[calc(100vh-var(--m-top))] mt-[--m-top]">
      <slot />
    </main>
    <Footer />
  </Body>
</template>