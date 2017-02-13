// Package logging provides structured logging instance
package logging

import (
	"os"

	"github.com/uber-go/zap"
)

// Global instance
var logger zap.Logger

// Logger returns the logger instance
func Logger() zap.Logger {
	if logger == nil {
		// Get hostname as common logging context
		hostname, err := os.Hostname()
		if err != nil {
			panic(err)
		}

		// Intialize the logger, with JSON encoder,
		// INFO level, or higher to standard out.
		// Any errors during logging will be written to standard error.
		logger = zap.New(zap.NewJSONEncoder(),
			zap.InfoLevel,
			zap.Fields(zap.String("host", hostname)))
	}
	return logger
}
