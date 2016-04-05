package glogger

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"sync"
	"time"
)

const (
	dailyDateFormat      = "2006-01-02"
	maxWaitTime          = 100 * time.Millisecond
	bufferSize           = 100
	defaultFileMode      = 0600
	defaultDirectoryMode = 0755
)

type (
	RotateRule interface {
		ShallRotate() bool
		SetRotateTime()
		GetBackupFilename(filename string) string
	}

	RotateLogger struct {
		filename       string
		backupFilename string
		rule           RotateRule

		level int

		fp      *os.File
		channel chan []byte
		done    chan bool

		waitGroup sync.WaitGroup
	}

	DailyRotateLogger struct {
		rotateTime string
	}
)

func (drl *DailyRotateLogger) ShallRotate() bool {
	return drl.rotateTime != getNowDate() && len(drl.rotateTime) > 0
}

func (drl *DailyRotateLogger) SetRotateTime() {
	drl.rotateTime = getNowDate()
}

func (drl *DailyRotateLogger) GetBackupFilename(filename string) string {
	return fmt.Sprintf("%s-%s", filename, getNowDate())
}

func NewRotateLogger(filename string, level int) (*RotateLogger, error) {
	if filename == "" {
		return nil, nil
	}

	l := &RotateLogger{
		filename: filename,
		rule: &DailyRotateLogger{
			rotateTime: getNowDate(),
		},
		level:   level,
		channel: make(chan []byte, bufferSize),
		done:    make(chan bool),
	}

	if err := l.init(); err != nil {
		return nil, err
	}

	l.start()
	return l, nil
}

func (rl *RotateLogger) init() error {
	rl.backupFilename = rl.rule.GetBackupFilename(rl.filename)

	if _, err := os.Stat(rl.filename); err != nil {
		basePath := path.Dir(rl.filename)
		if _, err = os.Stat(basePath); err != nil {
			if err = os.MkdirAll(basePath, defaultDirectoryMode); err != nil {
				return err
			}
		}
		if rl.fp, err = os.Create(rl.filename); err != nil {
			return err
		}
	} else if rl.fp, err = os.OpenFile(rl.filename, os.O_APPEND|os.O_WRONLY, defaultFileMode); err != nil {
		return err
	}
	return nil
}

func (rl *RotateLogger) start() {
	rl.waitGroup.Add(1)

	go func() {
		defer rl.waitGroup.Done()

		for {
			select {
			case event, ok := <-rl.channel:
				if ok {
					rl.write(event)
				} else {
					return
				}
			}
		}
	}()
}

func (rl *RotateLogger) rotate() error {
	if rl.fp != nil {
		if err := rl.fp.Close(); err != nil {
			return err
		}
		rl.fp = nil
	}

	_, err := os.Stat(rl.filename)
	if err == nil && len(rl.backupFilename) > 0 {
		err = os.Rename(rl.filename, rl.backupFilename)
	}

	rl.backupFilename = rl.rule.GetBackupFilename(rl.filename)
	rl.fp, err = os.Create(rl.filename)
	return nil
}

func (rl *RotateLogger) Write(content []byte) (int, error) {
	select {
	case <-rl.done:
	default:
		select {
		case rl.channel <- content:
			return len(content), nil
		case <-time.After(maxWaitTime):
			return 0, errors.New("Timeout on writting log")
		}

	}
	return 0, nil
}

func (rl *RotateLogger) write(content []byte) {
	if rl.rule.ShallRotate() {
		if err := rl.rotate(); err != nil {
			log.Println(err)
		}
	}
	if rl.fp != nil {
		rl.fp.Write(content)
	}
}

func (rl *RotateLogger) SetLevel(level int) {
	rl.level = level
}

func (rl *RotateLogger) GetLevel() int {
	return rl.level
}

func (rl *RotateLogger) Close() error {
	close(rl.done)
	close(rl.channel)
	rl.waitGroup.Wait()
	if err := rl.fp.Sync(); err != nil {
		return err
	}
	return rl.fp.Close()
}

func getNowDate() string {
	return time.Now().Format(dailyDateFormat)
}
