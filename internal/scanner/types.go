package scanner

import "time"

type Target struct {
	URL string
}

type Finding struct {
	PluginID    string
	Title       string
	Description string
	Severity    string
	Evidence    string
	Timestamp   time.Time
}

type ScanResult struct {
	Target   string
	Findings []Finding
}
