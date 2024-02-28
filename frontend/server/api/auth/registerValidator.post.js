import { sendError } from "h3";

export default defineEventHandler(async (event) => {
  const body = await readBody(event);

  const {
    firstName,
    lastName,
    email,
    nickname,
    password,
    repeatPassword,
    aboutMe,
    avatarImg,
  } = body;

  const requiredFields = [
    firstName,
    lastName,
    email,
    nickname,
    password,
    repeatPassword,
  ];
  
  if (requiredFields.some((field) => field == "")) {
    const errorMessage = "Invalid entries";
    return sendError(
      event,
      createError({ statusCode: 400, statusMessage: errorMessage })
    );
  }

  if (password !== repeatPassword) {
    return sendError(
      event,
      createError({
        statusCode: 400,
        statusMessage: "Passwords do not match",
      })
    );
  }

  const emailRegex = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$/;
  if (!emailRegex.test(email)) {
    return sendError(
      event,
      createError({
        statusCode: 400,
        statusMessage: "Email syntax not accepted",
      })
    );
  }

  const userData = {
    email: email.trim(),
    firstName: firstName.trim(),
    lastName: lastName.trim(),
    avatarImage: avatarImg,
    nickname: nickname.trim(),
    aboutMe: aboutMe.trim(),
    password: password.trim(),
  };
  
    const userSession = await $fetch("http://localhost:8081/registrations", {
      method: "POST",
      body: JSON.stringify(userData),
    });

    if (userSession.error) {
      // LOGic handling Error from Server
      return sendError(
        event,
        createError({
          statusCode: 400,
          statusMessage: `Rejection from server : ${userSession.error}`
        })
      );
    } else {
      return {
        // server return the idSession will use to establish cookie
        userSession,
      };
    }
});
