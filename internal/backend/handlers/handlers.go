package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kiing-dom/live-code-stats/internal/backend/stats"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var delta stats.Stats
	json.NewDecoder(r.Body).Decode(&delta)

	stats.UpdateStats(delta)

	w.WriteHeader(http.StatusOK)
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(stats.GetStats())
}
