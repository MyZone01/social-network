export const connNotifSocket = async (ws: WebSocket | undefined, id: string = "") => {
    const isSecure = location.protocol === "https:";
    const url = (isSecure ? "wss://" : "ws://") + location.host + `/api/socket?userId=${id}&channel=notif`;
    if (ws) {
        ws.close();
    }

    ws = new WebSocket(url);
    return ws;
}