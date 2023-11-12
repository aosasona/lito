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
		Debug(msg string, params ...Param)

		// Info logs an info message
		Info(msg string, params ...Param)

		// Warn logs a warning message
		Warn(msg string, params ...Param)

		// Error logs an error message
		Error(msg string, params ...Param)

		// Fatal logs a fatal message and exits
		Fatal(msg string, params ...Param)

		// Log logs a message with the default log level
		Log(msg string, params ...Param)

		Sync() error

		SetLogFile(path string)
	}

	Param struct {
		Key   string `json:"key"`
		Value any    `json:"value"`
	}

	Log struct {
		Level     string  `json:"level"`
		Timestamp int64   `json:"timestamp"`
		Message   string  `json:"message"`
		Params    []Param `json:"meta,omitempty"`
	}

	Logs []Log
)

func Field(key string, value any) Param {
	return Param{
		Key:   key,
		Value: value,
	}
}

func (l *Log) String() string {
	encoded, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(encoded)
}

// Marshaing a Param to JSON
func (p *Param) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		p.Key: p.Value,
	})
}

// Unmarshaling a Param from JSON
func (p *Param) UnmarshalJSON(data []byte) error {
	var obj map[string]any
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}

	for key, value := range obj {
		p.Key = key
		p.Value = value
	}

	return nil
}
