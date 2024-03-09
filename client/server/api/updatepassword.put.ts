import { sendError, getSession, useSession } from 'h3'
import { sessionUpdater } from '../utils/createHandler'
import { secure } from '../utils/transformer'
import { fetcher } from '../utils/fetcher'

export default defineEventHandler(async (event) => {
    const body = await readBody(event)
    const token = event.context.token

    const session = await getSession(event, {
        password: "5ec0312f-223f-4cc0-aa0f-303ff39fe1b2",
        name: "server-store",
        // generateId: () => { return '' }
    })
    
    if (session.data.sessionToken != token) {
        return sendError(event, createError({
            statusCode: 400,
            statusMessage: 'No user session available'
        }))
    } else {
        try {
            const result = await fetcher('http://localhost:8081/updatepassword', "PUT", JSON.stringify(body), token)
            await sessionUpdater(result.session, result.user, event)
            const cleanInfos = {
                session: result.session,
                message: result.message,
                user: secure(result.user),
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


