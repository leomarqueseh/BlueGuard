package logger

import (
	"log"
	"os"
)

var Logger *log.Logger

func Init() {
	Logger = log.New(os.Stdout, "[BLUEGUARD] ", log.LstdFlags|log.Lshortfile)
}
