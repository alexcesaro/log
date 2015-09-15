package log

import (
	"testing"
)

func TestNullLogger(t *testing.T) {
	// There isn't much to test, just making sure it doesn't explode on basic usage.
	var logger Logger = NullLogger
	logger.Info("hello, world")
}
