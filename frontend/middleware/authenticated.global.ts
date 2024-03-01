import { decoder } from "@/server/utils/transformer";
import { useGlobalAuthStore } from "@/stores/useGlobalStateAuthStore";
import { ConsoleMessage } from "puppeteer";
let checker: Boolean
interface ServerResponse {
  error: String,
  message: String,
}

export default defineNuxtRouteMiddleware(async (to, from) => {
  if (process.server) { return }

  const authStore = useGlobalAuthStore();
  // const pass = authStore.token
  const token = authStore.token// decoder(pass.toString())
  
  const response = await $fetch('/api/auth/session', {
    method: 'POST',
    body: {
      token
    }
  })
  .then((res) => {
    checker = res ? true : false
    const isAuthenticated = authStore.isAuthenticated && checker;  
    if (!isAuthenticated && to.path !== "/auth") {
      authStore.logout();
      return navigateTo("/auth");
    }
    if (isAuthenticated && to.path === '/auth') {
      return navigateTo(from.path)
    }
  })
  .catch((error) => {
    return
  } )
});
