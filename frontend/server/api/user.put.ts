export default defineEventHandler(async (event) => {
    const body = await readBody(event)
    const token = event.headers.get('Authorization')
    console.log(body);
    const repone = await fetch('http://localhost:8081/editeuser', {
        method: 'PUT',
        headers: {
            Authorization: `${token}`,
            'Content-Type': 'application/json',

        },
        body,
    }).then(async (res) => await res.json()).catch((err) => {
        console.log(err);
        return {
            status: 500,
            body: 'Internal server error',
        };
    });
})