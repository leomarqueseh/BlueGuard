package plugins

import (
	"github.com/leomarqueseh/BlueGuard/internal/core"
	"github.com/leomarqueseh/BlueGuard/internal/scanner"
)

//
// 🔹 Interface padrão dos plugins
//
type Plugin interface {
	Name() string
	Run(ctx *core.ScanContext, target scanner.Target) ([]scanner.Finding, error)
}
