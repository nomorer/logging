package glogger

import (
	"fmt"
)

type Level int

//level constants
// higher level means more serious : DEBUG < INFO < WARN < ERROR < FATAL
const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

func (l Level) String() string {
	if levelStr, err := l.string(); err == nil {
		return levelStr
	}
	return "<NA>"
}

func (l Level) string() (string, error) {
	switch l {
	case LevelDebug:
		return "DEBUG", nil
	case LevelInfo:
		return "INFO", nil
	case LevelWarn:
		return "WARN", nil
	case LevelError:
		return "ERROR", nil
	case LevelFatal:
		return "FATAL", nil
	default:
		return "", fmt.Errorf("Unknown Level: %d", int(l))
	}
}
