package daemon

import (
	"context"
	"testing"
)

func TestAutoRemediator_EvaluateAndApply(t *testing.T) {
	remediator := NewAutoRemediator(true)
	ctx := context.Background()

	// Deployment is allowed natively in V2 MVP
	_, err := remediator.EvaluateAndApply(ctx, "patch", "Deployment")
	if err != nil {
		t.Errorf("Expected Deployment auto-remediation to succeed, got error: %v", err)
	}

	// Pod is NOT allowed natively
	_, err = remediator.EvaluateAndApply(ctx, "patch", "Pod")
	if err == nil {
		t.Errorf("Expected Pod auto-remediation to fail, got nil")
	}
}
