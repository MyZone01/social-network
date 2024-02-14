// import { getUserByEmail } from "~/server/db/users.js";
// import bcrypt from "bcrypt"
import { userTransformer } from "~/server/transformers/user.js";
// import { createRefreshToken } from "~/server/db/refreshToken.js";
// import { generateTokens, sendRefreshToken } from "~/server/utils/jwt.js";

export default defineEventHandler(async (event) => {
    const body = await readBody(event)

    const { email, password} = body

    const requiredFields = [email, password];
    if (requiredFields.some(field => field == "")) {
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
        access_Token: "accessToken",
        user: userTransformer(user)
    }
} ) 