package scanner

import (
	"context"
	"time"

	"github.com/yourusername/blueguard/internal/plugins"
)

type Engine struct {
	registry *plugins.Registry
}

func NewEngine(reg *plugins.Registry) *Engine {
	return &Engine{
		registry: reg,
	}
}

func (e *Engine) Scan(ctx context.Context, target Target) (*ScanResult, error) {
	result := &ScanResult{
		Target:   target.URL,
		Findings: []Finding{},
	}

	for _, plugin := range e.registry.All() {

		select {
		case <-ctx.Done():
			return result, ctx.Err()
		default:
		}

		findings, err := plugin.Run(ctx, target)
		if err != nil {
			continue
		}

		result.Findings = append(result.Findings, findings...)
	}

	return result, nil
}
