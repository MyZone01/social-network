export default defineEventHandler(async (event) => {
    const token = event.context.token;
    console.log("token: " + token + ``);
    if (!token) {
        return {
            status: 401,
            body: 'Unauthorized',
        };
    }
    
    
    const response = await fetch('http://localhost:8081/getAllFollowers', {
        method: 'GET',
        headers: {
            Authorization: `Bearer ${token}`,
        },
    }).then(async (res) => {
        return await res.json()
    }).catch((err) => {
        console.log(err);
        return {
            status: 500,
            body: 'Internal server error',
        };
    });

    console.log("res from server", response)
    if (response.status !== 200) {
        return {
            status: response.status,
            body: response.message,
        };
    }

    return {
        status: 200,
        body: response.data,
    };
});