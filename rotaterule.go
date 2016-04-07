package glogger

import (
	"fmt"
	"time"
)

const (
	DailyRotate = iota
	HourlyRotate
)

const (
	dailyDateFormat  = "2006-01-02"
	hourlyDateFormat = "2006-01-02_15"
)

type (
	RotateRule struct {
		rotateType int
		rotateTime string
	}
)

func NewRotateRule(rotateType int) *RotateRule {
	return &RotateRule{
		rotateTime: getFormatDate(rotateType),
	}
}

func (rr *RotateRule) ShallRotate() bool {
	return rr.rotateTime != getFormatDate(rr.rotateType) && len(rr.rotateTime) > 0
}

func (rr *RotateRule) SetRotateTime() {
	rr.rotateTime = getFormatDate(rr.rotateType)
}

func (rr *RotateRule) GetBackupFilename(filename string) string {
	return fmt.Sprintf("%s-%s", filename, getFormatDate(rr.rotateType))
}

func getFormatDate(rotateType int) string {
	switch rotateType {
	case HourlyRotate:
		return time.Now().Format(hourlyDateFormat)
	case DailyRotate:
		return time.Now().Format(dailyDateFormat)
	default:
		return time.Now().Format(dailyDateFormat)
	}
}
