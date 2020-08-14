package logger

import (
	"fmt"
	"io"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
)

type logEventType string

const (
	Debug logEventType = "DEBUG"
	Info  logEventType = "INFO"
	Warn  logEventType = "WARNING"
	Error logEventType = "ERROR"
)

type LogEvent struct {
	Timestamp time.Time
	Level     logEventType
	Message   string
}

func NewLogRotation(logDir string, logFileName string) (io.Writer, error) {
	tmpl := fmt.Sprintf("%s/%s.%s.log", logDir, "%Y%m%d%H%M", logFileName)

	logger, err := rotatelogs.New(
		tmpl,
		rotatelogs.WithRotationTime(time.Hour*24),
		rotatelogs.WithMaxAge(-1),
		rotatelogs.WithRotationCount(7),
	)

	if err != nil {
		return nil, err
	}

	return logger, nil
}

func NewLogEvent(level logEventType, msg string) *LogEvent {
	return &LogEvent{
		Timestamp: time.Now(),
		Level:     level,
		Message:   msg,
	}
}
