package message

import "time"

type LogMessage struct {
	Level     string
	LevelCode int
	Message   string
	Time      time.Time
	Arguments []any
}
