export default defineEventHandler(async event => {
    try {
        if (!event.context.token) {
            return createError({
                statusCode: 401,
                message: "You don't have the rights to access this resource",
            });
        }

        const token = event.context.token;
        const queryObj = getQuery(event)
        const groupId = queryObj.gid


        const response = await $fetch('http://localhost:8081/demand-access', {
            method: 'POST',
            headers: {
                Authorization: `${token}`
            },
            query: {
                group_id: groupId,
            }
        })
        return response

    } catch (e: any) {
        console.log(e)
        setResponseStatus(event, 500, e.message)
    }

})