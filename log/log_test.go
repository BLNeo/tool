package log

import "testing"

func TestLog(t *testing.T) {
	Init("test")

	Logger.Info("test")
}
