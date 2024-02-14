export default defineEventHandler(async (event) => {
    const body = await readBody(event)

    let userData
    const { loginAccess, registration } = body
    if (loginAccess) {
        userData = loginAccess
    } else if (registration) {
        userData = registration
    }
    
    if (!true) {
        return sendError(event, createError({
            statusCode: 400,
            statusMessage: 'This is the Error from the Server'
        }))
    }
    // Token LOGIC IMPLEMENTATION
    // generate Token
    // Access Token
    // refresh Token
    
    // const { accessToken, refreshToken } = generateTokens(user)
    // Save refreshToken inside db
    // await createRefreshToken({
        // token: refreshToken,
        // userId: user.id
    // })
    
    // add http only cookie
    // sendRefreshToken(event, refreshToken)

    return {
        body: userData
    }
})