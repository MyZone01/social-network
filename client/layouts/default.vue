<template>
  <layout-header />
  <main w-full flex flex-col items-start justify-start overflow-hidden h-screen bg-gray-100 dark:bg-gray-950 py-6 px-8>
    <slot />
  </main>
  <layout-footer />
</template>
