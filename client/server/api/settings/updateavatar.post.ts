import { sendError, getSession, useSession } from 'h3'
import { sessionUpdater } from '../../utils/sessionHandler'
import { fetcher } from '../../utils/fetcher'

export default defineEventHandler(async (event) => {
    const reader = await readMultipartFormData(event);
    if (!reader) return { status: 400, body: 'Bad request', ok: false }

    const { file } = await processParts(reader);

    const body = new FormData();
    body.append('file', new Blob([file.data]), file.filename);

    console.log(body)

    const token = event.context.token
    
    const session = await getSession(event, {
        password: "5ec0312f-223f-4cc0-aa0f-303ff39fe1b2",
        name: "server-store",
        // generateId: () => { return '' }
    })
    
    if (session.data.sessionToken != token) {
        return sendError(event, createError({
            statusCode: 400,
            statusMessage: 'No user session available'
        }))
    } else {
        const updateValue = session.data.userInfos
        try {
            const _response = await fetch(`${process.env.BACKEND_URL}` + "/upload", {
                method: "POST",
                headers: {
                    Authorization: `Bearer ${token}`,
                },
                body
            });
            
            const response = JSON.parse(await _response.text()) as { imageurl: string };
            
            console.log(response)
            
            if (!response.imageurl) {
                return { status: 200, body: "Error while updating avatar image", session: "", ok: false };
            }
            
            updateValue.avatarImage = response.imageurl
            const _response2 = await fetch(`${process.env.BACKEND_URL}` + "/updateuser", {
                method: "PUT",
                headers: {
                  Accept: "application/json",
                  "Content-Type": "application/json",
                  Authorization: `Bearer ${token}`,
                },
                body: JSON.stringify(updateValue)
              });

              await sessionUpdater(token, _response2, event)
            return response.imageurl
    //         body['email'] = session.data.userInfos.email
    //         const result = await fetcher(`${process.env.BACKEND_URL}`+"/updatepassword", "PUT", JSON.stringify(body), token)

    //         const { password: _password, ...userWithoutPassword } = result.data;
    //         const cleanInfos = {
    //             message: result.message,
    //             user: userWithoutPassword,
    //         }
    //         return cleanInfos
        } catch (err) {
            console.log(err)
            return sendError(event, createError({
                statusCode: 500,
                statusMessage: 'Internal server error' + err
            }))
        }
    }
});


