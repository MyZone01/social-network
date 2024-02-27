import { Login } from "@/server/models/login";
import { encoder, secure } from "@/server/utils/transformer";

export default defineEventHandler(async (event) => {
  const body = await readBody(event);

  const data = body.data;

  const login = new Login(data);
  const [isValid, message] = login.validate();
  if (!isValid) {
    return {
      status: 400,
      body: message,
    };
  }
  
  const response = await fetcher("http://localhost:8081/login", "POST", JSON.stringify(login), "");

  if (response.status !== "200") {
    return { status: 400, body: response.message, session: response.session, ok: false }
  }
  return {
    status: 200,
    body: "User registered successfully",
    session: encoder(response.session),
    user: secure(response.user),
    ok: true
  };
})