package formatter

import (
	"fmt"
	"github.com/asstroneer/go-log/log/message"
	"testing"
	"time"
)

func TestJsonFormatter(t *testing.T) {
	type test struct {
		input    *message.LogMessage
		expected string
	}

	nowTime := time.Now()

	tests := []test{
		{
			input: &message.LogMessage{
				Level:     "info",
				LevelCode: 1,
				Message:   "Hello, %s",
				Time:      nowTime,
				Arguments: []interface{}{"World"},
			},
			expected: fmt.Sprintf("{\"timestamp\":\"%s\",\"level\":\"info\",\"message\":\"Hello, World\"}", nowTime.Format(time.RFC3339)),
		},
		{
			input: &message.LogMessage{
				Level:     "warning",
				LevelCode: 1,
				Message:   "Hello",
				Time:      nowTime,
				Arguments: []interface{}{},
			},
			expected: fmt.Sprintf("{\"timestamp\":\"%s\",\"level\":\"warning\",\"message\":\"Hello\"}", nowTime.Format(time.RFC3339)),
		},
	}

	consoleFormatter := &JsonFormatter{}
	for _, tt := range tests {
		actual, err := consoleFormatter.Format(tt.input)
		if err != nil {
			t.Errorf("Format(%q) error: %s", tt.input, err)
		}
		if actual != tt.expected {
			t.Errorf("Format(%q) expected: %q, actual: %q", tt.input, tt.expected, actual)
		}
	}
}
