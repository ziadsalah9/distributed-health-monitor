package websocket

import (
	"sync"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Hub struct {
	clients map[*websocket.Conn]bool
	broadcast chan interface{}
	mu sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan interface{}),
	}
}

func (h *Hub) Run() {
	for {
		msg := <-h.broadcast
		h.mu.Lock()
		for client := range h.clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(h.clients, client)
			}
		}
		h.mu.Unlock()
	}
}

func (h *Hub) HandleWS(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	h.mu.Lock()
	h.clients[conn] = true
	h.mu.Unlock()
}

func (h *Hub) BroadcastUpdate(data interface{}) {
	h.broadcast <- data
}