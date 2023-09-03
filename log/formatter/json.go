package formatter

import (
	"encoding/json"
	"fmt"
	"github.com/asstroneer/go-log/log/message"
	"time"
)

type jsonLogMessage struct {
	Timestamp string `json:"timestamp"`
	Level     string `json:"level"`
	Message   string `json:"message"`
}

type JsonFormatter struct{}

func (f *JsonFormatter) Format(msg *message.LogMessage) (string, error) {
	sprintingMessage := fmt.Sprintf(msg.Message, msg.Arguments...)

	jsonStruct := jsonLogMessage{
		Timestamp: time.Now().Format(time.RFC3339),
		Level:     msg.Level,
		Message:   sprintingMessage,
	}
	jsonString, err := json.Marshal(jsonStruct)
	if err != nil {
		return "", err
	}

	return string(jsonString), nil
}
