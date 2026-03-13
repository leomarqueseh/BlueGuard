package httpx

type Result struct {
	URL        string
	StatusCode int
	Body       []byte
	Headers    map[string][]string
}

