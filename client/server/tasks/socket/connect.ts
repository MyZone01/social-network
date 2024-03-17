import { WebSocketClient } from "~/server/utils/server_socket";
import { setServerSocket, server_socket } from "~/stores/socketCon";

export default defineTask({
    async run() {
        if (!server_socket) {
            let ws = new WebSocketClient('ws://localhost:8081/socket?key=' + "socket");
            setServerSocket(ws)
        }
        return { result: 'ok' }
    }
})
