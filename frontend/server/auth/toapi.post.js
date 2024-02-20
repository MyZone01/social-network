import sendError from 'h3'

export default defineEventHandler(async (event) => {
    const body = await readBody(event)

    const { email, password} = body

    const requiredFields = [email, password]
    if (requiredFields.some(field => field == '')) {
        return sendError(event, createError({
            statusCode: 400,
            statusMessage: 'Invalid params'
        }))
    }

    // Logic checking user registered and Password matched
    // Could generate new accessToken abd refreshSessionToken

    // temporary for ogin test only
    alert('passed')
    return {
        hello: 'welcome'
    }
})