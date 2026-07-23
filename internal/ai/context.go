package ai

import (
	"encoding/json"
	"fmt"
)

// ContextBuilder is responsible for taking raw Kubernetes objects and
// converting them into an optimized prompt representation.
type ContextBuilder struct {
}

func NewContextBuilder() *ContextBuilder {
	return &ContextBuilder{}
}

// BuildResourceContext converts a K8s resource (like a Pod) into a minimal JSON string.
// In a production scenario, we would strip out noisy fields like managedFields.
func (c *ContextBuilder) BuildResourceContext(resource interface{}) (string, error) {
	// For V1, a simple JSON marshal is often enough.
	// We can implement recursive map cleanup to remove "managedFields" if necessary.
	bytes, err := json.MarshalIndent(resource, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal resource: %w", err)
	}
	return string(bytes), nil
}

// BuildDiagnosticPrompt constructs a prompt to diagnose a failing resource.
func (c *ContextBuilder) BuildDiagnosticPrompt(resourceType, resourceName, resourceJSON string) string {
	return fmt.Sprintf(`
You are an expert Kubernetes SRE assistant. 
Diagnose the following failing %s named "%s".

Resource YAML/JSON:
%s

Identify the root cause and provide clear, actionable steps to remediate the issue.
`, resourceType, resourceName, resourceJSON)
}
