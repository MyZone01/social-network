import { sendError, getSession, useSession } from 'h3'
import { sessionUpdater } from '../utils/createSession'
import { secure } from '../utils/transformer'

interface Response {
    message: string,
    status: number,
    user: object,
    session: string
}

export default defineEventHandler(async (event) => {
    const body = await readBody(event)
    const token = event.headers.get('Authorization')

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
            const response = await fetch('http://localhost:8081/updatepassword', {
                method: 'PUT',
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(body),
            })
            const result = await response.json()
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


