package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/kiing-dom/live-code-stats/internal/types"
)

var clients = make(map[*websocket.Conn]bool)
var mu sync.Mutex
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[websocket] upgrade error: %v", err)
		return
	}

	mu.Lock()
	clients[conn] = true
	mu.Unlock()
	log.Printf("[websocket] client connected (%d total)", len(clients))

	go func() {
		defer func() {
			mu.Lock()
			delete(clients, conn)
			mu.Unlock()
			conn.Close()
			log.Printf("[websocket] client disconnected (%d total)", len(clients))
		}()
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				break
			}
		}
	}()
}

func Broadcast(data types.Stats) {
	mu.Lock()
	defer mu.Unlock()

	for client := range clients {
		err := client.WriteJSON(data)
		if err != nil {
			log.Printf("[websocket] write error, dropping client: %v", err)
			client.Close()
			delete(clients, client)
		}
	}
}
