export default defineEventHandler(async (event) => {
    const body = await readBody(event)
    const token = event.context.token;
    if (!token) {
        return {
            status: 401,
            body: 'Unauthorized',
        };
    }
    console.log(body)
    const response = await fetch('http://localhost:8081/getMessages', {
        method: 'POST',
        headers: {
            Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(body),
       
    }).then(async (res) => {
        return await res.json()
    }).catch((err) => {
        console.log(err, "loololo");
        return {
            status: 500,
            body: 'Internal server error',
        };
    });
   
   
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