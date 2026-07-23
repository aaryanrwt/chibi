package daemon

import (
	"context"
	"fmt"
)

// AutoRemediator executes suggested fixes if they match safety policies.
type AutoRemediator struct {
	autoApplyEnabled bool
}

func NewAutoRemediator(enabled bool) *AutoRemediator {
	return &AutoRemediator{
		autoApplyEnabled: enabled,
	}
}

// EvaluateAndApply checks a proposed patch against safety policies.
func (r *AutoRemediator) EvaluateAndApply(ctx context.Context, patch string, resourceType string) (string, error) {
	// For V2 MVP, we simulate a strict policy check.
	// We only allow "Deployment" restarts natively without user confirmation.
	
	if !r.autoApplyEnabled {
		return "", fmt.Errorf("auto-remediation is disabled (Dry-Run mode)")
	}

	if resourceType == "Deployment" {
		// Mock apply
		return "Successfully auto-applied deployment patch.", nil
	}

	return "", fmt.Errorf("resource %s requires manual confirmation", resourceType)
}
