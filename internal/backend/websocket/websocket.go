package websocket

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/kiing-dom/live-code-stats/internal/types"
)

var clients = make(map[*websocket.Conn]bool)
var upgrader = websocket.Upgrader{}

func Broadcast(data types.Stats) {
	for client := range clients {
		err := client.WriteJSON(data)
		if err != nil {
			fmt.Println(err)

			client.Close()
			delete(clients, client)
		}
	}
}

func HandleWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	clients[conn] = true
}
