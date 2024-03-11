

export default defineEventHandler(async (event) => {
    const body = await readBody(event);
    const token = event.context.token;
    if (!token) {
        return {
            status: 401,
            body: 'Unauthorized',
        };
    }
    
    const response = await fetcher('http://localhost:8081/follower', 'POST', JSON.stringify(body), `Bearer ${token}`);
    if (response.status !== 200) {
        return {
            status: response.status,
            body: response.message,
        };
    }
    return {
        status: 200,
        body: response.message,
    };


});