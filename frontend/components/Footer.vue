<template>
  <footer class="fixed bottom-0 w-screen flex items-center justify-center text-sm shadow-inner">
    Copyright ©️ {{ new Date().getFullYear() }} | Made with <span mx-1 class="heart">❤️</span> and ☕ by @01Talents
  </footer>
</template>
<style scoped lang="css">
footer {
  @apply h-[25px]
}

@-webkit-keyframes beat {
  0% {
    transform: scale(1);
  }

  50% {
    transform: scale(1.35);
  }
}

@keyframes beat {
  0% {
    transform: scale(1);
  }

  50% {
    transform: scale(1.35);
  }
}

span.heart {
  transform-origin: center;
  animation: beat 1.5s infinite alternate;
}
</style>