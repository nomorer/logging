package glog

import (
	"testing"
	"os"
)

func TestSetup(t *testing.T) {
	path:= "/tmp/test_glogger"
	if err := Setup(path, LevelDebug, DailyRotate);err != nil {
		t.Errorf("Error at Setup log: %v", err)
	}

	if GetLevel() != LevelDebug {
		t.Error("Error at Level()")
	}

	SetLevel(LevelInfo)
	if GetLevel() != LevelInfo {
		t.Error("Error at SetLevel()")
	}

	Debug("some msg")
	Debugf("some %s", "msg")

	Info("some msg")
	Infof("some %s", "msg")

	Warn("some msg")
	Warnf("some %s", "msg")

	Error("some msg")
	Errorf("some %s", "msg")

	Close()

	os.RemoveAll(path)
}
