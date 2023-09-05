package logger

import "encoding/json"

type LogLevel string

const (
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
	LogLevelFatal LogLevel = "fatal"
)

type (
	Logger interface {
		// Debug logs a debug message
		Debug(msg string)

		// Info logs an info message
		Info(msg string)

		// Warn logs a warning message
		Warn(msg string)

		// Error logs an error message
		Error(msg string)

		// Fatal logs a fatal message and exits
		Fatal(msg string)

		// Log logs a message with the default log level
		Log(msg string)
	}

	Param struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	Log struct {
		Level     string  `json:"level"`
		Timestamp int64   `json:"timestamp"`
		Message   string  `json:"message"`
		Params    []Param `json:"params"`
	}
)

func (l *Log) String() string {
	encoded, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(encoded)
}
