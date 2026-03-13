package plugins

import (
	"context"

	"github.com/yourusername/blueguard/internal/scanner"
)

type Plugin interface {
	ID() string
	Name() string
	Description() string
	Run(ctx context.Context, target scanner.Target) ([]scanner.Finding, error)
}

