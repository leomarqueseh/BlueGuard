package httpx

import "net/http"

// Result representa uma resposta HTTP padronizada
// usada pelo scanner e adaptada pela analysis
type Result struct {
	URL        string
	StatusCode int
	Headers    http.Header
	Body       string
}
