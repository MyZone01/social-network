import { ServerResponse, User } from "~/types";
import { serialize, sign } from "~/server/utils/cookie";

export default defineEventHandler(async (event) => {
  const body = await readBody<{ username: string; password: string; rememberMe: boolean }>(event);

  const { username, password, rememberMe } = body;

  const response = await $fetch<ServerResponse<User>>("http://localhost:8081/login", {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: {
      email: username,
      password,
    },
  });

  if (response.status !== "200") {
    return sendError(event, createError({
      statusCode: 400 as number,
      statusMessage: response.message
    }))
  }

  const config = useRuntimeConfig();
  const userWithPassword = response.user;
  const session = serialize({ session: response.session });
  const signedSession = sign(session, config.cookieSecret);

  setCookie(event, config.cookieName, signedSession, {
    httpOnly: true,
    path: "/",
    sameSite: "strict",
    secure: process.env.NODE_ENV === "production",
    expires: rememberMe
      ? new Date(Date.now() + config.cookieRememberMeExpires)
      : new Date(Date.now() + config.cookieExpires),
  });

  const { password: _password, ...userWithoutPassword } = userWithPassword;

  return {
    user: userWithoutPassword,
  };
});