import { Login } from "@/server/models/login";

export default defineEventHandler(async (event) => {
  console.log("Login endpoint called");
  const body = await readBody(event);

  const data = body.data;

  const login = new Login(data);
  const [isValid, message] = login.validate();
  console.log(login);

  if (!isValid) {
    return {
      status: 400,
      body: message,
    };
  }

  const response = await fetcher("http://localhost:8081/login", "POST", JSON.stringify(login), "");
  console.log(response);

  if (response.status !== "200") {
    return { status: 400, body: response.message, session: response.session, ok: false }
  }
  return {
    status: 200,
    body: "User registered successfully",
    session: response.session,
    ok: true
  };
})