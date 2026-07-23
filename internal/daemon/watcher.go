package daemon

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	// Register SQLite driver.
	_ "modernc.org/sqlite"

	"chibi/internal/k8s"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PredictiveDaemon watches cluster state to detect anomalies.
type PredictiveDaemon struct {
	db *sql.DB
}

func NewPredictiveDaemon(dbPath string) (*PredictiveDaemon, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite state db: %w", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS anomaly_events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		namespace TEXT,
		resource_name TEXT,
		event_type TEXT,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		return nil, fmt.Errorf("failed to init db schema: %w", err)
	}

	return &PredictiveDaemon{
		db: db,
	}, nil
}

// Start watching events in a non-blocking goroutine.
func (d *PredictiveDaemon) Start(ctx context.Context) {
	go func() {
		// V2 MVP: Poll every 30s instead of complex informers for simplicity
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				d.analyzeEvents(ctx)
			}
		}
	}()
}

func (d *PredictiveDaemon) analyzeEvents(ctx context.Context) {
	if k8s.GlobalClient == nil {
		return
	}
	events, err := k8s.GlobalClient.Clientset.CoreV1().Events("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return // log error
	}

	for _, e := range events.Items {
		if e.Type == "Warning" {
			// Store in SQLite state DB
			_, _ = d.db.ExecContext(ctx, "INSERT INTO anomaly_events (namespace, resource_name, event_type) VALUES (?, ?, ?)",
				e.Namespace, e.InvolvedObject.Name, e.Reason)
		}
	}
}
