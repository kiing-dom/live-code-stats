package backend

import "sync"

type Stats struct {
	Lines      int `json:"lines"`
	Errors     int `json:"errors"`
	Keystrokes int `json:"keystrokes"`
}

var stats = Stats{}
var mu sync.Mutex

func UpdateStats(delta Stats) {
	mu.Lock()
	defer mu.Unlock()

	stats.Lines += delta.Lines
	stats.Errors += delta.Errors
	stats.Keystrokes += delta.Keystrokes
}

func GetStats() Stats {
	mu.Lock()
	defer mu.Unlock()

	return stats
}
