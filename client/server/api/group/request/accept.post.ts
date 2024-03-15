export default defineEventHandler(async event => {
    try {
        const token = getHeader(event,'Authorization')
        const queryObj = getQuery(event)
        const groupId = queryObj.gId
        const rId = queryObj.rId


        const response = await $fetch('http://localhost:8081/accept-access-demand', {
            method: 'POST',
            headers: {
                Authorization: `${token}`
            },

        query:{
            group_id:groupId,
            requesting_id:rId
        }
        })
        return response

    } catch (e: any) {
        console.log(e)
        setResponseStatus(event,500,e.message)
    }

})