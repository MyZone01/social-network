import { getSession, useSession, defineEventHandler } from "h3"

export default defineEventHandler(async (event) => {
    const body = await readBody(event);
    const token = body.token

    const sessionId = await getSession(event, {
        password: "5ec0312f-223f-4cc0-aa0f-303ff39fe1b2",
        name: 'server-store'
    })
    const checker = sessionId.data.sessionToken
    try {
        if (checker != token) {
            const session = await useSession(event, {
                password: "5ec0312f-223f-4cc0-aa0f-303ff39fe1b2",
                name: 'server-store'
            })
            await session.clear()
            return false
        }
        const response = await fetch('http://localhost:8081/checksession', {
            headers: {
                "Authorization": `Bearer ${token}`
            }
        })
        const data = await response.json()
        if (data.message) {
            return true
        } else {
            return false
        }
    } catch (error) {
        const session = await useSession(event, {
            password: "5ec0312f-223f-4cc0-aa0f-303ff39fe1b2",
            name: 'server-store'
        })
        await session.clear()
        return false
    }
})
// name: "server-store",
// generateId: () => { return token }