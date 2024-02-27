import { Login } from "@/server/models/login";
import { secure } from "@/server/utils/transformer";
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

  if (response.status !== "200") {
    return { status: 400, body: response.message, session: response.session, ok: false }
  }
  return {
    status: 200,
    body: "User registered successfully",
    session: response.session,
    user: secure(response.user),
    ok: true
  };
})