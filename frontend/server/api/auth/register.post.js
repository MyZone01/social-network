import { sendError } from "h3"
// import { createUser } from "~/server/db/users.js"
import { userTransformer } from "~/server/transformers/user.js"

export default defineEventHandler(async (event) => {
    const body = await readBody(event)
    // console.log(body)
    
    const { firstName, lastName, email, nickname, password, repeatPassword, aboutMe, avatarImg } = body
    
    const requiredFields = [ firstName, lastName, email, nickname, password, repeatPassword, aboutMe, avatarImg ];
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

    // VALIDATION LOGIC to Implement
    
    // LOGIC TO STORE USER into DataBase
    // const user = await createUser(userData)
    
    
    // better use Transformer for user Infos exposition
    const userData = {
        nickname,
        email,
        firstName,
        lastName,
        avatarImg,
        aboutMe
    }
    return {
        body: userData
    }

})