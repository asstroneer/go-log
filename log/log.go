package log

import (
	"fmt"
	"github.com/asstroneer/go-log/log/formatter"
	"github.com/asstroneer/go-log/log/message"
	"github.com/asstroneer/go-log/log/writer"
	"time"
)

const LogLevelFatal = "fatal"
const LogLevelError = "error"
const LogLevelWarning = "warning"
const LogLevelInfo = "info"
const LogLevelDebug = "debug"

var messageFormatter formatter.MessageFormatter
var messageWriter writer.MessageWriter
var logLevel = 0

func SetWriter(mWriter writer.MessageWriter) {
	messageWriter = mWriter
}

func SetFormatter(mFormatter formatter.MessageFormatter) {
	messageFormatter = mFormatter
}

func SetLogLevel(level string) {
	logLevel = logLevelToCode(level)
}

func Debug(msg string, args ...any) {
	log(LogLevelDebug, msg, args)
}

func Info(msg string, args ...any) {
	log(LogLevelInfo, msg, args)
}

func Warning(msg string, args ...any) {
	log(LogLevelWarning, msg, args)
}

func Error(msg string, args ...any) {
	log(LogLevelError, msg, args)
}

func Fatal(msg string, args ...any) {
	log(LogLevelFatal, msg, args)
}

func log(level, msg string, args []any) {
	logMessage := message.LogMessage{
		Message:   msg,
		Level:     level,
		LevelCode: logLevelToCode(level),
		Arguments: args,
		Time:      time.Now(),
	}

	if !compareLogLevel(logMessage.LevelCode) {
		return
	}

	formattedMessage, err := messageFormatter.Format(&logMessage)
	if err != nil {
		fmt.Printf("Error in log: %s\n", err.Error())
		return
	}
	err = messageWriter.Write(logMessage.Level, formattedMessage)
	if err != nil {
		fmt.Printf("Error in log: %s\n", err.Error())
		return
	}
}

func logLevelToCode(level string) int {
	switch level {
	case LogLevelFatal:
		return 0
	case LogLevelError:
		return 1
	case LogLevelWarning:
		return 2
	case LogLevelInfo:
		return 3
	case LogLevelDebug:
		return 4
	}

	return 100
}

func compareLogLevel(messageLevel int) bool {
	return !(messageLevel > logLevel) || messageLevel == logLevel
}
