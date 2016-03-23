package glogger

import (
	"fmt"
	"io"
	"log"
	"time"
)

const (
	defaultTimeFormat   = "2006-01-02 15:04:05"
	defaultformatString = "%s ▶ %.3s %s"
)

var (
	writer io.WriteCloser
)

func Info(v ...interface{}) {
	output(writer, "INFO", fmt.Sprintln(v...))
}

func Infof(format string, v ...interface{}) {
	output(writer, "INFO", fmt.Sprintln(fmt.Sprintf(format, v...)))
}

func Warn(v ...interface{}) {
	output(writer, "WARN", fmt.Sprintln(v...))
}

func Warnf(format string, v ...interface{}) {
	output(writer, "WARN", fmt.Sprintln(fmt.Sprintf(format, v...)))
}

func Error(v ...interface{}) {
	output(writer, "ERROR", fmt.Sprintln(v...))
}

func Errorf(format string, v ...interface{}) {
	output(writer, "ERROR", fmt.Sprintln(fmt.Sprintf(format, v...)))
}

func Debug(v ...interface{}) {
	output(writer, "DEBUG", fmt.Sprintln(v...))
}

func Debugf(format string, v ...interface{}) {
	output(writer, "DEBUG", fmt.Sprintln(fmt.Sprintf(format, v...)))
}

func Fatal(v ...interface{}) {
	output(writer, "FATAL", fmt.Sprintln(v...))
}

func Fatalf(format string, v ...interface{}) {
	output(writer, "FATAL", fmt.Sprintln(fmt.Sprintf(format, v...)))
}

func Setup(path string) error {
	var err error
	if writer, err = createRotateLogger(path); err != nil {
		return err
	}

	return nil
}

func createRotateLogger(filename string) (io.WriteCloser, error) {
	return NewRotateLogger(filename)
}

func output(writer io.WriteCloser, level, content string) {
	logContent := fmt.Sprintf(defaultformatString, time.Now().Format(defaultTimeFormat), level, content)
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
