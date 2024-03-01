import { getSession, sendError } from "h3"

export default defineEventHandler(async (event) => {
    const body = await readBody(event);
    const token = body.token

    const sessionId = await getSession(event, {
        password: "5ec0312f-223f-4cc0-aa0f-303ff39fe1b2",
        name: "server-store",
        generateId: () => { return '' }
    })
    if (token == sessionId.id) {
        return { user: sessionId.data.userInfos }
    } else {
        sendError(event, createError({
            statusCode: 400,
            statusMessage: "Bad Request: No user informations available."
        }))
    }
})