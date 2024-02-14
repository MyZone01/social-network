import { sendError } from "h3"
// import { createUser } from "~/server/db/users.js"
import { userTransformer } from "~/server/transformers/user.js"

export default defineEventHandler(async (event) => {
    const body = await readBody(event)

    const { nickname, email, password, repeatPassword, name } = body

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
        password,
        name, 
        profileImage: 'https://picsum.photos/200/200'
    }

    const user = await createUser(userData)
    return {
        body: userTransformer(user)
    }

})