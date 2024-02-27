export default defineEventHandler(async (event) => {
    const body = await readBody(event);
    const token = body.token

    try {
        await useFetch('http://localhost:8081/checksession', {
            headers: {
                "Authorization": `Bearer ${token}`
            }
        })
        return true
    } catch (error) {
        return false
    }
})