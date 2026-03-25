package stats

import (
	"sync"

	"github.com/kiing-dom/live-code-stats/internal/types"
)

var curr = types.Stats{}
var mu sync.Mutex

func UpdateStats(delta types.Stats) types.Stats {
	mu.Lock()
	defer mu.Unlock()

	curr.Lines += delta.Lines
	curr.Errors += delta.Errors
	curr.Keystrokes += delta.Keystrokes

	updated := curr
	mu.Unlock()

	return updated
}

func GetStats() types.Stats {
	mu.Lock()
	defer mu.Unlock()

	return curr
}
