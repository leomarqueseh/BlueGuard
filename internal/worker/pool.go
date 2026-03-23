package worker

import (
	"context"
	"sync"

	"github.com/leomarqueseh/BlueGuard/internal/core"
	"github.com/leomarqueseh/BlueGuard/internal/plugins"
	"github.com/leomarqueseh/BlueGuard/internal/scanner"
)

type Pool struct {
	Plugins []plugins.Plugin
	Workers int
	Ctx     *core.ScanContext
}

func NewPool(p []plugins.Plugin, workers int, ctx *core.ScanContext) *Pool {
	return &Pool{
		Plugins: p,
		Workers: workers,
		Ctx:     ctx,
	}
}

func (p *Pool) Run(ctx context.Context, targets []scanner.Target) []scanner.Finding {

	jobs := make(chan scanner.Target)
	results := make(chan []scanner.Finding)

	var wg sync.WaitGroup

	for w := 0; w < p.Workers; w++ {

		wg.Add(1)

		go func() {
			defer wg.Done()

			for target := range jobs {

				for _, plugin := range p.Plugins {

					findings, err := plugin.Run(p.Ctx, target)
					if err != nil {
						continue
					}

					results <- findings
				}
			}
		}()
	}

	go func() {

		for _, t := range targets {
			jobs <- t
		}

		close(jobs)
		wg.Wait()
		close(results)
	}()

	var all []scanner.Finding

	for r := range results {
		all = append(all, r...)
	}

	return all
}
