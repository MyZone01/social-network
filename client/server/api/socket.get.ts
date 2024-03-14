import WebSocket from "ws";

let wss: WebSocket.Server | null = null;

export default defineEventHandler(async (event) => {
    if (!wss) {
        let server = { server: event.node.res.socket?.server };
        wss = new WebSocket.Server(server);
        wss.on('connection', function connection(ws) {
            ws.on('message', function incoming(message) {
                console.log('received: %s', message);
            });
            ws.send('something');
        });
    }
});