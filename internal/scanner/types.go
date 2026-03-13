package scanner

type Target struct {
	URL string
}

type Finding struct {
	Title       string
	Description string
	Severity    string
	Target      string
	Evidence    string
}

type ScanResult struct {
	Target   string
	Findings []Finding
}
