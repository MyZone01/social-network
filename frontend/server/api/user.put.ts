import { sendError } from 'h3'
import { getSession } from "h3"

export default defineEventHandler(async (event) => {
    const body = await readBody(event)

    const token = event.headers.get('Authorization')
    console.log("FROM SERVER", body);

    const session = await getSession(event, {
        password: "5ec0312f-223f-4cc0-aa0f-303ff39fe1b2",
        name: "server-store",
        generateId: () => { return '' }
    })
    console.log(session)

    const response = await fetch('http://localhost:8081/edituser', {
        method: 'PUT',
        headers: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(body),
    })
    .then(async (res) => {
        await res.json()
        console.log(res)
        console.log(res.json())
    })
    .catch((err) => {
        console.log(err);
        return sendError(event, createError({
            statusCode: 500,
            statusMessage: 'Internal server error' + err
        }))
    });
})