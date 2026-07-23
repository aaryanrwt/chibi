package ai

import (
	"context"
	"fmt"
)

// Provider defines the interface that any AI provider must implement.
type Provider interface {
	// Name returns the name of the provider (e.g., "openai", "anthropic").
	Name() string
	
	// Generate generates a response for a given system prompt and user prompt.
	Generate(ctx context.Context, systemPrompt, userPrompt string) (string, error)
	
	// GenerateStream returns a channel that streams the response.
	GenerateStream(ctx context.Context, systemPrompt, userPrompt string) (<-chan string, error)
}

// Engine orchestrates AI interactions.
type Engine struct {
	ActiveProvider Provider
	Providers      map[string]Provider
}

// NewEngine creates a new AI Engine.
func NewEngine() *Engine {
	return &Engine{
		Providers: make(map[string]Provider),
	}
}

// RegisterProvider adds a provider to the engine.
func (e *Engine) RegisterProvider(p Provider) {
	e.Providers[p.Name()] = p
}

// SetActiveProvider sets the provider to use for generation.
func (e *Engine) SetActiveProvider(name string) error {
	p, ok := e.Providers[name]
	if !ok {
		return fmt.Errorf("provider %s not found", name)
	}
	e.ActiveProvider = p
	return nil
}
