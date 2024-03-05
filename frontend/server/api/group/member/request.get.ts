export default defineEventHandler(async event => {
    try {
        const gid = getRouterParam(event, 'id')
        const token = getHeader(event,'Authorization')
        const queryObj = getQuery(event)
        const groupId = gid
        const ismember = queryObj.m === '1'
        const isuser = queryObj.u === '1'


        const response = await $fetch("http://localhost:8081/get-group", {
            method: 'GET',
            headers: {
                'Authorization': `${token}`
            },
            query: {
                group_id: groupId,
                isMemberNeeded: ismember,
                isUserNeeded: isuser
            }
        })
        return response

    } catch (e: any) {
        console.log(e);
        setResponseStatus(event,500,e.message)
    }

})