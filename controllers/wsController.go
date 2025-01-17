package controllers

import (
	"myapp/models"

	"github.com/gofiber/websocket/v2"

	"log"
	"sync"
)

// WebSocketController handles WebSocket connections
type WebSocketController struct {
	connections []*websocket.Conn // Active WebSocket connections
	mu          sync.Mutex        // Mutex for safe concurrent access
}

// NewWebSocketController initializes the WebSocketController
func NewWebSocketController() *WebSocketController {
	return &WebSocketController{}
}

// HandleWebSocketConnection handles incoming WebSocket connections and broadcasts messages
func (ws *WebSocketController) HandleWebSocketConnection(c *websocket.Conn) {
	// Lock to safely modify the connections slice
	ws.mu.Lock()
	ws.connections = append(ws.connections, c)
	ws.mu.Unlock()

	defer func() {
		// On close, remove the connection from the list
		ws.mu.Lock()
		for i, conn := range ws.connections {
			if conn == c {
				ws.connections = append(ws.connections[:i], ws.connections[i+1:]...)
				break
			}
		}
		ws.mu.Unlock()
		c.Close()
	}()

	// Read and broadcast messages
	for {
		msgType, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		// Create message model
		message := models.Message{
			Content: string(msg),
		}

		// Broadcast the message to all connected clients
		ws.broadcastMessage(msgType, []byte("Ticket number "+message.Content+"proceed to counter 1"))
	}
}

// broadcastMessage sends a message to all connected clients
func (ws *WebSocketController) broadcastMessage(msgType int, message []byte) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	for _, conn := range ws.connections {
		err := conn.WriteMessage(msgType, message)
		if err != nil {
			log.Println("Error broadcasting message:", err)
		}
	}
}
