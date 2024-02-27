import { sendError } from "h3";

export default defineEventHandler(async (event) => {
  const body = await readBody(event);

  const { email, password } = body;

  const requiredFields = [email, password];
  if (requiredFields.some((field) => field == "")) {
    return sendError(
      event,
      createError({
        statusCode: 400,
        statusMessage: "Fields cannot be empty",
      })
    );
  }

  // Check if User mail or nickname syntax is correct
  const emailRegex = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$/;
  const userEmail = emailRegex.test(email);
  if (!userEmail || (email.includes(" ") && email.length >= 3)) {
    return sendError(
      event,
      createError({
        statusCode: 400,
        statusMessage: "Email or password is invalid",
      })
    );
  }

  const loginAccess = {
    email: email,
    password: password,
  };

  const userSession = await $fetch("http://localhost:8081/login", {
    method: "POST",
    body: JSON.stringify(loginAccess),
  });

  if (userSession.error) {
    // LOGic handling Error from Server
    return sendError(
      event,
      createError({
        statusCode: 400,
        statusMessage: `Not Valid: ${userSession.error}`
      })
    );
  } else {

    return {
      // server return the idSession will use to establish cookie
      userSession,
    };
  }
});
