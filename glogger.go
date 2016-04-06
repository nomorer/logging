package glogger

import (
	"fmt"
	"log"
	"time"
)

const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal

	defaultTimeFormat   = "2006-01-02 15:04:05"
	defaultformatString = "%s ▶ %.3s %s"
)

var (
	// higher means more serious, log level : DEBUG < INFO < WARN < ERROR < FATAL
	LevelName [5]string = [5]string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	writer    *RotateLogger
)

func Debug(v ...interface{}) {
	output(LevelDebug, fmt.Sprintln(v...))
}

func Debugf(format string, v ...interface{}) {
	output(LevelDebug, fmt.Sprintln(fmt.Sprintf(format, v...)))
}

func Info(v ...interface{}) {
	output(LevelInfo, fmt.Sprintln(v...))
}

func Infof(format string, v ...interface{}) {
	output(LevelInfo, fmt.Sprintln(fmt.Sprintf(format, v...)))
}

func Warn(v ...interface{}) {
	output(LevelWarn, fmt.Sprintln(v...))
}

func Warnf(format string, v ...interface{}) {
	output(LevelWarn, fmt.Sprintln(fmt.Sprintf(format, v...)))
}

func Error(v ...interface{}) {
	output(LevelError, fmt.Sprintln(v...))
}

func Errorf(format string, v ...interface{}) {
	output(LevelError, fmt.Sprintln(fmt.Sprintf(format, v...)))
}

func Fatal(v ...interface{}) {
	output(LevelFatal, fmt.Sprintln(v...))
}

func Fatalf(format string, v ...interface{}) {
	output(LevelFatal, fmt.Sprintln(fmt.Sprintf(format, v...)))
}

func Setup(path string, level int) error {
	var err error
	if writer, err = NewRotateLogger(path, level); err != nil {
		return err
	}
	return nil
}

func SetLevel(level int) {
	writer.SetLevel(level)
}

func GetLevel(level int) int {
	return writer.GetLevel()
}

func output(level int, content string) {
	if level < writer.GetLevel() {
		return
	}

	//the writer may be close
	logContent := fmt.Sprintf(defaultformatString, time.Now().Format(defaultTimeFormat), LevelName[level], content)
	if writer != nil {
		buf := make([]byte, len(logContent))
		copy(buf, logContent)
		writer.Write(buf)
	} else {
		log.Print(logContent)
	}

}

func Close() error {
	if writer != nil {
		if err := writer.Close(); err != nil {
			return err
		} else {
			writer = nil
		}
	}
	return nil
}
