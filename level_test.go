package logging

import (
	"testing"
)

func TestString(t *testing.T) {
	if LevelInfo.String() != "INFO" {
		t.Errorf("expected INFO, %v", LevelInfo.String())
	}

	if Level(100).String() != "<NA>" {
		t.Errorf("expected <NA>, %v", Level(100).String())
	}
}
