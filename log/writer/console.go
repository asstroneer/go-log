package writer

import (
	"errors"
	"fmt"
	"os"
)

type ConsoleWriter struct{}

func (c *ConsoleWriter) Write(level, msg string) error {
	switch level {
	case "debug":
		_, err := os.Stdout.WriteString(fmt.Sprintf("%s\n", msg))
		return err
	case "info":
		_, err := os.Stdout.WriteString(fmt.Sprintf("%s\n", msg))
		return err
	case "warning":
		_, err := os.Stdout.WriteString(fmt.Sprintf("%s\n", msg))
		return err
	case "error":
		_, err := os.Stderr.WriteString(fmt.Sprintf("%s\n", msg))
		return err
	case "fatal":
		_, err := os.Stderr.WriteString(fmt.Sprintf("%s\n", msg))
		return err
	}

	return errors.New("invalid log level")
}
