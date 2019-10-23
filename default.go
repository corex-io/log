package log

import (
	"os"

	"github.com/corex-io/log/file"
)

var log = New(Writer(os.Stderr))

//DefaultFileLog default file log
func DefaultFileLog(path ...string) *Log {
	if len(path) == 1 {
		return New(Writer(file.New(file.Path(path[0]))))
	}
	return New(Writer(file.New()))
}

// DefaultStdLog default std log
func DefaultStdLog() *Log {
	return log
}

// Debugf default debug
func Debugf(format string, v ...interface{}) {
	log.Debugf(format, v...)
}

// Infof default infof
func Infof(format string, v ...interface{}) {
	log.Infof(format, v...)
}

// Warnf default warnf
func Warnf(format string, v ...interface{}) {
	log.Warnf(format, v...)
}

// Errorf default errorf
func Errorf(format string, v ...interface{}) {
	log.Errorf(format, v...)
}
