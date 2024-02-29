import { Login } from "@/server/models/login";
import { sendError, useSession, setCookie } from 'h3'
import { secure } from "@/server/utils/transformer";

export default defineEventHandler(async (event) => {
  const body = await readBody(event);

  const data = body.data;

  const login = new Login(data);
  const [isValid, message] = login.validate();
  if (!isValid) {
    return sendError(event, createError({
      statusCode: 400,
      statusMessage: message
    }))
  }
  const response = await fetcher("http://localhost:8081/login", "POST", JSON.stringify(login), "");
  
  if (response.status !== "200") {
    return sendError(event, createError({
      statusCode: 400,
      statusMessage: message
    }))
  }
  
  const serverSession = await useSession(event, {
    password: "5ec0312f-223f-4cc0-aa0f-303ff39fe1b2",
    name: "server-store",
    cookie: {
      httpOnly: true,
      secure: true,
      sameSite: "strict",
    },
    maxAge: 60 * 60 * 24 * 7,
    generateId: () => { return response.session }
  })
  await serverSession.update({
    userInfos: response.user
  })

  return {
    status: 200,
    body: "User registered successfully",
    session: response.session,
    user: secure(response.user),
    ok: true
  };
})

