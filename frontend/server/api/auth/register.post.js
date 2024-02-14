import { sendError } from "h3"
// import { createUser } from "~/server/db/users.js"
import { userTransformer } from "~/server/transformers/user.js"

export default defineEventHandler(async (event) => {
    const body = await readBody(event)
    console.log(body)
    
    const { firstName, lastName, email, nickname, password, repeatPassword, aboutMe, avatarImg} = body
    console.log(firstName)

    const requiredFields = [nickname, email, password, repeatPassword, name];
    if (requiredFields.some(field => field == "")) {
        const errorMessage = 'Invalid params';
        return sendError(event, createError({ statusCode: 400, statusMessage: errorMessage }));
    }

    if (password !== repeatPassword) {
        return sendError(event, createError({
            statusCode: 400,
            statusMessage: 'Passwords do not match'
        }))
    }

    const userData = {
        nickname,
        email,
        firstName,
        lastName,
        avatarImg,
        aboutMe
    }

    // const user = await createUser(userData)
    return {
        body: userData
    }

})