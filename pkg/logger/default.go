package logger

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type DefaultLogger struct {
	DefaultLogLevel LogLevel
	Path            string
	buffer          Logs
	maxBufferSize   int
	mutex           sync.Mutex
}

var DefaultLogHandler = &DefaultLogger{
	DefaultLogLevel: LogLevelInfo,
	Path:            "lito.log",
}

func New(level LogLevel, path string) Logger {
	return &DefaultLogger{
		DefaultLogLevel: level,
		Path:            path,
		buffer:          make([]Log, 0),
		maxBufferSize:   256,
	}
}

func (l *DefaultLogger) SetLogFile(path string) {
	l.Path = path
}

func (l *DefaultLogger) Log(msg string, params ...Param) {
	l.log(l.DefaultLogLevel, msg, params)
}

func (l *DefaultLogger) Debug(msg string, params ...Param) {
	l.log(LogLevelDebug, msg, params)
}

func (l *DefaultLogger) Info(msg string, params ...Param) {
	l.log(LogLevelInfo, msg, params)
}

func (l *DefaultLogger) Warn(msg string, params ...Param) {
	l.log(LogLevelWarn, msg, params)
}

func (l *DefaultLogger) Error(msg string, params ...Param) {
	l.log(LogLevelError, msg, params)
}

func (l *DefaultLogger) Fatal(msg string, params ...Param) {
	l.log(LogLevelFatal, msg, params)
	os.Exit(1)
}

func (l *DefaultLogger) Sync() error {
	return l.flush()
}

func (l *DefaultLogger) SetLogLevel(level LogLevel) *DefaultLogger {
	l.DefaultLogLevel = level
	return l
}

func (l *DefaultLogger) SetMaxBufferSize(size int) *DefaultLogger {
	l.maxBufferSize = size
	return l
}

func (l *DefaultLogger) log(level LogLevel, msg string, params []Param) {
	log := l.makeLog(level, msg, params)
	l.appendToBuffer(log)
	fmt.Println(log)
}

func (l *DefaultLogger) makeLog(level LogLevel, msg string, params []Param) *Log {
	t := time.Now().Unix()
	return &Log{
		Level:     string(level),
		Timestamp: t,
		Message:   msg,
		Params:    params,
	}
}

// appendToBuffer appends a log to the buffer and flushes the buffer if it is full
func (l *DefaultLogger) appendToBuffer(log *Log) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.buffer = append(l.buffer, *log)

	if len(l.buffer) >= l.maxBufferSize {
		go func() {
			err := l.flush()
			if err != nil {
				fmt.Println(l.makeLog(LogLevelError, fmt.Sprintf("Failed to flush logger buffer: %s", err.Error()), nil))
			}
		}()
	}

	return nil
}

func (l *DefaultLogger) commit() error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	f, err := os.OpenFile(l.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	var bufferString string
	for _, log := range l.buffer {
		bufferString += log.String() + "\n"
	}

	_, err = f.WriteString(bufferString)
	if err != nil {
		return err
	}

	l.buffer = l.buffer[:0] // reset buffer by slicing it to 0
	return nil
}

func (l *DefaultLogger) flush() error {
	if err := l.commit(); err != nil {
		return err
	}
	return nil
}
