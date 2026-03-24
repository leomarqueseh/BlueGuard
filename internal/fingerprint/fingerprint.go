package fingerprint

import "strings"

type Result struct {
	Technologies []string
}

func Detect(headers map[string][]string, body string) Result {

	var tech []string

	// 🔹 Server header
	if server, ok := headers["Server"]; ok {
		s := strings.ToLower(strings.Join(server, " "))

		if strings.Contains(s, "nginx") {
			tech = append(tech, "nginx")
		}

		if strings.Contains(s, "apache") {
			tech = append(tech, "apache")
		}

		if strings.Contains(s, "cloudflare") {
			tech = append(tech, "cloudflare")
		}
	}

	// 🔹 X-Powered-By
	if xp, ok := headers["X-Powered-By"]; ok {
		x := strings.ToLower(strings.Join(xp, " "))

		if strings.Contains(x, "php") {
			tech = append(tech, "php")
		}

		if strings.Contains(x, "asp.net") {
			tech = append(tech, "asp.net")
		}
	}

	// 🔹 Body detection (básico)
	bodyLower := strings.ToLower(body)

	if strings.Contains(bodyLower, "wp-content") {
		tech = append(tech, "wordpress")
	}

	return Result{
		Technologies: unique(tech),
	}
}

func unique(input []string) []string {

	keys := make(map[string]bool)
	var result []string

	for _, v := range input {
		if !keys[v] {
			keys[v] = true
			result = append(result, v)
		}
	}

	return result
}
