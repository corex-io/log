package log_test

import (
	"testing"

	"github.com/corex-io/log"
)

func TestDefaultLog(t *testing.T) {
	log.Debugf("debugf")
}

func TestFileLog(t *testing.T) {
	// f := file.New()
	// l := log.New(log.Writer(f))
	for {
		log.Debugf("hello")
	}
}
