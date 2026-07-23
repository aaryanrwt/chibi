package ai

import (
	"strings"
	"testing"
)

func TestContextBuilder_BuildDiagnosticPrompt(t *testing.T) {
	builder := NewContextBuilder()
	prompt := builder.BuildDiagnosticPrompt("Pod", "test-pod", "{\"status\": \"Failed\"}")

	if !strings.Contains(prompt, "test-pod") {
		t.Errorf("Expected prompt to contain pod name, got: %s", prompt)
	}
	if !strings.Contains(prompt, "Failed") {
		t.Errorf("Expected prompt to contain resource JSON, got: %s", prompt)
	}
}
