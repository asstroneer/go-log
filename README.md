# GoLog

Small library for logging in Go.

## Usage

Usage example:

```go
package main

import (
	"github.com/asstroneer/go-log/log"
	"github.com/asstroneer/go-log/log/formatter"
	"github.com/asstroneer/go-log/log/writer"
)

func main() {
	log.SetWriter(&writer.ConsoleWriter{})          // Set log writer
	log.SetFormatter(&formatter.ConsoleFormatter{}) // Set log formatter
	log.SetLogLevel(log.LogLevelInfo)               // Set min log level

	log.Fatal("This is an fatal message")
	log.Error("This is an error message")
	log.Warning("This is an warning message")
	log.Info("This is an info message")
	log.Debug("This is an debug message")

	log.Error("This is an error message with int argument: %v-%v", 1, 123)
}

```