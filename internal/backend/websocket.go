package backend

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/kiing-dom/live-code-stats/internal/backend/stats"
)

var clients = make(map[*websocket.Conn]bool)
var upgrader = websocket.Upgrader{}

func Broadcast() {
	for client := range clients {
		err := client.WriteJSON(stats.GetStats())
		if err != nil {
			fmt.Println(err)

			client.Close()
			delete(clients, client)
		}
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	clients[conn] = true
}
