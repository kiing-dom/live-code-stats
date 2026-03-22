package main

import (
	"fmt"
	"net/http"

	"github.com/kiing-dom/live-code-stats/internal/backend/handlers"
)

func main() {
	http.HandleFunc("/update", handlers.UpdateHandler)
	http.HandleFunc("/stats", handlers.StatsHandler)

	fmt.Println("server running on :8080")
	http.ListenAndServe(":8080", nil)
}
