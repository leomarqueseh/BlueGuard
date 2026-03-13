package plugins

import (
	"sync"

	"blueguard/internal/core"
)

// RunAll executa todos os plugins registrados
func RunAll(ctx *core.ScanContext) ([]core.Finding, error) {

	var (
		findings []core.Finding
		mu       sync.Mutex
		wg       sync.WaitGroup
	)

	for _, plugin := range Registry {

		wg.Add(1)

		go func(p Plugin) {
			defer wg.Done()

			results, err := p.Run(ctx)
			if err != nil {
				return
			}

			if len(results) == 0 {
				return
			}

			mu.Lock()
			findings = append(findings, results...)
			mu.Unlock()

		}(plugin)

	}

	wg.Wait()

	return findings, nil
}
