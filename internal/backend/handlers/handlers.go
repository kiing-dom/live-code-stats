package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kiing-dom/live-code-stats/internal/backend/stats"
	"github.com/kiing-dom/live-code-stats/internal/backend/websocket"
	"github.com/kiing-dom/live-code-stats/internal/types"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var delta types.StatsDelta
	if err := json.NewDecoder(r.Body).Decode(&delta); err != nil {
		log.Printf("[handlers] failed to decode update payload: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updated := stats.UpdateStats(delta)
	websocket.Broadcast(updated)

	w.WriteHeader(http.StatusOK)
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[handlers] stats requested")
	json.NewEncoder(w).Encode(stats.GetStats())
}
