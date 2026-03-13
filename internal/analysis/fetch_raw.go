package analysis

import (
	"crypto/tls"
	"io"
	"net/http"
	"time"
)

type RawResponse struct {
	URL        string
	StatusCode int
	Body       string
}

// FetchRaw executa uma requisição direta (usado em fuzzing leve)
func FetchRaw(url string) (*RawResponse, error) {

	client := &http.Client{
		Timeout: 8 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))

	return &RawResponse{
		URL:        url,
		StatusCode: resp.StatusCode,
		Body:       string(bodyBytes),
	}, nil
}
