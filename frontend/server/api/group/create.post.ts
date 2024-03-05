export default defineEventHandler(async event => {
    const token = getHeader(event,'Authorization')

    try {
        const payload = await readBody(event)
        const response = await $fetch('http://localhost:8081/create-group', {
            method:'POST',
            headers: {
                'Authorization': `${token}`
            },
            body: JSON.stringify(payload)

        })
        console.log('from api/group/create');
        // event.node.res.writeHead(201, 'ok')
        return response


    } catch (e: any) {
        console.log(e);
        event.node.res.writeHead(500, e.message)
    }   

})