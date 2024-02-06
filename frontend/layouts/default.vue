<template>
  <Body class="dark:bg-gray-900 bg-white overflow-x-hidden">
    <Header />
    <main class="flex flex-col w-screen h-screen mt-[--m-top]">
      <slot />
    </main>
    <Footer />
  </Body>
</template>