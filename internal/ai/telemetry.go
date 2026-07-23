package ai

import (
	"fmt"
	"chibi/internal/observability"
)

// TelemetryCorrelator merges raw Kubernetes manifests with Prometheus time-series data.
type TelemetryCorrelator struct {
	promClient *observability.PrometheusClient
}

func NewTelemetryCorrelator(promClient *observability.PrometheusClient) *TelemetryCorrelator {
	return &TelemetryCorrelator{
		promClient: promClient,
	}
}

// BuildCorrelatedContext creates a combined context block containing both object state and metrics.
func (t *TelemetryCorrelator) BuildCorrelatedContext(namespace, podName string, resourceContext string) string {
	metricInfo := "Metrics unavailable (Prometheus not configured)"
	if t.promClient != nil {
		// For V2 MVP, we simulate parsing the complex PromQL return value
		metricInfo = "CPU usage: 0.15 cores (simulated correlator output)"
	}

	return fmt.Sprintf(`
Resource Context:
%s

Time-Series Telemetry (Last 5m):
%s
`, resourceContext, metricInfo)
}
