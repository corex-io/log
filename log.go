package log

import (
	"fmt"
	"io"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
)

// Level
const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

// Logger log interface
type Logger interface {
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
}

// Log log
type Log struct {
	opts   Options
	lv     int
	writer io.Writer
}

var lvs = []string{"[D]", "[I]", "[W]", "[E]", "[F]"}

// New Log
func New(opts ...Option) *Log {
	options := newOptions(opts...)
	return &Log{
		opts:   options,
		lv:     DEBUG,
		writer: io.MultiWriter(options.writers...),
	}
}

const prefix = "%s %s %s:%s:%d %s\n"

func (l *Log) output(level int, format string, v ...interface{}) {
	if level < l.lv {
		return
	}
	var funcName string
	pc, file, line, ok := runtime.Caller(3)
	if ok {
		file = filepath.Base(file)
		pcs := strings.Split(runtime.FuncForPC(pc).Name(), ".")
		if len(pcs) >= 2 {
			funcName = pcs[len(pcs)-1]
		}
	} else {
		file = "-"
		funcName = "-"
		line = 0
	}
	text := fmt.Sprintf(prefix, time.Now().Format(l.opts.dateFormat), lvs[level], file, funcName, line, format)
	fmt.Fprintf(l.writer, text, v...)
}

// SetLevel set level
func (l *Log) SetLevel(lv int) {
	l.lv = lv
}

// Debugf log debug info
func (l *Log) Debugf(format string, v ...interface{}) {
	l.output(DEBUG, format, v...)
}

// Infof log info info
func (l *Log) Infof(format string, v ...interface{}) {
	l.output(INFO, format, v...)
}

// Warnf log warn info
func (l *Log) Warnf(format string, v ...interface{}) {
	l.output(WARN, format, v...)
}

// Errorf log error info
func (l *Log) Errorf(format string, v ...interface{}) {
	l.output(ERROR, format, v...)
}

// Panicf log panic info
func (l *Log) Panicf(format string, v ...interface{}) {
	l.output(FATAL, format, v...)
	fmt.Fprintf(l.writer, string(debug.Stack()))
}
