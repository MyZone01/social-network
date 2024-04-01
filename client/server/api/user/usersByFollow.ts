

export default defineEventHandler(async (event) => {

    const token = event.context.token;
    if (!token) {
        return {
            status: 401,
            body: 'Unauthorized',
        };
    }
    const response = await fetch('http://localhost:8081/usersByFollow', {
        headers: {
            Authorization: `Bearer ${token}`,
        }
    });
    // console.log(response);
    return response
})
