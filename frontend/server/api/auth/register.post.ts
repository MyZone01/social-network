import { Register } from "@/server/models/register";
import { fetcher } from "@/server/utils/fetcher";
import { processParts } from "@/server/utils/processParts";
import { encoder, secure } from "@/server/utils/transformer";


export default defineEventHandler(async (event) => {
  const reader = await readMultipartFormData(event);
  if (!reader) return { status: 400, body: 'Bad request', ok: false }

  const { file, jsonData } = await processParts(reader);

  const register = new Register(jsonData);
  const [isValid, message] = register.validate();
  if (!isValid) {
    return { status: 400, body: message, ok: false }
  }

  const response = await fetcher("http://localhost:8081/registration", "POST", JSON.stringify(register), "")
  if (response.status !== "200") {
    return { status: 400, body: response.message, ok: false }
  }

  if (!file) {
    return {
      status: 200,
      body: 'No file uploaded',
      session: encoder(response.session),
      user: secure(response.user),
      ok: true
    };
  }

  const body = new FormData();
  body.append('file', new Blob([file.data]), file.filename);

  const response2 = await fetcher("http://localhost:8081/upload", "POST", body, response.session)
  if (!response2.imageurl) {
    return { status: 400, body: response.message, session: response.session, ok: true }
  }

  register.avatarImage = response2.imageurl;

  const response3 = await fetcher("http://localhost:8081/updateuser", "PUT", JSON.stringify(register), response.session)
  if (response3.status !== "200") {
    return { status: 400, body: response3.message, session: response.session, ok: true }
  }
  return {
    status: 200,
    body: "User registered successfully",
    session: response.session,
    user: response.user,
    ok: true
  };
});