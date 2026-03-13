package recon

import (
	"github.com/leomarqueseh/BlueGuard/internal/analysis"
	"github.com/leomarqueseh/BlueGuard/internal/core"
)

func Run(ctx *core.ScanContext) []analysis.Finding {

	// Recon pode fazer enumeração depois.
	// Por enquanto apenas chama analysis direto.
	return analysis.RunAll(ctx.Target)
}
