package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Hub struct {
	clients map[string]*websocket.Conn
	mu      sync.Mutex
}

type Message struct {
	SenderId   string `json:"senderId"`
	ReceiverId string `json:"receiverId"`
	Message    string `json:"message"`
	Type       string `json:"type"`
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[string]*websocket.Conn),
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Hub) HandleConnection(c *gin.Context) {
	userID := c.Query("userId")

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "userId is required",
		})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}

	h.mu.Lock()
	h.clients[userID] = conn
	h.mu.Unlock()

	log.Println("User connected:", userID)

	defer func() {
		h.mu.Lock()

		if h.clients[userID] == conn {
			delete(h.clients, userID)
		}

		h.mu.Unlock()

		conn.Close()
		log.Println("User disconnected:", userID)
	}()

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket disconnected:", err)
			break
		}
		var message Message

		if err := json.Unmarshal(data, &message); err != nil {
			log.Println("Invalid message:", err)
			continue
		}

		if message.ReceiverId == "" || message.Message == "" {
			log.Println("Invalid receiver or empty message")
			continue
		}

		message.SenderId = userID
		message.Type = "chat"

		response, err := json.Marshal(message)
		if err != nil {
			log.Println("JSON encode error:", err)
			continue
		}

		h.SendToUser(userID, response)

		if message.ReceiverId != userID {
			h.SendToUser(message.ReceiverId, response)
		}
	}
}

func (h *Hub) SendToUser(userID string, message []byte) {
	h.mu.Lock()
	defer h.mu.Unlock()

	conn, exists := h.clients[userID]
	if !exists {
		log.Println("User not connected:", userID)
		return
	}

	if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
		log.Println("WebSocket write error:", err)

		delete(h.clients, userID)
		conn.Close()
	}
}
