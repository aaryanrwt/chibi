package plugin

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Manager handles discovery and execution of external plugins.
type Manager struct {
	pluginDir string
}

func NewManager(pluginDir string) *Manager {
	if pluginDir == "" {
		home, _ := os.UserHomeDir()
		pluginDir = filepath.Join(home, ".chibi", "plugins")
	}
	return &Manager{
		pluginDir: pluginDir,
	}
}

// Discover finds all executable binaries in the plugin directory.
func (m *Manager) Discover() ([]string, error) {
	var plugins []string
	if _, err := os.Stat(m.pluginDir); os.IsNotExist(err) {
		return plugins, nil // No plugins directory yet
	}

	entries, err := os.ReadDir(m.pluginDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read plugin directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			// In a real implementation, we would verify if it's executable.
			plugins = append(plugins, entry.Name())
		}
	}
	return plugins, nil
}

// ExecutePlugin runs an external binary via subprocess and returns its output.
// For V1 MVP, this executes the binary using os/exec.
func (m *Manager) ExecutePlugin(name string, args []string) (string, error) {
	binaryPath := filepath.Join(m.pluginDir, name)
	if _, err := os.Stat(binaryPath); os.IsNotExist(err) {
		return "", fmt.Errorf("plugin %s not found", name)
	}

	cmd := exec.Command(binaryPath, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(out), fmt.Errorf("plugin execution failed: %w", err)
	}

	return string(out), nil
}
