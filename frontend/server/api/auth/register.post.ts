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
    dateOfBirth,
    aboutMe,
    form,
    avatarUrl
  } = body;

  const requiredFields = [
    firstName,
    lastName,
    email,
    password,
    repeatPassword,
    dateOfBirth
  ];
  if (requiredFields.some((field) => field == "")) {
    const errorMessage = "some Fields can not be empty";
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

  let [year, month, day] = dateOfBirth.split("-");
  let formattedDateOfBirth = new Date(year, month - 1, day);
  
  const userData = {
    email: email.trim(),
    firstName: firstName.trim(),
    lastName: lastName.trim(),
    avatarImage: 'uploads/default.jpg',
    nickname: nickname.trim(),
    aboutMe: aboutMe.trim(),
    password: password.trim(),
    dateOfBirth: formattedDateOfBirth
  };

  const userSession = await $fetch("http://localhost:8081/registration", {
    method: "POST",
    body: JSON.stringify(userData),
  });

  if (!userSession) {
    // LOGic handling Error from Server
    return sendError(
      event,
      createError({
        statusCode: 400,
        statusMessage: `Rejection from server : ${userSession}`
      })
    );
  } else {
    return {
      // server return the idSession will use to establish cookie
      userSession,
    };
  }
});
