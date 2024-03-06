import { getSession } from 'h3'

export default defineEventHandler(async (event) => {
    const body = await readBody(event)
    const token = body.token

    const sesssion = await getSession(event, {
        password: '5ec0312f-223f-4cc0-aa0f-303ff39fe1b2',
        name: 'server-store',
        generateId: () => { return '' }
    })

    console.log(sesssion)
    console.log(token)

    try {
        await useFetch('http://localhost:8081/checksession', {
            headers: {
                Authorization: `Bearer ${token}`
            }
        })
        return true
    } catch (error) {
        return false
    }
})