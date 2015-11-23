package log_test

import (
	"flag"
	"github.com/opinionated/utils/log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	log.InitStd()
	os.Exit(m.Run())
}

func TestWarning(t *testing.T) {
	log.Init(os.Stdout, os.Stdout, os.Stdout)

	log.Warn("hello world")
	log.Warn("hello", "world")
	log.Warnf("problem: %s is: %d\n", "one", 1)
}

func TestInitStd(t *testing.T) {
	log.InitStd()
	log.Info("hello world", "from", "info")
	log.Infof("one is: %d", 1)
	log.Info("two")
	log.Infof("three is: %d", 3)
}

func TestLine(t *testing.T) {
	log.Info("test")
}
