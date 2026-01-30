package analysis

import (
	"fmt"
	"sync"
)

var (
	Verbose bool
	mu      sync.Mutex
)

// logf imprime logs apenas se verbose estiver ativo
func logf(format string, args ...any) {
	if !Verbose {
		return
	}

	mu.Lock()
	defer mu.Unlock()

	fmt.Printf("[DEBUG] "+format+"\n", args...)
}
