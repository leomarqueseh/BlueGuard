package scanner

type Target struct {
	URL string
}

type Finding struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Severity    string `json:"severity"`
	Target      string `json:"target"`
	Score       float64 `json:"score"`
}
