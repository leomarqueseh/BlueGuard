package httpx

import (
	"crypto/tls"
	"io"
	"net/http"
	"time"
)

func Fetch(host string, timeout time.Duration) (*Result, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // necessário em takeover
		},
	}

	client := &http.Client{
		Timeout: timeout,
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 5 {
				return http.ErrUseLastResponse
			}
			return nil
		},
	}

	// 👉 HTTPS FIRST
	urls := []string{
		"https://" + host,
		"http://" + host,
	}

	for _, u := range urls {

		resp, err := client.Get(u)
		if err != nil {
			continue
		}

		defer resp.Body.Close()

		bodyBytes, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))

		return &Result{
			URL:        u,
			StatusCode: resp.StatusCode,
			Headers:    resp.Header,
			Body:       string(bodyBytes),
		}, nil
	}

	return nil, nil
}
