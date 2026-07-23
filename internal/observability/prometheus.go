package observability

import (
	"context"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

type PrometheusClient struct {
	v1api v1.API
}

func NewPrometheusClient(address string) (*PrometheusClient, error) {
	client, err := api.NewClient(api.Config{
		Address: address,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create prometheus client: %w", err)
	}
	return &PrometheusClient{
		v1api: v1.NewAPI(client),
	}, nil
}

// GetPodCPUUsage retrieves the CPU usage for a pod over the last 5 minutes.
func (p *PrometheusClient) GetPodCPUUsage(ctx context.Context, namespace, pod string) (model.Value, error) {
	query := fmt.Sprintf(`rate(container_cpu_usage_seconds_total{namespace="%s", pod="%s"}[5m])`, namespace, pod)
	result, _, err := p.v1api.Query(ctx, query, time.Now(), v1.WithTimeout(5*time.Second))
	if err != nil {
		return nil, fmt.Errorf("prometheus query failed: %w", err)
	}
	return result, nil
}
