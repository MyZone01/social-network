// import { useStore } from "pinia";
// import { useGlobalAuthStore } from "@/stores/useGlobalStateAuthStore";
// // import { navigateTo } from "nuxt/app"

export default defineNuxtRouteMiddleware((to, from) => {
    const authStore = useGlobalAuthStore();;
    const isAuthenticated = authStore.isAuthenticated;

  if (!isAuthenticated && to.path !== "/auth") {
    return navigateTo("/auth");
  }
  if (isAuthenticated && to.path === "/auth") {
    return navigateTo("/");
  }

});

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
