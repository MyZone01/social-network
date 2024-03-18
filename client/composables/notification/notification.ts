export const connNotifSocket = async (ws: WebSocket | undefined, id: string = "") => {
    const isSecure = location.protocol === "https:";
    const url = (isSecure ? "wss://" : "ws://") + location.host + `/api/socket?userId=${id}&channel=notif`;
    if (ws) {
        ws.close();
    }

    ws = new WebSocket(url);
    return ws;
}

export const useClearNotif = async (id: string, action: string = "clear") => {
    const data = { type: action, id: id };    
    const response = await fetch("/api/notification/clearnotification", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    }).then(async (res) => await res.json()).catch((err) => {
        return {
            status: 500,
            body: "Internal server error",
        };
    });
    return response;
}
