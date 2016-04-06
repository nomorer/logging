package glogger

import (
	"fmt"
	"time"
)

const (
	dailyDateFormat = "2006-01-02"
)

type (
	RotateRule interface {
		ShallRotate() bool
		SetRotateTime()
		GetBackupFilename(filename string) string
	}

	DailyRotateRule struct {
		rotateTime string
	}
)

func (drr *DailyRotateRule) ShallRotate() bool {
	return drr.rotateTime != getCurrentDailyFormatDate() && len(drr.rotateTime) > 0
}

func (drr *DailyRotateRule) SetRotateTime() {
	drr.rotateTime = getCurrentDailyFormatDate()
}

func (drr *DailyRotateRule) GetBackupFilename(filename string) string {
	return fmt.Sprintf("%s-%s", filename, getCurrentDailyFormatDate())
}

func getCurrentDailyFormatDate() string {
	return time.Now().Format(dailyDateFormat)
}
