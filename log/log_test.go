package log_test

import (
	"github.com/opinionated/utils/log"
	"os"
	"testing"
)

func TestWarning(t *testing.T) {
	log.Init(os.Stdout, os.Stdout, os.Stdout)

	log.Warn("hello world")
	log.Warn("hello", "world")
	log.Warnf("problem: %s is: %d\n", "one", 1)
}

func TestInitStd(t *testing.T) {
	log.InitStd()
	log.Info("hello world", "from", "info")
}
