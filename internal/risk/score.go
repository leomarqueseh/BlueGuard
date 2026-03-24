package risk

import "github.com/leomarqueseh/BlueGuard/internal/scanner"

func CalculateScore(f scanner.Finding) float64 {

	switch f.Severity {
	case "HIGH":
		return 9.0
	case "MEDIUM":
		return 6.5
	case "LOW":
		return 3.5
	case "INFO":
		return 1.0
	default:
		return 0.0
	}
}
