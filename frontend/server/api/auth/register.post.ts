import { Register } from "@/server/models/register";
import { fetcher } from "@/server/utils/fetcher";
import { processParts } from "@/server/utils/processParts";
import { secure } from "@/server/utils/transformer";
import { sessionCreator } from "@/server/utils/createSession";

export default defineEventHandler(async (event) => {
  const reader = await readMultipartFormData(event)
  if (!reader) return { status: 400, body: 'Bad request', ok: false }

  const { file, jsonData } = await processParts(reader)

  const register = new Register(jsonData)
  const [isValid, message] = register.validate()
  if (!isValid) {
    return sendError(event, createError({
      statusCode: 400,
      statusMessage: message
    }))
  }

  const response = await fetcher('http://localhost:8081/registration', 'POST', JSON.stringify(register), '')
  if (response.status !== '200') {
    return sendError(event, createError({
      statusCode: 400,
      statusMessage: response.message
    }))
  }

  
  if (!file) {
    await sessionCreator(response.session, response.user, event)
    return {
      status: 200,
      body: 'No file uploaded',
      session: response.session,
      user: secure(response.user),
      ok: true
    }
  }

  const body = new FormData()
  body.append('file', new Blob([file.data]), file.filename)

  const response2 = await fetcher('http://localhost:8081/upload', 'POST', body, response.session)
  if (!response2.imageurl) {
    return { status: 200, body: response.message, session: response.session, ok: false }
  }

  register.avatarImage = response2.imageurl;
  const response3 = await fetcher("http://localhost:8081/updateuser", "PUT", JSON.stringify(register), response.session)
  if (response3.status !== 200) {
    return { status: 200, body: response3.message, session: response.session, ok: false }
  }

  response.user.avatarImage = register.avatarImage
  await sessionCreator(response.session, response.user, event)
  // const serverSession = await useSession(event, {
  //   password: "5ec0312f-223f-4cc0-aa0f-303ff39fe1b2",
  //   name: "server-store",
  //   cookie: {
  //     httpOnly: true,
  //     secure: true,
  //     sameSite: "strict",
  //   },
  //   maxAge: 60 * 60 * 24 * 7,
  //   generateId: () => { return response.session }
  // })
  // await serverSession.update({
  //   userInfos: response.user
  // })
  return {
    status: 200,
    body: 'User registered successfully',
    session: response.session,
    user: secure(response.user),
    ok: true
  }
})