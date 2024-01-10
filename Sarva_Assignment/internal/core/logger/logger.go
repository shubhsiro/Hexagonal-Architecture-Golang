package logger

import (
	"github.com/hashicorp/go-hclog"
)

// Logger represents the core logic for logging.
type Log struct {
	logger hclog.Logger
}

// NewLogger creates a new instance of Logger.
func NewLogger(logger hclog.Logger) *Log {
	return &Log{logger: logger}
}

// Log logs a message.
func (l *Log) Log(message string) {
	l.logger.Info(message)
}

// Trace logs a trace message.
func (l *Log) Trace(message string) {
	l.logger.Trace(message)
}

// Debug logs a debug message.
func (l *Log) Debug(message string) {
	l.logger.Debug(message)
}

// Metrics logs metrics.
func (l *Log) Metrics(metrics map[string]interface{}) {
	// Convert the metrics map to a list of key-value pairs
	var keyValuePairs []interface{}
	for key, value := range metrics {
		keyValuePairs = append(keyValuePairs, key, value)
	}

	// Log the metrics using hclog.Logger
	l.logger.Info("Metrics logged", keyValuePairs...)
}
