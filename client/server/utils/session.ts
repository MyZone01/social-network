import type { H3Event } from "h3";
import { User } from "~/types";

export async function getUserFromToken(event: H3Event) {
  const config = useRuntimeConfig();

  const cookie = getCookie(event, config.cookieName);
  if (!cookie) return {user: null, token: null};;

  const unsignedToken = unsign(cookie, config.cookieSecret);
  if (!unsignedToken) return {user: null, token: null};

  const token = deserialize(unsignedToken);

  try {
    const response = await $fetch<User>("http://localhost:8081/me", {
      method: "GET",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
        Authorization: `Bearer ${token.session}`,
      },
    });
    return { user: response, token: token.session };
  } catch (error) {
    deleteCookie(event, config.cookieName, {
      httpOnly: true,
      path: "/",
      sameSite: "strict",
      secure: process.env.NODE_ENV === "production",
    });
    return {user: null, token: null};;
  }
}