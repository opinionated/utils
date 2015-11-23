package log

import (
	"io"
	"log"
	"os"
)

var (
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
)

// Info writes an info message to the log.
func Info(s ...interface{}) {
	info.Println(s...)
}

// Infof writes a formatted info message to the log.
func Infof(format string, args ...interface{}) {
	info.Printf(format, args...)
}

// Warn writes a warning message to the warning log.
func Warn(s ...interface{}) {
	warning.Println(s...)
}

// Warnf write a formatted warning message to the warning log.
func Warnf(format string, args ...interface{}) {
	warning.Printf(format, args...)
}

// Error write a message to the log. Does not fatal or panic.
func Error(s ...interface{}) {
	err.Println(s...)
}

// Fatal writes a message to the error log then calls sys.exit(1)
func Fatal(s ...interface{}) {
	err.Fatal(s...)
}

// Initialize logs to write to standard output
func InitStd() {
	Init(os.Stdout, os.Stdout, os.Stdout)
}

// Init initializes the logs.
func Init(i io.Writer, w io.Writer, e io.Writer) {
	info = log.New(i, "INFO: ", log.Ltime|log.Lshortfile)
	warning = log.New(i, "INFO: ", log.Ltime|log.Lshortfile)
	err = log.New(i, "INFO: ", log.Ltime|log.Lshortfile)
}
