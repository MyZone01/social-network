export default defineEventHandler(async event => {
    try {
        const token = getHeader(event,'Authorization')
        const queryObj = getQuery(event)
        const groupId = queryObj.gid


        const response = await $fetch(`${process.env.BACKEND_URL}`+"/get-all-access-demand', {
            method: 'GET',
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
        setResponseStatus(event,500,e.message)
    }

})