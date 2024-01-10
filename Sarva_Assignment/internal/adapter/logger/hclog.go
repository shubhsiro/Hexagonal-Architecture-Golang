package logger

import (
	"github.com/hashicorp/go-hclog"
)

// HCLoggerAdapter represents the adapter for the HCLogger.
type HCLoggerAdapter struct {
	logger hclog.Logger
}

// NewHCLoggerAdapter creates a new instance of HCLoggerAdapter.
func NewHCLoggerAdapter(logger hclog.Logger) *HCLoggerAdapter {
	return &HCLoggerAdapter{logger: logger}
}

// Log logs a message.
func (l *HCLoggerAdapter) Log(message string) {
	l.logger.Info(message)
}

// Trace logs a trace message.
func (l *HCLoggerAdapter) Trace(message string) {
	l.logger.Trace(message)
}

// Debug logs a debug message.
func (l *HCLoggerAdapter) Debug(message string) {
	l.logger.Debug(message)
}

// Metrics logs metrics.
func (l *HCLoggerAdapter) Metrics(metrics map[string]interface{}) {
	var keyValuePairs []interface{}

	for key, value := range metrics {
		keyValuePairs = append(keyValuePairs, key, value)
	}

	l.logger.Info("Metrics logged", keyValuePairs...)
}
