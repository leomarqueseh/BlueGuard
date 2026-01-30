package analysis

import "strings"

func MatchFingerprint(fp *Fingerprint, resp *httpxResult) bool {

	// AND logic
	for _, m := range fp.HTTP.Matchers {
		if !matchOne(m, resp) {
			return false
		}
	}

	// NEGATIVE logic
	for _, n := range fp.HTTP.Negative {
		if matchOne(n, resp) {
			return false
		}
	}

	return true
}

func matchOne(m Matcher, resp *httpxResult) bool {
	switch m.Type {

	case "body":
		for _, w := range m.Words {
			if strings.Contains(resp.Body, w) {
				return true
			}
		}

	case "status":
		for _, s := range m.Status {
			if resp.StatusCode == s {
				return true
			}
		}
	}

	return false
}
