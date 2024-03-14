class WebSocketClient {
    private socket: WebSocket | null = null;
    private readonly url: string;
    private readonly reconnectInterval: number;

    constructor(url: string, reconnectInterval = 5000) {
        this.url = url;
        this.reconnectInterval = reconnectInterval;
        this.connect();
    }

    private connect() {
        this.socket = new WebSocket(this.url);

        this.socket.onopen = (event: Event) => {
            console.log('WebSocket connection established');
        };

        this.socket.onerror = (error: Event) => {
            console.error('WebSocket error:', error);
        };

        this.socket.onclose = (event: CloseEvent) => {
            console.log('WebSocket connection closed. Reconnecting...');
            setTimeout(() => this.connect(), this.reconnectInterval);
        };
    }

    public onmessage(handler: (params: MessageEvent) => void) {
        if (this.socket && this.socket.readyState === WebSocket.OPEN) {
            this.socket.onmessage = handler;
        } else {
            console.error('Cannot set message handler, WebSocket is not open');
        }
    }
    public send(data: string) {
        if (this.socket && this.socket.readyState === WebSocket.OPEN) {
            this.socket.send(data);
        } else {
            console.error('Cannot send message, WebSocket is not open');
        }
    }
}

export const client = new WebSocketClient('ws://localhost:8080/socket');

client.onmessage((event) => {
    console.log('Message received:', event.data);
});
