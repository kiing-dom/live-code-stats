package backend

import (
	"encoding/json"
	"net/http"
)

func updateHandler(w http.ResponseWriter, r *http.Request) {
	var delta Stats
	json.NewDecoder(r.Body).Decode(&delta)

	UpdateStats(delta)

	w.WriteHeader(http.StatusOK)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(GetStats())
}
