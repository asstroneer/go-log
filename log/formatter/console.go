package formatter

import (
	"fmt"
	"github.com/asstroneer/go-log/log/message"
	"time"
)

type ConsoleFormatter struct{}

func (f *ConsoleFormatter) Format(msg *message.LogMessage) (string, error) {
	logMessage := fmt.Sprintf(msg.Message, msg.Arguments...)
	return fmt.Sprintf("[%s][%s] %s\n", time.Now().Format(time.RFC3339), msg.Level, logMessage), nil
}
