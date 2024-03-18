export const getMessage = async (receiverId: string) => {


    const response = await fetch('/api/getReceiver', {
        body: JSON.stringify({
            receiver_id: receiverId,
        }),
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        }
    }).then(async (res) => await res.json()).catch((err) => {
        return {
            status: 500,
            body: 'Internal server error',
        };
    })
    console.log(response)
    return response
}
