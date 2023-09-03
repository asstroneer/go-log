package message

type LogMessage struct {
	Level     string
	LevelCode int
	Message   string
	Arguments []any
}
