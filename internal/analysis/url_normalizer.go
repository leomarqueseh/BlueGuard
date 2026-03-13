package analysis

import (
	"net/url"
	"strings"
)

func NormalizeURL(raw string) (host string, path string, params map[string]string, base string, ok bool) {

	u, err := url.Parse(raw)
	if err != nil || u.Host == "" {
		return "", "", nil, "", false
	}

	host = u.Host
	path = u.Path
	if path == "" {
		path = "/"
	}

	params = make(map[string]string)

	q := u.Query()
	for key := range q {
		params[key] = "FUZZ"
	}

	// Base URL normalizada
	if len(params) > 0 {
		var b strings.Builder
		b.WriteString(u.Scheme + "://" + host + path + "?")
		first := true
		for k := range params {
			if !first {
				b.WriteString("&")
			}
			b.WriteString(k + "=FUZZ")
			first = false
		}
		base = b.String()
	} else {
		base = u.Scheme + "://" + host + path
	}

	return host, path, params, base, true
}
