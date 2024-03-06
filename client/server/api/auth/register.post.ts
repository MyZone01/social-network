import { processParts } from "~/server/utils/processParts";
import { Register } from "~/server/models/register";
import { ServerResponse, User } from "~/types";

export default defineEventHandler(async (event) => {
  const reader = await readMultipartFormData(event);
  if (!reader) return { status: 400, body: 'Bad request', ok: false }

  const { file, jsonData } = await processParts(reader);

  const register = new Register(jsonData);
  const [isValid, message] = register.validate();
  if (!isValid) {
    return sendError(event, createError({
      statusCode: 400,
      statusMessage: message
    }))
  }

  const response = await $fetch<ServerResponse<User>>("http://localhost:8081/registration", {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: register,
  });

  if (response.status !== "200") {
    return sendError(event, createError({
      statusCode: 400,
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
    expires: true
      ? new Date(Date.now() + config.cookieRememberMeExpires)
      : new Date(Date.now() + config.cookieExpires),
  });

  const { password: _password, ...userWithoutPassword } = userWithPassword;

  if (!file) {
    console.log("No File given")
    return {
      status: 200,
      body: 'No file uploaded',
      session: response.session,
      user: userWithoutPassword,
      ok: true
    };
  }

  const body = new FormData();
  body.append('file', new Blob([file.data]), file.filename);

  const response2 = await $fetch<{imageurl: string}>("http://localhost:8081/upload", {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
      Authorization: `Bearer ${response.session}`,
    },
    body
  });
  if (!response2.imageurl) {
    return { status: 200, body: response.message, session: response.session, ok: false };
  }

  register.avatarImage = response2.imageurl;

  const response3 = await $fetch<ServerResponse<{}>>("http://localhost:8081/updateuser", {
    method: "PUT",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
      Authorization: `Bearer ${response.session}`,
    },
    body: register
  });
  if (response3.status !== "200") {
    return { status: 200, body: response3.message, session: response.session, ok: false }
  }
  return {
    status: 200,
    body: "User registered successfully",
    session: response.session,
    user: userWithoutPassword,
    ok: true
  };
});