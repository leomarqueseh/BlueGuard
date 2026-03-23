package httpclient

import (
	"net"
	"net/http"
	"time"
)

type Client struct {
	http *http.Client
}

func New(timeout time.Duration) *Client {

	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: timeout,
		}).DialContext,
		TLSHandshakeTimeout: timeout,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
	}

	return &Client{
		http: &http.Client{
			Timeout:   timeout,
			Transport: transport,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				if len(via) >= 5 {
					return http.ErrUseLastResponse
				}
				return nil
			},
		},
	}
}
