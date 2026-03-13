package logging

import (
	"fmt"
	"sync"
)

var (
	Verbose bool
	mu sync.Mutex
)

func Info(format string, args ...any) {
	log("[INFO] ", format, args...)
}

func Debug(format string, args ...any) {
	if !Verbose {
		return
	}
	log("[DEBUG] ", format, args...)
}

func Warn(format string, args ...any) {
	log("[WARN] ", format, args...)
}

func log(prefix, format string, args ...any) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Printf(prefix+format+"\n", args...)
}
