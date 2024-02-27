import axios from "axios";

export default defineNuxtRouteMiddleware(async (to, from) => {
  const authStore = useGlobalAuthStore();
  let tokenValid;
  const authenticate = authStore.isAuthenticated;
  const token = authStore.token 


  try {
    await $fetch("http://localhost:8081/checksession", {
      headers: {
        Authorization: `Bearer ${authStore.token}`,
      },
    }).then(response => {
      console.log("\n\n>>>>>>>>>>> response <<<<<<<<<<<<<<< ", response);

      if (response) {
        if (to.path !== "/auth") {
          return navigateTo("/auth");
        } else {
          return navigateTo("/");
        }
      }
    }).catch(err => {
      throw err
    })
  } catch (err) {
    if (to.path !== '/auth') {

      return navigateTo("/auth");
    }
  };

})
