package glogger

import (
	"errors"
	"log"
	"os"
	"path"
	"sync"
	"time"
)

const (
	maxWaitTime          = 100 * time.Millisecond
	bufferSize           = 100
	defaultFileMode      = 0600
	defaultDirectoryMode = 0755
)

type (
	RotateLogger struct {
		filename       string
		backupFilename string
		level          int
		rule           *RotateRule
		fp             *os.File
		msg            chan []byte
		done           chan bool
		waitGroup      sync.WaitGroup
	}
)

func NewRotateLogger(filename string, level, rotateType int) (*RotateLogger, error) {
	l := &RotateLogger{
		filename: filename,
		rule:     NewRotateRule(rotateType),
		level:    level,
		msg:      make(chan []byte, bufferSize),
		done:     make(chan bool),
	}

	if err := l.init(); err != nil {
		return nil, err
	}

	l.run()
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

func (rl *RotateLogger) run() {
	rl.waitGroup.Add(1)

	go func() {
		defer rl.waitGroup.Done()

		for {
			select {
			case msg, ok := <-rl.msg:
				if ok {
					rl.write(msg)
				} else {
					return
				}
			case <-rl.done:
				if len(rl.msg) == 0 {
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
		return 0, errors.New("Error: log file closed")
	default:
		select {
		case rl.msg <- content:
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
		} else {
			rl.rule.SetRotateTime()
		}
	}
	if rl.fp != nil {
		rl.fp.Write(content)
	}
}

func (rl *RotateLogger) SetLevel(level int) {
	rl.level = level
}

func (rl *RotateLogger) Level() int {
	return rl.level
}

func (rl *RotateLogger) Close() error {
	close(rl.done)
	close(rl.msg)
	rl.waitGroup.Wait()
	if err := rl.fp.Sync(); err != nil {
		return err
	}
	return rl.fp.Close()
}
