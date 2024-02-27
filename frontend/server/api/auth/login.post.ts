import { Login } from "@/server/models/login";
import { sendError } from 'h3'

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
  console.log(response)

  if (response.status !== "200") {
    return sendError(event, createError({
      statusCode: 400,
      statusMessage: message
    }))
  }
  return {
    status: 200,
    body: "User registered successfully",
    session: response.session,
    user: response.user,
    ok: true
  };
})