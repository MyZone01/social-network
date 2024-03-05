import { ServerResponse } from "~/types";

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();

  const cookie = getCookie(event, config.cookieName);
  if (!cookie) return null;

  const unsignedToken = unsign(cookie, config.cookieSecret);
  if (!unsignedToken) return null;

  const token = deserialize(unsignedToken);

  const response = await $fetch<ServerResponse<{}>>("http://localhost:8081/logout", {
    method: "DELETE",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
      Authorization: `Bearer ${token.session}`,
    },
  });

  deleteCookie(event, config.cookieName, {
    httpOnly: true,
    path: "/",
    sameSite: "strict",
    secure: process.env.NODE_ENV === "production",
  });

  return {
    user: null,
  };
});