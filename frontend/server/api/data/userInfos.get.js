import axios from "axios";
import { useGlobalAuthStore } from "../../../stores/useGlobalStateAuthStore";
import { sendError } from "h3";

export default defineEventHandler(async (event) => {
  const store = useGlobalAuthStore()
  let userInfos

  await axios
    .get('http://localhost:8081/userinfos', {
      headers: {
        Authorization: `Bearer ${store.token}`,
      },
    })
    .then((res) => {
      userInfos = !res.error ? res : false
    })
    .catch((err) => {
        throw err
    })

  if (!userInfos) {
    // LOGic handling Error from Server
    return sendError(
      event,
      createError({
        statusCode: 400,
        statusMessage: `Not Valid: ${userSession.error}`,
      })
    )
  } else {
    return {
      // server return the idSession will use to establish cookie
      userInfos,
    }
  }
})
