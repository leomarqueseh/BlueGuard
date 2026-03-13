package core

import "time"

type ScanContext struct {
	Target    string
	Timeout   time.Duration
	UserAgent string
}
