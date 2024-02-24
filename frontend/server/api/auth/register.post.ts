import { Register } from "@/server/models/register";

// export default defineEventHandler(async (event) => {
//   const body = await readBody(event);

//   const {
//     firstName,
//     lastName,
//     email,
//     nickname,
//     password,
//     repeatPassword,
//     dateOfBirth,
//     aboutMe,
//     form,
//     avatarUrl
//   } = body;

//   const requiredFields = [
//     firstName,
//     lastName,
//     email,
//     password,
//     repeatPassword,
//     dateOfBirth
//   ];
//   if (requiredFields.some((field) => field == "")) {
//     const errorMessage = "some Fields can not be empty";
//     return sendError(
//       event,
//       createError({ statusCode: 400, statusMessage: errorMessage })
//     );
//   }

//   if (password !== repeatPassword) {
//     return sendError(
//       event,
//       createError({
//         statusCode: 400,
//         statusMessage: "Passwords do not match",
//       })
//     );
//   }

//   const emailRegex = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$/;
//   if (!emailRegex.test(email)) {
//     return sendError(
//       event,
//       createError({
//         statusCode: 400,
//         statusMessage: "Email syntax not accepted",
//       })
//     );
//   }

//   let [year, month, day] = dateOfBirth.split("-");
//   let formattedDateOfBirth = new Date(year, month - 1, day);

//   const userData = {
//     email: email.trim(),
//     firstName: firstName.trim(),
//     lastName: lastName.trim(),
//     avatarImage: 'uploads/default.jpg',
//     nickname: nickname.trim(),
//     aboutMe: aboutMe.trim(),
//     password: password.trim(),
//     dateOfBirth: formattedDateOfBirth
//   };

//   const userSession = await $fetch("http://localhost:8081/registration", {
//     method: "POST",
//     body: JSON.stringify(userData),
//   });

//   if (!userSession) {
//     // LOGic handling Error from Server
//     return sendError(
//       event,
//       createError({
//         statusCode: 400,
//         statusMessage: `Rejection from server : ${userSession}`
//       })
//     );
//   } else {
//     return {
//       // server return the idSession will use to establish cookie
//       userSession,
//     };
//   }
// });

export default defineEventHandler(async (event) => {

  console.log("Registering a new user");
  // const body = await readBody(event);
  const reader = await readMultipartFormData(event);
  if (!reader) return { status: 400, body: 'Bad request' }
  let file;
  let jsonData;

  console.log(reader);
  

  for await (const part of reader) {
    if (part.name === "file") {
      // This is a file part, store it in file variable
      file = part;
    } else if (part.name === 'data') {
      const data = Buffer.from(part.data).toString();
      jsonData = JSON.parse(data);
    }
  }

  const register = new Register(jsonData);
  const [isValid, message] = register.validate();
  if (!isValid) {
    return { status: 400, body: message }
  }

  const response1: any = await $fetch("http://localhost:8081/registration", {
    method: "POST",
    body: JSON.stringify(register),
  })
  const reponseJson = await JSON.parse(response1);
  if (reponseJson.status !== "200") {
    return { status: 400, body: response1.message }
  }

  console.log("status3");

  // Create a new FormData instance
  if (!file) {
    return {
      status: 200,
      body: 'No file uploaded',
    };
  }

  console.log("file", file);


  // const body = new FormData();
  // // Append the file to the FormData instance
  // body.append('file', new Blob([file.data]), file.filename);

  // const response = await $fetch('http://localhost:8081/upload', {
  //   method: 'POST',
  //   headers: {
  //     Authorization: reponseJson.session,
  //   },
  //   body,
  // }).then(async (res: any) => await res.json()).catch((err) => {
  //   console.log(err);
  //   return {
  //     status: 500,
  //     body: 'Internal server error',
  //   };
  // });
  // console.log("response", response);
});