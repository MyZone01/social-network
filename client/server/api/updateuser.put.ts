import { sendError, getSession, defineEventHandler, useSession } from 'h3'
import { secure } from '../utils/transformer'
import { fetcher } from '../utils/fetcher';
import { sessionUpdater } from '../utils/createHandler'

export default defineEventHandler(async (event) => {
    const body = await readBody(event)
    const token = event.context.token

    const session = await getSession(event, {
        password: "5ec0312f-223f-4cc0-aa0f-303ff39fe1b2",
        name: "server-store",
    })
    
    console.log("Body from Client", body, token)
    console.log(token, "<=====>" , session.data.sessionToken)
    console.log("Body from Server", session.data.userInfos)
    if (session.data.sessionToken != token) {
        return sendError(event, createError({
            statusCode: 400,
            statusMessage: 'No user session available'
        }))
    } else {
        try {
            body["password"] = session.data.userInfos.password
            console.log("Response", body)
            const result = await fetcher('http://localhost:8081/edituser', "PUT", JSON.stringify(body), token)
            await sessionUpdater(result.session, result.user, event)
            const { password: _password, ...userWithoutPassword } = result.user;
            const cleanInfos = {
                session: result.session,
                message: result.message,
                user: userWithoutPassword,
            }
            return cleanInfos
        } catch (err) {
            return sendError(event, createError({
                statusCode: 500,
                statusMessage: 'Internal server error' + err
        }))
        }
    }
});


