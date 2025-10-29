let socket: WebSocket | null = null;

export function connectWebSocket(url: string): void {
  socket = new WebSocket("ws://localhost:8081/ws/echo");

    socket.onopen = () => {
        socket ?.send("Hello from frontend");
    }
    socket.onmessage = (event) => {
        onMessage(event.data)
    }
    socket.onerror = (err) => {
        console.error("WebSocket error:", err);
    }
    socket.onclose = () => {
        console.log("Websocket closed")
    }
}