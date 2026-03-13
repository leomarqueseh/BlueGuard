package plugins

import (
	"context"

	"github.com/leomarqueseh/BlueGuard/internal/scanner"
)

func RunAll(ctx context.Context, registry *Registry, target scanner.Target) ([]scanner.Finding, error) {

	var findings []scanner.Finding

	for _, plugin := range registry.All() {

		result, err := plugin.Run(ctx, target)
		if err != nil {
			continue
		}

		findings = append(findings, result...)
	}

	return findings, nil
}
