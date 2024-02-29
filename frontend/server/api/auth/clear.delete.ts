import { useSession } from 'h3'

export default defineEventHandler(async (event) => {

    const session = await useSession(event, {
        password: "5ec0312f-223f-4cc0-aa0f-303ff39fe1b2",
        name: "server-store",
        generateId: () => { return '' }
    })

    await session.clear();
})