package main

import (
	"github.com/asstroneer/go-log/log"
	"github.com/asstroneer/go-log/log/formatter"
	"github.com/asstroneer/go-log/log/writer"
)

func main() {
	log.SetWriter(&writer.ConsoleWriter{})
	log.SetFormatter(&formatter.ConsoleFormatter{})
	log.SetLogLevel(log.LogLevelInfo)

	log.Fatal("This is an fatal message")
	log.Error("This is an error message")
	log.Warning("This is an warning message")
	log.Info("This is an info message")
	log.Debug("This is an debug message")

	log.Error("This is an error message with int argument: %v-%v", 1, 123)
}
