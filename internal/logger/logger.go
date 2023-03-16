package logger

type LogLevel string

type Log struct {
	Level     LogLevel `json:"level"`
	Message   string   `json:"message"`
	TimeStamp string   `json:"timestamp"`
}

func New(dir string) {
	if dir == "" {
		dir = "./logs"
	}
}
