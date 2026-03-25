package stats

import (
	"log"
	"sync"

	"github.com/kiing-dom/live-code-stats/internal/types"
)

var curr = types.Stats{}
var mu sync.Mutex

func UpdateStats(delta types.StatsDelta) types.Stats {
	mu.Lock()
	defer mu.Unlock()

	if delta.Lines != nil {
		curr.Lines = *delta.Lines
	}
	if delta.Errors != nil {
		curr.Errors = *delta.Errors
	}
	if delta.Keystrokes != nil {
		curr.Keystrokes = *delta.Keystrokes
	}
	if delta.FileName != nil {
		curr.FileName = *delta.FileName
	}

	updated := curr
	log.Printf("[stats] lines=%d errors=%d keystrokes=%d file=%s", updated.Lines, updated.Errors, updated.Keystrokes, updated.FileName)

	return updated
}

func GetStats() types.Stats {
	mu.Lock()
	defer mu.Unlock()

	return curr
}
