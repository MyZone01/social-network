import { Register } from '@/server/models/register'
import { fetcher } from '@/server/utils/fetcher'
import { processParts } from '@/server/utils/processParts'

export default defineEventHandler(async (event) => {
  const reader = await readMultipartFormData(event)
  if (!reader) return { status: 400, body: 'Bad request', ok: false }

  const { file, jsonData } = await processParts(reader)

  const register = new Register(jsonData)
  const [isValid, message] = register.validate()
  if (!isValid) {
    return sendError(event, createError({
      statusCode: 400,
      statusMessage: message
    }))
  }

  const response = await fetcher('http://localhost:8081/registration', 'POST', JSON.stringify(register), '')
  if (response.status !== '200') {
    return sendError(event, createError({
      statusCode: 400,
      statusMessage: response.message
    }))
  }

  if (!file) {
    console.log('No File given')
    return {
      status: 200,
      body: 'No file uploaded',
      session: response.session,
      user: response.user,
      ok: true
    }
  }

  const body = new FormData()
  body.append('file', new Blob([file.data]), file.filename)

  const response2 = await fetcher('http://localhost:8081/upload', 'POST', body, response.session)
  if (!response2.imageurl) {
    return { status: 200, body: response.message, session: response.session, ok: false }
  }

  register.avatarImage = response2.imageurl

  const response3 = await fetcher('http://localhost:8081/updateuser', 'PUT', JSON.stringify(register), response.session)
  if (response3.status !== '200') {
    return { status: 200, body: response3.message, session: response.session, ok: false }
  }
  return {
    status: 200,
    body: 'User registered successfully',
    session: response.session,
    user: response.user,
    ok: true
  }
})