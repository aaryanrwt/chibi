package wasm

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

// Manager handles the discovery and highly-secure execution of WASM plugins.
type Manager struct {
	pluginDir string
}

func NewManager(pluginDir string) *Manager {
	if pluginDir == "" {
		home, _ := os.UserHomeDir()
		pluginDir = filepath.Join(home, ".chibi", "plugins", "wasm")
	}
	return &Manager{
		pluginDir: pluginDir,
	}
}

// ExecutePlugin loads a .wasm file and executes it in a zero-trust sandbox.
func (m *Manager) ExecutePlugin(ctx context.Context, name string, _ []string) (string, error) {
	wasmPath := filepath.Join(m.pluginDir, fmt.Sprintf("%s.wasm", name))
	wasmBytes, err := os.ReadFile(wasmPath)
	if err != nil {
		return "", fmt.Errorf("failed to read wasm plugin: %w", err)
	}

	// Create a new zero-trust WebAssembly runtime.
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)

	// Instantiate WASI, which provides sandboxed host-environment capabilities.
	// We DO NOT mount the host filesystem or network, enforcing strict security.
	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	// This is a minimal implementation for V2 MVP.
	_, err = r.Instantiate(ctx, wasmBytes)
	if err != nil {
		return "", fmt.Errorf("failed to instantiate wasm plugin: %w", err)
	}

	return "WASM execution completed securely.", nil
}
