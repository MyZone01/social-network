<template>
  <Body class="dark:bg-gray-900 bg-white overflow-hidden">
    <Header />
    <Sidebar />
    <main class="ml-48 flex flex-col w-screen h-screen overflow-y-auto mt-[--m-top] p-4">
      <slot />
    </main>
    <Footer />
  </Body>
</template>