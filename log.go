package glogger

import (
	"fmt"
	"log"
	"os"
	"time"
	"errors"
)

const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

const (
	defaultTimeFormat   = "2006-01-02 15:04:05"
	defaultPrefixFormat = "%s ▶ %.3s %s"
)

var (
	// higher means more serious, log level : DEBUG < INFO < WARN < ERROR < FATAL
	levelNames [5]string = [5]string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
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
	writer.Close()
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	output(LevelFatal, fmt.Sprintln(fmt.Sprintf(format, v...)))
	writer.Close()
	os.Exit(1)
}

func Setup(path string, level, rotateType int) error {
	if level < LevelDebug || level > LevelFatal {
		return errors.New("None Exist Level")
	}

	if rotateType < DailyRotate || level > HourlyRotate {
		return errors.New("None Exist Rotate Type")
	}
	var err error
	if writer, err = NewRotateLogger(path, level, rotateType); err != nil {
		return err
	}
	return nil
}

func SetLevel(level int) {
	writer.SetLevel(level)
}

func Level(level int) int {
	return writer.Level()
}

func output(level int, content string) {
	if level < writer.Level() {
		return
	}

	//the writer may be close
	logContent := fmt.Sprintf(defaultPrefixFormat, time.Now().Format(defaultTimeFormat), levelNames[level], content)
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