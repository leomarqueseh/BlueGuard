package scanner

import (
	"github.com/leomarqueseh/BlueGuard/internal/core"
)

// Engine executa plugins em um target
type Engine struct {
	plugins []Plugin
}

// Plugin interface
type Plugin interface {
	Name() string
	Run(ctx *core.ScanContext, target Target) ([]Finding, error)
}

// NewEngine cria engine
func NewEngine(plugins []Plugin) *Engine {
	return &Engine{
		plugins: plugins,
	}
}

// Scan executa todos plugins e retorna findings
func (e *Engine) Scan(ctx *core.ScanContext, target Target) []Finding {

	var findings []Finding

	for _, plugin := range e.plugins {

		result, err := plugin.Run(ctx, target)
		if err != nil {
			continue
		}

		findings = append(findings, result...)
	}

	return findings
}
