package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kiing-dom/live-code-stats/internal/backend/stats"
	"github.com/kiing-dom/live-code-stats/internal/backend/websocket"
	"github.com/kiing-dom/live-code-stats/internal/types"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var delta types.Stats
	json.NewDecoder(r.Body).Decode(&delta)

	updated := stats.UpdateStats(delta)

	websocket.Broadcast(updated)

	w.WriteHeader(http.StatusOK)
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(stats.GetStats())
}
