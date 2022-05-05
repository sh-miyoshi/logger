package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/pkg/errors"
)

var (
	logEnvDebug = false
	logger      = log.New(os.Stdout, "", log.LstdFlags)
)

func writeLog(level string, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	_, fname, line, _ := runtime.Caller(2)

	logger.Printf("%s:%d [%s] %s\n", fname, line, level, msg)
}

func writeCustomLog(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	logger.Printf("%s\n", msg)
}

// Init method initialize variables for logger
func Init(debugMode bool, fileName string) error {
	logEnvDebug = debugMode
	if fileName != "" {
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}

		//do not call file.Close() because logger write log through file.Writer
		logger.SetOutput(file)
	}
	return nil
}

// Debug method outputs log as DEBUG Level
func Debug(format string, a ...interface{}) {
	if logEnvDebug {
		writeLog("DEBUG", format, a...)
	}
}

// Info method outputs log as INFO Level
func Info(format string, a ...interface{}) {
	writeLog("INFO", format, a...)
}

// Error method outputs log as ERROR Level
func Error(format string, a ...interface{}) {
	writeLog("ERROR", format, a...)
}

// ErrorCustom outputs custom format log as ERROR Level
func ErrorCustom(format string, a ...interface{}) {
	writeCustomLog(format, a...)
}

// ErrorWithStack shows error with stack trace
// err must be from "github.com/pkg/errors"
func ErrorWithStack(msg string, err error) {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	e, ok := errors.Cause(err).(stackTracer)
	if !ok {
		writeCustomLog("No stacked error: %v", err)
	}
	st := e.StackTrace()
	writeLog("ERROR", "%s: %v", msg, err)
	reverse(st)
	for i, frame := range st[3:] {
		fname, line := fileInfo(frame)
		tail := "..."
		if i == len(st)-(3+1) {
			tail = "!"
		}

		logger.Printf("%s:%d [ERROR] error caused from here %s\n", fname, line, tail)
	}
}

func fileInfo(frame errors.Frame) (string, int) {
	pc := uintptr(frame) - 1

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "unknown", 0
	}
	return fn.FileLine(pc)
}

func reverse(s errors.StackTrace) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
