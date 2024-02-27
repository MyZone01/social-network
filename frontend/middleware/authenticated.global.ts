import { decoder } from "@/server/utils/transformer";
import { useGlobalAuthStore } from "@/stores/useGobalStateAuthStore";
let checker: Boolean

export default defineNuxtRouteMiddleware(async (to, from) => {
  if (process.server) { return }

  const authStore = useGlobalAuthStore();
  // const pass = authStore.token
  const token = authStore.token// decoder(pass.toString())
  
  const response = await useFetch('/api/auth/session', {
    method: 'POST',
    body: {
      token
    }
  })
  .then((res) => {
    checker = res ? true : false
  })

  const isAuthenticated = authStore.isAuthenticated && checker;  

  if (!isAuthenticated && to.path !== "/auth") {
    authStore.logout();
    return navigateTo("/auth");
  }
});
