import { decoder } from "@/server/utils/transformer";
import axios from "axios";


export default defineNuxtRouteMiddleware(async (to, from) => {
  const { token, isAuthenticated, user } = useGlobalAuthStore();
  let tokenValid: Boolean;
  let decodeToken: String
  let cookie = useCookie('token')
  console.log(cookie)

  // if (cookie.value !== token && to.path !== '/auth') {
  //   return navigateTo('/auth')
  // }
  if (typeof cookie == 'string') {
    // decodeToken = decoder(cookie.toString())
  }



  // tokenValid = false
  // await axios
  //   .get("http://localhost:8081/checksession", {
  //     headers: {
  //       Authorization: `Bearer ${token}`,
  //     },
  //   })
  //   .then(({ data }) => {
  //     tokenValid = !data.error ? true : false;
  //   })
  //   .catch((err) => {
  //     tokenValid = false
  //     return;
  //   });

  // const authenticated = isAuthenticated && tokenValid;

  // if (authenticated === false && to.path !== '/auth') {
  //   return navigateTo('/auth')
  // }
})