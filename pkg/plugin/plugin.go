package plugin

import "context"

// Plugin defines the interface that all external Chibi plugins must implement.
type Plugin interface {
	// Name returns the unique identifier for the plugin.
	Name() string

	// Description returns a short summary of what the plugin does.
	Description() string

	// Execute runs the plugin logic. It receives the arguments passed by the user.
	Execute(ctx context.Context, args []string) (string, error)
}
