package core

import (
	"github.com/leomarqueseh/BlueGuard/internal/analysis"
)

func RunScan(target string) []analysis.Finding {
	return analysis.RunAll(target)
}
