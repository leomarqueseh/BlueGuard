package plugins

import (
	"context"

	"github.com/leomarqueseh/BlueGuard/internal/scanner"
)

type Plugin interface {
	Name() string
	Run(ctx context.Context, target scanner.Target) ([]scanner.Finding, error)
}
