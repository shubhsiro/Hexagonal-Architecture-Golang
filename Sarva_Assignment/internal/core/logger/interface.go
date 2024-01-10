package logger

// Logger represents the core logic for logging.
type Logger interface {
	Log(message string)
	Trace(message string)
	Debug(message string)
	Metrics(metrics map[string]interface{})
}
