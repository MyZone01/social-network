import { clearSession } from 'h3'

export default defineEventHandler(async (event) => {

    const sesssion = await getSession(event, {
        password: "5ec0312f-223f-4cc0-aa0f-303ff39fe1b2",
        name: "server-store",
        generateId: () => { return '' }
    })

    // const session = await clearSession(event, {
    //     password: "5ec0312f-223f-4cc0-aa0f-303ff39fe1b2",
    //     name: "server-store",
    //     generateId: () => { return '' }
    // })

    const session = await useSession(event, {
        password: "5ec0312f-223f-4cc0-aa0f-303ff39fe1b2",
    });
    
    await session.clear();

})