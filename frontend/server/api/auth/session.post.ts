import { getSession, useSession } from "h3"
import { useGlobalAuthStore } from "@/stores/useGlobalStateAuthStore";

export default defineEventHandler(async (event) => {
    const body = await readBody(event);
    const store = useGlobalAuthStore()
    const token = body.token

    const session = await useSession(event, {
        password: "5ec0312f-223f-4cc0-aa0f-303ff39fe1b2",
        name: "server-store",
    })
    try {
        const sessionId = await getSession(event, {
            password: "5ec0312f-223f-4cc0-aa0f-303ff39fe1b2",
            name: "server-store",
            generateId: () => { return '' }
        })
        console.log(sessionId.data.userInfos)
        if (sessionId.id != token) {
            await store.logout()
            await session.clear()
            return false
        }
        const response = await $fetch('http://localhost:8081/checksession', {
            headers: {
                "Authorization": `Bearer ${token}`
            }
        })
        console.log(response)
        return true
    } catch (error) {
        await store.logout()
        await session.clear()
        return false
    }
})