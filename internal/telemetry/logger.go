package telemetry

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

// InitLogger initializes the global Zap logger based on the debug flag.
func InitLogger(debug bool) error {
	var err error
	if debug {
		Log, err = zap.NewDevelopment()
	} else {
		Log, err = zap.NewProduction()
	}
	return err
}

// Sync flushes any buffered log entries.
func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}
