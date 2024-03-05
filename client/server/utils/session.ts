import type { H3Event } from "h3";
import { User } from "~/types";

export async function getUserFromToken(event: H3Event) {
  const config = useRuntimeConfig();

  const cookie = getCookie(event, config.cookieName);
  if (!cookie) return null;

  const unsignedToken = unsign(cookie, config.cookieSecret);
  if (!unsignedToken) return null;

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
    return response;
  } catch (error) {
    deleteCookie(event, config.cookieName, {
      httpOnly: true,
      path: "/",
      sameSite: "strict",
      secure: process.env.NODE_ENV === "production",
    });
    return null;
  }
}