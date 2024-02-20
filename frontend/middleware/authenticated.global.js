import axios from "axios";

<<<<<<< HEAD
export default defineNuxtRouteMiddleware((to, from) => {
    const authStore = useGlobalAuthStore()
    const isAuthenticated = true

  if (!isAuthenticated && to.path !== '/auth') {
    return navigateTo('/auth')
  }
  if (isAuthenticated && to.path === '/auth') {
    return navigateTo('/feed')
  }
})

// export default defineEventHandler(async (event) => {
//     const endpoints = [
//         '/api/auth/user'
//     ]

//     const isHandledByThisMiddleware = endpoints.some(endpoint => {
//         const pattern = new UrlPattern(endpoint)

//         return pattern.match(event.req.url)
//     })

//     if (!isHandledByThisMiddleware) {
//         return
//     }

//     const token = event.req.headers['authorization']?.split(' ')[1]

//     const decoded = decodeAccessToken(token)

//     if (!decoded) {
//         return sendError(event, createError({
//             statusCode: 401,
//             statusMessage: 'Unauthorized'
//         }))
//     }

//     try {
//         const userId = decoded.userId
//         const user = await getuserById(userId)

//         event.context.auth = { user }
//     } catch (error) {
//         return
//     }
// })
=======
export default defineNuxtRouteMiddleware(async (to, from) => {
  const authStore = useGlobalAuthStore();
  let tokenValid;
  const authenticate = authStore.isAuthenticated;

  await axios
    .get("http://localhost:8081/checksession", {
      headers: {
        Authorization: `Bearer ${authStore.token}`,
      },
    })
    .then((res) => {
      tokenValid = !res.error ? true : false;
    })
    .catch((err) => {
      return;
    });

  const isAuthenticated = authenticate && tokenValid;
  console.log(isAuthenticated)
  if (!isAuthenticated && to.path !== "/auth") {
    authStore.logout();
    return navigateTo("/auth");
  }

  if (isAuthenticated && to.path === "/auth") {
    return navigateTo("/");
  }
});
>>>>>>> origin/master
