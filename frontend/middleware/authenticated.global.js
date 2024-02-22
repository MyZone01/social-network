import axios from "axios";

export default defineNuxtRouteMiddleware(async (to, from) => {
  const authStore = useGlobalAuthStore();
  let tokenValid;
  const authenticate = authStore.isAuthenticated;
  
  await axios
    .get("http://localhost:8081/checksession", {
      headers: {
        Authorization: `Bearer ${authStore.token}`,
      },
    })
    .then((res) => {
      tokenValid = !res.error ? true : false;
    })
    .catch((err) => {
      return;
    });

  const isAuthenticated = authenticate && tokenValid;

  // console.log(authStore.user)

  if (!isAuthenticated && to.path !== "/auth") {
    authStore.logout();
    return navigateTo("/auth");
  }

  if (isAuthenticated && to.path === "/auth") {
    return navigateTo("/");
  }
});