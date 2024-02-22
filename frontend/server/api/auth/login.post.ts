import { sendError } from "h3";
import axios from "axios";
import { useGlobalAuthStore } from "../../../stores/useGobalStateAuthStore";

interface response {
  idSession: String
  userInfos: Object
}

export default defineEventHandler(async (event) => {
  const body = await readBody(event);
  const store = useGlobalAuthStore()

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

  try {
    const result: response = await $fetch("http://localhost:8081/login", {
      method: "POST",
      body: JSON.stringify(loginAccess),
    })

    console.log(result)
    let userInfos;
    console.log("USER INFOS VARIABLE => ", result);
    
  } catch(err) {
    return sendError(event, createError({
      statusCode: 400,
        statusMessage: `Not Valid: ${err}`,
      })
    );
  }
})
// await axios
// .get("http://localhost:8081/userinfos", {
  //     headers: {
//       Authorization: `Bearer ${res.idSession}`,
//     },
//   })
//   .then((res) => {
//     console.log("USER INFOS VARIABLE => ", res.data);
//     console.log("USER INFOS VARIABLE => ", userInfos);
//     userInfos = !res.error ? res : {};
//   })
//   .catch((err) => {
  //     throw err;
  //   });
//     return {
//       // server return the idSession will use to establish cookie
//       result
//     };
//   })
//   .catch((error) => {
//   })

//   if (!response) {
//     // LOGic handling Error from Server
//   } else {
//   }
// });
