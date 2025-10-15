package controllers

import (
	"context"
	"time"
	"fmt"
)

// Analyzer scans cluster resources and produces metrics-ready queries.
// TODO: break into smaller components (metrics client, cost calculator, recommendation store)

type Analyzer struct {
	// TODO: add Kubernetes client and logger when wiring controller-runtime
}

// NewAnalyzer creates a new Analyzer
func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}

// Run performs a single analysis pass. It should be scheduled periodically by the controller.
func (a *Analyzer) Run(ctx context.Context) error {
	// TODO: list pods, deployments, nodes, pvcs and gather request/limit info
	// TODO: query prometheus via PrometheusClient for 7-day usage stats
	// TODO: compute waste and call Recommender to produce CostRecommendation objects

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(100 * time.Millisecond):
		// placeholder work
		fmt.Println("Analyzer: placeholder run")
	}

	return nil
}
