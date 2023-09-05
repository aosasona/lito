package logger

import "time"

type DefaultLogger struct {
	DefaultLogLevel LogLevel
	Path            string
	buffer          []Log
	maxBufferSize   int
}

var DefaultLogHandler = DefaultLogger{
	DefaultLogLevel: LogLevelInfo,
	Path:            "lito.log",
}

func New(level LogLevel, path string) Logger {
	return &DefaultLogger{
		DefaultLogLevel: level,
		Path:            path,
		buffer:          make([]Log, 0),
		maxBufferSize:   100,
	}
}

func (l *DefaultLogger) Log(msg string) {}

func (l *DefaultLogger) Debug(msg string) {}

func (l *DefaultLogger) Info(msg string) {}

func (l *DefaultLogger) Warn(msg string) {}

func (l *DefaultLogger) Error(msg string) {}

func (l *DefaultLogger) Fatal(msg string) {}

func (l *DefaultLogger) SetLogLevel(level LogLevel) *DefaultLogger {
	l.DefaultLogLevel = level
	return l
}

func (l *DefaultLogger) SetMaxBufferSize(size int) *DefaultLogger {
	l.maxBufferSize = size
	return l
}

func (l *DefaultLogger) makeLog(level LogLevel, msg string) *Log {
	t := time.Now().Unix()
	return &Log{
		Level:     string(level),
		Timestamp: t,
		Message:   msg,
	}
}

// appendToBuffer appends a log to the buffer and flushes the buffer if it is full
func (l *DefaultLogger) appendToBuffer(log *Log) error {
	return nil
}

func (l *DefaultLogger) commit() error {
	return nil
}

func (l *DefaultLogger) flush(log *Log) error {
	return nil
}
