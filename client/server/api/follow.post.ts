

export default defineEventHandler(async (event) => {
    const body = await readBody(event);
    const token = event.context.token;
    if (!token) {
        return {
            status: 401,
            body: 'Unauthorized',
        };
    }

    const response = await fetcher('http://localhost:8081/follower', 'POST', JSON.stringify(body), token);
    return response
   
});