package scanner

import "context"

type Plugin interface {
	Name() string
	Run(ctx context.Context, target Target) ([]Finding, error)
}

type Engine struct {
	plugins []Plugin
}

func NewEngine(p []Plugin) *Engine {
	return &Engine{
		plugins: p,
	}
}

func (e *Engine) Scan(ctx context.Context, target Target) (*ScanResult, error) {

	result := &ScanResult{
		Target:   target.URL,
		Findings: []Finding{},
	}

	for _, plugin := range e.plugins {

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
