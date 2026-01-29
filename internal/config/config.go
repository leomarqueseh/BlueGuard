package config

import "time"

type ScanConfig struct {
	Target string

	Passive     bool
	PassiveLite bool
	Active      bool
	Full        bool
	Stealth     bool

	Rate    int
	Delay   time.Duration
	Timeout time.Duration

	HeadOnly  bool
	NoAlive   bool
	NoCrawler bool
}
