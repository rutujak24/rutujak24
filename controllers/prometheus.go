package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

// PrometheusClient is a thin wrapper around the official Prometheus HTTP client.
// TODO: add authentication, TLS, timeouts, retries.

type PrometheusClient struct {
	client v1.API
}

func NewPrometheusClient(address string) (*PrometheusClient, error) {
	cfg := api.Config{Address: address}
	c, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &PrometheusClient{client: v1.NewAPI(c)}, nil
}

// QueryInstant executes an instant query
func (p *PrometheusClient) QueryInstant(ctx context.Context, q string) (string, error) {
	// TODO: return typed results; this scaffold returns raw string for simplicity
	result, warnings, err := p.client.Query(ctx, q, time.Now())
	if err != nil {
		return "", err
	}
	if len(warnings) > 0 {
		fmt.Printf("prometheus warnings: %v\n", warnings)
	}
	return result.String(), nil
}
