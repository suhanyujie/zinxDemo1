package mylog

import (
	"testing"

	"go.uber.org/zap"
)

func TestLog1(t *testing.T) {
	logger := GetLogger()
	logger.Info("test11111", zap.String("tag1", "tag1_v1"))
}
