import { useGlobalAuthStore } from "@/stores/useGlobalStateAuthStore";
let checker: Boolean

export default defineNuxtRouteMiddleware(async (to, from) => {
  if (process.server) { return }

  const authStore = useGlobalAuthStore();
  const token = authStore.token
  const response = await $fetch('/api/auth/session', {
    method: 'POST',
    body: {
      token
    }
  })
  .then((res) => {
    checker = res ? true : false
    const isAuthenticated = authStore.isAuthenticated && checker;
    if (!isAuthenticated) {
      authStore.logout();
      return navigateTo("/auth");
    }
    // if (isAuthenticated && to.path === '/auth') {
    //   return navigateTo('/auth')
    // }
  })
  .catch((error) => {
    authStore.logout();
    return
  } )
});
