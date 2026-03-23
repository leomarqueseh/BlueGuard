package plugins

import (
	"github.com/leomarqueseh/BlueGuard/internal/core"
	"github.com/leomarqueseh/BlueGuard/internal/scanner"
)

type Plugin interface {
	Name() string
	Run(ctx *core.ScanContext, target scanner.Target) ([]scanner.Finding, error)
}
