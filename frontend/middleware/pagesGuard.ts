import { decoder } from "@/server/utils/transformer";
import axios from "axios";


export default defineNuxtRouteMiddleware(async (to, from) => {
    const {token, isAuthenticated, user} = useGlobalAuthStore();
    let tokenValid: Boolean;
    let decodeToken = decoder(String(token))
    
    tokenValid = false
    await axios
      .get("http://localhost:8081/checksession", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then(({ data }) => {
        tokenValid = !data.error ? true : false;
      })
      .catch((err) => {
        tokenValid = false
        return;
    });

    const authenticated = isAuthenticated && tokenValid;
    if (authenticated === false) {
      return navigateTo('/auth')
    }
  })