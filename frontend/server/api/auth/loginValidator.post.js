// import { getUserByEmail } from "~/server/db/users.js";
// import bcrypt from "bcrypt"
import { userTransformer } from '~/server/transformers/user.js'
// import { createRefreshToken } from "~/server/db/refreshToken.js";
// import { generateTokens, sendRefreshToken } from "~/server/utils/jwt.js";

export default defineEventHandler(async (event) => {
    const body = await readBody(event)

    const { email, password} = body

    const requiredFields = [email, password]
    if (requiredFields.some(field => field == '')) {
        return sendError(event, createError({
            statusCode: 400,
            statusMessage: 'Invalid entries'
        }))
    }

    // Check if User is registered
    const user = true //await getUserByEmail(email)

    if (!user) {
        return sendError(event, createError({
            statusCode: 400,
            statusMessage: 'Email or password is invalid'
        }))
    }

    // Compare password with password stored in DataBase
    const passwordMatched = true //await bcrypt.compare(password, user.password)
    if (!passwordMatched) {
        return sendError(event, createError({
            statusCode: 400,
            statusMessage: 'Email or password is invalid'
        }))
    }

    const loginAccess = {
        email: email,
        password: password
    }

    try {
        const userSession = await $fetch('http://localhost:8081/registration', {
            method: 'POST',
            body: JSON.stringify(loginAccess)
        })

        if (userSession.error) {
            // LOGic handling Error from Server
            return sendError(event, createError({
                statusCode: 400,
                statusMessage: error.statusMessage
            }))
        } else {
            return {
                // filtering what data to expose and which one to set from there
                body: userTransformer(userSession)
            }
        }
    } catch (error) {
        // relate It to the server error
        return sendError(event, createError({
            statusCode: 400,
            statusMessage: error.statusMessage
        }))
    }
})