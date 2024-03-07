import { User } from "~/types";

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();

  const cookie = getCookie(event, config.cookieName);

  if (!cookie) return null;

  const unsignedToken = unsign(cookie, config.cookieSecret);

  if (!unsignedToken) return null;

  const token = deserialize(unsignedToken);

  const response = await $fetch<User>("http://localhost:8081/get-all-groups?isMemberNeeded=true&isUserNeeded=true", {
    method: "GET",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
      Authorization: `Bearer ${token.session}`,
    },
  });

  return response;
});