package log

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var (
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
)

func getCallerString() string {
	var buffer bytes.Buffer

	_, file, line, ok := runtime.Caller(2)
	if !ok {
		warning.Println("getCallerString could not get its caller!")
		return ""
	}

	// get the index
	opinionatedIndex := strings.LastIndex(file, "opinionated")
	if opinionatedIndex == -1 {
		warning.Println("getCallerString not called from an opinionated repo")
		return ""
	}

	buffer.WriteString(file[opinionatedIndex+11+1:])
	buffer.WriteRune(':')
	buffer.WriteString(strconv.Itoa(line))
	buffer.WriteRune(' ')
	return buffer.String()
}

// Info writes an info message to the log.
func Info(s ...interface{}) {
	outer := fmt.Sprintln(s...)
	info.Print(getCallerString(), outer)
}

// Infof writes a formatted info message to the log. Ends with new line.
func Infof(format string, args ...interface{}) {
	outer := fmt.Sprintf(format, args...)
	info.Print(getCallerString(), outer)
}

// Warn writes a warning message to the warning log.
func Warn(s ...interface{}) {
	outer := fmt.Sprintln(s...)
	warning.Print(getCallerString(), outer)
}

// Warnf write a formatted warning message to the warning log. Ends with new line.
func Warnf(format string, args ...interface{}) {
	outer := fmt.Sprintf(format, args...)
	warning.Print(getCallerString(), outer)
}

// Error write a message to the log. Does not fatal or panic.
func Error(s ...interface{}) {
	outer := fmt.Sprintln(s...)
	err.Print(getCallerString(), outer)
}

// Fatal writes a message to the error log then calls sys.exit(1).
func Fatal(s ...interface{}) {
	outer := fmt.Sprintln(s...)
	err.Fatal(getCallerString(), outer)
}

// Initialize logs to write to standard output.
func InitStd() {
	Init(os.Stdout, os.Stdout, os.Stdout)
}

// Init initializes the logs.
func Init(i io.Writer, w io.Writer, e io.Writer) {
	info = log.New(i, "INFO: ", log.Ltime)
	warning = log.New(i, "INFO: ", log.Ltime)
	err = log.New(i, "INFO: ", log.Ltime)
}
