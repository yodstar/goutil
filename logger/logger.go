package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Level
const (
	NONE = iota
	FATAL
	ALERT
	ERROR
	WARN
	NOTICE
	INFO
	DEBUG
	TRACE
	ALL
)

// Mapping
var levelMap = map[string]int{
	"-":      NONE,
	"FATAL":  FATAL,
	"ALERT":  ALERT,
	"ERROR":  ERROR,
	"WARN":   WARN,
	"NOTICE": NOTICE,
	"INFO":   INFO,
	"DEBUG":  DEBUG,
	"TRACE":  TRACE,
	"ALL":    ALL,
}

// Prefix
var levelPrefix = [ALL]string{
	"-",
	"[FATAL] ",
	"[ALERT] ",
	"[ERROR] ",
	"[WARN] ",
	"[NOTICE] ",
	"[INFO] ",
	"[DEBUG] ",
	"[TRACE] ",
}

// Logger
type Logger struct {
	log.Logger
	level   int
	outfile string
	datefmt string
	nowdate string
	filter  int
	handler func(string)
}

// Logger.prepare
func (l *Logger) prepare() error {
	if l.outfile == "" {
		return nil
	}
	nowdate := time.Now().Format(l.datefmt)
	if l.nowdate != nowdate {
		outfile := fmt.Sprintf(l.outfile, nowdate)
		if err := os.MkdirAll(filepath.Dir(outfile), 0644); err != nil {
			return err
		}
		output, err := os.OpenFile(outfile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		l.SetOutput(output)
		l.nowdate = nowdate
	}
	return nil
}

// Logger.SetLevel
func (l *Logger) SetLevel(level interface{}) {
	switch lv := level.(type) {
	case int:
		l.level = lv
	case string:
		if v, ok := levelMap[level.(string)]; ok {
			l.level = v
		} else {
			l.level = 0
		}
	default:
		l.level = 0
	}
}

// Logger.SetOutFile
func (l *Logger) SetOutFile(outfile, datefmt string) {
	l.outfile = outfile
	l.datefmt = datefmt
}

// Logger.SetFilter
func (l *Logger) SetFilter(lv interface{}, fn func(string)) {
	switch lv := lv.(type) {
	case int:
		l.filter = lv
	case string:
		if v, ok := levelMap[lv]; ok {
			l.filter = v
		} else {
			l.filter = 0
		}
	default:
		l.filter = 0
	}
	l.handler = fn
}

// Logger.Outputf
func (l *Logger) Outputf(level int, args ...interface{}) {
	if level > l.level && level > l.filter {
		return
	}
	var format string
	if len(args) > 1 {
		format = levelPrefix[level] + args[0].(string)
		args = args[1:]
	} else {
		format = levelPrefix[level] + "%s"
	}
	output := fmt.Sprintf(format, args...)
	if l.filter >= level {
		l.handler(output)
	}
	if l.level >= level {
		if err := l.prepare(); err != nil {
			l.Printf("[ALERT] %s", err.Error())
		}
		l.Output(3, output)
	}
}

// Logger.Fatal
func (l *Logger) Fatal(args ...interface{}) {
	l.Outputf(FATAL, args...)
}

// Logger.Alert
func (l *Logger) Alert(args ...interface{}) {
	l.Outputf(ALERT, args...)
}

// Logger.Error
func (l *Logger) Error(args ...interface{}) {
	l.Outputf(ERROR, args...)
}

// Logger.Warn
func (l *Logger) Warn(args ...interface{}) {
	l.Outputf(WARN, args...)
}

// Logger.Notice
func (l *Logger) Notice(args ...interface{}) {
	l.Outputf(NOTICE, args...)
}

// Logger.Info
func (l *Logger) Info(args ...interface{}) {
	l.Outputf(INFO, args...)
}

// Logger.Debug
func (l *Logger) Debug(args ...interface{}) {
	l.Outputf(DEBUG, args...)
}

// Logger.Trace
func (l *Logger) Trace(args ...interface{}) {
	l.Outputf(TRACE, args...)
}

// New
func New(out io.Writer, prefix string, flag int) *Logger {
	logger := &Logger{level: INFO}
	logger.SetOutput(out)
	logger.SetPrefix(prefix)
	logger.SetFlags(flag)
	return logger
}

// LOG - Default logger
var LOG = New(os.Stderr, "", log.LstdFlags)

// SetFlags
func SetFlags(flag int) {
	LOG.SetFlags(flag)
}

// SetLevel
func SetLevel(level interface{}) {
	LOG.SetLevel(level)
}

// SetOutFile
func SetOutFile(outfile, datefmt string) {
	LOG.SetOutFile(outfile, datefmt)
}

// SetFilter
func SetFilter(lv interface{}, fn func(string)) {
	LOG.SetFilter(lv, fn)
}

// Fatal
func Fatal(args ...interface{}) {
	LOG.Outputf(FATAL, args...)
}

// Alert
func Alert(args ...interface{}) {
	LOG.Outputf(ALERT, args...)
}

// Error
func Error(args ...interface{}) {
	LOG.Outputf(ERROR, args...)
}

// Warn
func Warn(args ...interface{}) {
	LOG.Outputf(WARN, args...)
}

// Notice
func Notice(args ...interface{}) {
	LOG.Outputf(NOTICE, args...)
}

// Info
func Info(args ...interface{}) {
	LOG.Outputf(INFO, args...)
}

// Debug
func Debug(args ...interface{}) {
	LOG.Outputf(DEBUG, args...)
}

// Trace
func Trace(args ...interface{}) {
	LOG.Outputf(TRACE, args...)
}
