package httpclient

import (
	"io"
	"net/http"
)

type Response struct {
	StatusCode int
	Body       []byte
	Headers    http.Header
	URL        string
}

func (c *Client) Get(url string, ua string) (*Response, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", ua)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	return &Response{
		StatusCode: resp.StatusCode,
		Body:       body,
		Headers:    resp.Header,
		URL:        resp.Request.URL.String(),
	}, nil
}
