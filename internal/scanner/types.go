package scanner

type Target struct {
	URL string
}

type Finding struct {
	Title       string
	Description string
	Severity    string
	Target      string
	Score       float64

	// 🔥 NOVO (PoC)
	Confirmed bool
}
