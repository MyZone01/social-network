import { sendError, getSession, useSession } from 'h3'
import { sessionUpdater } from '../utils/sessionHandler'
import { fetcher } from '../utils/fetcher'

export default defineEventHandler(async (event) => {
    const body = await readBody(event)
    const token = event.context.token

    const session = await getSession(event, {
        password: "5ec0312f-223f-4cc0-aa0f-303ff39fe1b2",
        name: "server-store",
        // generateId: () => { return '' }
    })
    console.log(session.data.sessionToken, "<======>", token)
    
    if (session.data.sessionToken != token) {
        return sendError(event, createError({
            statusCode: 400,
            statusMessage: 'No user session available'
        }))
    } else {
        try {
            const result = await fetcher('http://localhost:8081/updatepassword', "PUT", JSON.stringify(body), token)
            await sessionUpdater(token, result.data, event)

            console.log(result)
            const { password: _password, ...userWithoutPassword } = result.data;
            const cleanInfos = {
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

