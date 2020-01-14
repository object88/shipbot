package log

import (
	"fmt"
	"io"
	"os"
)

// Level is the logging level: None, Error, Info, Verbose
type Level int

const (
	// Error means that only errors will be written
	Error Level = iota

	// Info logging writes info, warning, and error
	Info

	// Verbose logs everything bug debug-level messages
	Verbose
)

// Log is a fairly basic logger
type Log struct {
	w   io.Writer
	lvl Level
}

// CreateLog will return a log for the given name, creating
// one with the provided writer as needed
func CreateLog(w io.Writer) *Log {
	return &Log{w: w, lvl: Error}
}

// Stderr gets the log for os.Stderr
func Stderr() *Log {
	return CreateLog(os.Stderr)
}

// Stdout gets the log for os.Stdout
func Stdout() *Log {
	return CreateLog(os.Stdout)
}

// Errorf will write if the log level is at least Error.
// If the pointer receiver is nil, the log for `os.Stdout` will be used.
func (l *Log) Errorf(msg string, v ...interface{}) {
	if l == nil {
		l = Stdout()
	}

	l.writeIf(Error, msg, v...)
}

// Infof will write if the log level is at least Info.
// If the pointer receiver is nil, the log for `os.Stdout` will be used.
func (l *Log) Infof(msg string, v ...interface{}) {
	if l == nil {
		l = Stdout()
	}

	l.writeIf(Info, msg, v...)
}

// Printf will always log the given message, regardless of log level set.
// If the pointer receiver is nil, the log for `os.Stdout` will be used.
func (l *Log) Printf(msg string, v ...interface{}) {
	if l == nil {
		l = Stdout()
	}

	l.write(Verbose, msg, v...)
}

// SetLevel will adjust the logger's level.  If the pointer receiver is nil,
// the log for `os.Stdout` will be used.
func (l *Log) SetLevel(lvl Level) {
	if l == nil {
		l = Stdout()
	}

	l.lvl = lvl
}

// Verbosef will write if the log level is at least Verbose.
// If the pointer receiver is nil, the log for `os.Stdout` will be used.
func (l *Log) Verbosef(msg string, v ...interface{}) {
	if l == nil {
		l = Stdout()
	}

	l.writeIf(Verbose, msg, v...)
}

func (l *Log) writeIf(lvl Level, msg string, v ...interface{}) {
	if l.lvl < lvl {
		return
	}

	l.write(lvl, msg, v...)
}

func (l *Log) write(lvl Level, msg string, v ...interface{}) {
	m := msg
	if v != nil {
		m = fmt.Sprintf(m, v...)
	}

	l.w.Write([]byte(m))
}
