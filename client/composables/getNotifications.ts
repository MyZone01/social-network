export const getNotifications = async () => {
    const res = await fetch("/api/notification").then()
    const responseInJsonFormat = await res.json().catch(err => ({ error: err }))

    return responseInJsonFormat
}