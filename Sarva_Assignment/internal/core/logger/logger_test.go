package logger

import (
	"testing"

	"github.com/hashicorp/go-hclog"
)

func TestLogger_Log(t *testing.T) {
	// Create a logger with info enabled
	loggerAdapter := NewLogger(hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Info,
		Output:     hclog.DefaultOutput,
		JSONFormat: true,
	}))

	// Call the Log method
	loggerAdapter.Log("Test log message")
}

func TestLogger_Trace(t *testing.T) {
	// Create a logger with trace enabled
	loggerAdapter := NewLogger(hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     hclog.DefaultOutput,
		JSONFormat: true,
	}))

	// Call the Trace method
	loggerAdapter.Trace("This is a trace message")
}

func TestLogger_Debug(t *testing.T) {
	// Create a logger with debug enabled
	loggerAdapter := NewLogger(hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Debug,
		Output:     hclog.DefaultOutput,
		JSONFormat: true,
	}))

	// Call the Debug function
	loggerAdapter.Debug("Test debug message")
}

func TestLogger_Metrics(t *testing.T) {
	// Create a logger with info enabled
	loggerAdapter := NewLogger(hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Info,
		Output:     hclog.DefaultOutput,
		JSONFormat: true,
	}))

	// Call the Metrics method
	loggerAdapter.Metrics(map[string]interface{}{"metric1": 42, "metric2": "value"})
}
