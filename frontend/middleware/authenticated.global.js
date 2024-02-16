// // import { navigateTo } from "nuxt/app"

export default defineNuxtRouteMiddleware((to, from) => {
    const authenticated = (event) => {
        if (getCookie(event), "token"){}
      }
    if (false) {
        return navigateTo("/auth")
    } else {
        return navigateTo("/")
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