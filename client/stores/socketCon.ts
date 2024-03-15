import { WebSocketClient } from "~/server/utils/server_socket"

export let server_socket: WebSocketClient | undefined

export const setServerSocket = (socket: WebSocketClient) => {
    server_socket = socket
}
